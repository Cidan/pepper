package schema

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Cidan/pepper/states"

	"github.com/Cidan/pepper/graph"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

type ShallowWalkFn func(string, string, string, ast.Node) error

type astVertex struct {
	state   string
	command string
	name    string
	n       ast.Node
}

type Schema struct {
	graph *graph.Digraph
	ast   []*ast.File
}

func New() *Schema {
	return &Schema{
		graph: graph.New(),
	}
}

// ReadFile reads single file and add to the AST list
func (s *Schema) ReadFile(path string) error {
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
func (s *Schema) ReadDir(dir string) error {

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

// Generate our full schema within a DAG and resolve
// any conflicts
func (s *Schema) Generate() error {
	for _, root := range s.ast {
		err := ShallowWalk(*root.Node.(*ast.ObjectList), s.createVertex)
		if err != nil {
			return err
		}
	}

	// Our graph now has every vertex, let's make the edges
	for vertex := range s.graph.Vertices() {
		v := vertex.(*astVertex)
		switch v.state {
		case "apt":
			o := &states.Apt{}
			err := hcl.DecodeObject(o, v.n)
			if err != nil {
				return err
			}
			if o.Requires != "" {
				suuid := strings.Replace(o.Requires, ".", "", -1)
				tuuid := v.state + v.command + v.name
				err := s.graph.LinkViaUUID(suuid, tuuid)
				if err == graph.ErrSourceVertexNotExists {
					return fmt.Errorf("unable to find 'requires' state '%s', which %s.%s.%s depends on",
						o.Requires, v.state, v.command, v.name)
				}
				if err == graph.ErrTargetVertexNotExists {
					return fmt.Errorf("unable to find target state %s.%s.%s which '%s' points to",
						v.state, v.command, v.name, o.Requires)
				}
				if err != nil {
					return err
				}
			} else {
				s.graph.LinkToRoot(v)
			}
		default:
			return errors.New("Unknown stanza " + v.state)
		}

	}
	op, err := s.graph.Print(s.graph.Root(), true)
	if err != nil {
		return err
	}
	fmt.Printf("Graph output success\n%s", op)
	return nil
}

func (s *Schema) createVertex(state, command, name string, n ast.Node) error {
	v := &astVertex{state, command, name, n}
	return s.graph.AddVertex(v, state+command+name)
}

// ShallowWalk will walk only the top level of the tree and call
// the supplied function with the key, command, name, and node under it.
func ShallowWalk(n ast.ObjectList, fn ShallowWalkFn) error {
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
