package plan

// A plan is a full set of actions to carry out on a single
// resource/server. It contains the full DAG for what steps
// to take in order to meet the required state.

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Cidan/pepper/graph"
	"github.com/Cidan/pepper/states"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/mitchellh/mapstructure"
)

// ShallowWalkFn func def
type ShallowWalkFn func(string, string, string, ast.Node) error

type astVertex struct {
	state   string
	command string
	name    string
	n       map[string]interface{}
	states  states.States
}

// Plan check
type Plan struct {
	graph *graph.Digraph
	ast   []*ast.File
}

// New Stuff
func New() *Plan {
	return &Plan{
		graph: graph.New(),
	}
}

// ReadFile reads single file and add to the AST list
func (s *Plan) ReadFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	hclRoot, err := hcl.ParseBytes(data)
	if err != nil {
		return err
	}

	s.ast = append(s.ast, hclRoot)
	return nil
}

// ReadDir will read an entire directory for HCL files
// and add it to the AST list
func (s *Plan) ReadDir(dir string) error {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := s.ReadFile(dir + "/" + file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// Generate our full Plan within a DAG and resolve
// any conflicts
func (s *Plan) Generate() error {
	for _, root := range s.ast {
		err := shallowWalk(*root.Node.(*ast.ObjectList), s.createVertex)
		if err != nil {
			return err
		}
	}

	// Our graph now has every vertex, let's make the edges
	for vertex := range s.graph.Vertices() {
		v := vertex.(*astVertex)
		err := s.checkReq(v)
		if err != nil {
			return err
		}
		err = s.getState(v)
		if err != nil {
			return err
		}
	}

	// TODO:
	// DAG is now created, let's walk the graph and generate
	// the execution steps.

	// TODO: move this to a print func
	op, err := s.graph.Print(s.graph.Root(), true)
	if err != nil {
		return err
	}
	fmt.Printf("Graph output success\n%s", op)
	return nil
}

// Execute the plan
func (s *Plan) Execute() {
	s.graph.Walk(func(v graph.Vertex) {
		fmt.Printf("executing %s\n", v.(*astVertex).name)
		v.(*astVertex).states.Execute()
	})
}

// getState will generate a state object for this node and
// update the node.
func (s *Plan) getState(v *astVertex) error {
	switch v.state {
	case "apt":
		var o *states.Apt
		if err := s.decode(v.n, &o); err != nil {
			return err
		}
		v.states = o
	case "shell":
		var o *states.Shell
		if err := s.decode(v.n, &o); err != nil {
			return err
		}
		v.states = o
	default:
		return errors.New("Unknown stanza " + v.state)
	}
	return nil
}

func (s *Plan) decode(m map[string]interface{}, raw interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		Metadata:    nil,
		Result:      raw,
	})

	if err != nil {
		return err
	}

	return decoder.Decode(m)
}

func (s *Plan) checkReq(v *astVertex) error {
	switch req := v.n["requires"].(type) {
	case []string:
		for _, r := range req {
			err := s.setEdge(r, v)
			if err != nil {
				return err
			}
		}
	case string:
		err := s.setEdge(req, v)
		if err != nil {
			return err
		}
	case nil:
		err := s.setEdge("", v)
		if err != nil {
			return err
		}
	}
	// Delete the requires stanza
	delete(v.n, "requires")
	return nil
}

func (s *Plan) setEdge(req string, v *astVertex) error {

	if req != "" {
		suuid := strings.Replace(req, ".", "", -1)
		tuuid := v.state + v.command + v.name
		err := s.graph.LinkViaUUID(suuid, tuuid)
		if err == graph.ErrSourceVertexNotExists {
			return fmt.Errorf("unable to find 'requires' state '%s', which %s.%s.%s depends on",
				req, v.state, v.command, v.name)
		}
		if err == graph.ErrTargetVertexNotExists {
			return fmt.Errorf("unable to find target state %s.%s.%s which '%s' points to",
				v.state, v.command, v.name, req)
		}
		if err != nil {
			return err
		}
	} else {
		s.graph.LinkToRoot(v)
	}
	return nil
}

func (s *Plan) createVertex(state, command, name string, n ast.Node) error {
	m := make(map[string]interface{})
	err := hcl.DecodeObject(&m, n)
	if err != nil {
		return err
	}
	v := &astVertex{state, command, name, m, nil}
	return s.graph.AddVertex(v, state+command+name)
}

// ShallowWalk will walk only the top level of the tree and call
// the supplied function with the key, command, name, and node under it.
func shallowWalk(n ast.ObjectList, fn ShallowWalkFn) error {
	for _, item := range n.Items {
		if len(item.Keys) < 3 {
			return errors.New("Invalid state")
		}
		err := fn(item.Keys[0].Token.Text, item.Keys[1].Token.Text, item.Keys[2].Token.Text, item.Val)
		if err != nil {
			return err
		}
	}
	return nil
}
