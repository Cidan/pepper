package graph

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Cidan/pepper/set"
)

var (
	// ErrCycle is returned when creating an edge between two vertices would result
	// in a cycle in the digraph
	ErrCycle = errors.New("digraph: cycle between edges")

	// ErrEdgeExists is returned when an edge between two vertices already exists
	ErrEdgeExists = errors.New("digraph: edge already exists")

	// ErrVertexExists is returned when a vertex with the same value already exists
	ErrVertexExists = errors.New("digraph: vertex already exists")

	// ErrVertexNotExists is returned when a vertex is used which does not exist
	ErrVertexNotExists = errors.New("digraph: vertex does not exist")
)

// Vertex represents a vertex or "node" in the digraph
type Vertex interface{}

// Digraph represents a "digraph", or directed graph data structure
type Digraph struct {
	m           sync.RWMutex
	adjList     map[Vertex]*AdjacencyList
	uuidMap     map[string]Vertex
	edgeCount   int
	root        Vertex
	vertexCount int
}

// New creates a new acyclic Digraph, and initializes its adjacency list
func New() *Digraph {
	return &Digraph{
		adjList: map[Vertex]*AdjacencyList{},
		uuidMap: map[string]Vertex{},
	}
}

// AddVertex tries to add a new vertex to the root of the adjacency list on the digraph
func (d *Digraph) AddVertex(vertex Vertex, uuid string) error {
	// Lock digraph while adding vertex
	d.m.Lock()
	defer d.m.Unlock()

	// Check for a previous, identical vertex
	if _, found := d.adjList[vertex]; found {
		return ErrVertexExists
	}

	if _, found := d.uuidMap[uuid]; found {
		return ErrVertexExists
	}
	// Check if this vertex is the first to be added to the digraph (the root)
	if d.root == nil {
		d.root = vertex
	}
	// Add the vertex to the adjacency list, initialize a new linked-list
	d.adjList[vertex] = NewAdjacencyList()
	d.uuidMap[uuid] = vertex
	d.vertexCount++

	return nil
}

// AddEdge tries to add a new edge between two vertices on the adjacency list
func (d *Digraph) AddEdge(source Vertex, target Vertex) error {
	// Ensure vertices are not identical
	if source == target {
		return ErrCycle
	}

	if _, found := d.adjList[source]; !found {
		return ErrVertexNotExists
	}

	if _, found := d.adjList[target]; !found {
		return ErrVertexNotExists
	}

	// Check if this digraph already has this edge
	if d.HasEdge(source, target) {
		// Return false, edge already exists
		return ErrEdgeExists
	}

	// Do a depth-first search from the target to the source to determine if a cycle will
	// result if this edge is created
	if d.DepthFirstSearch(target, source) {
		// Return false, a cycle will be created
		return ErrCycle
	}

	// Lock digraph while adding edge
	d.m.Lock()
	defer d.m.Unlock()

	// Retrieve adjacency list
	adjList := d.adjList[source]

	// Target was not found, so add an edge between source and target
	adjList.list.PushBack(target)
	d.edgeCount++

	// Store adjacency list
	d.adjList[source] = adjList

	return nil
}

func (d *Digraph) LinkToRoot(target Vertex) error {
	if d.root == target {
		return nil
	}
	return d.AddEdge(d.root, target)
}

func (d *Digraph) LinkViaUUID(source, target string) error {
	s, sok := d.uuidMap[source]
	t, tok := d.uuidMap[target]
	if !sok || !tok {
		return ErrVertexNotExists
	}
	return d.AddEdge(s, t)
}

// DepthFirstSearch searches the digraph for the target vertex, using the Depth-First
// Search algorithm, and returning true if a path to the target is found
func (d *Digraph) DepthFirstSearch(source Vertex, target Vertex) bool {
	// Lock completely while performing DFS
	d.m.Lock()
	defer d.m.Unlock()

	// Generate a set of locations which have been visited
	discovered := set.New()

	// Begin recursive Depth-First Search, looking for all vertices reachable from source
	d.dfs(discovered, source)

	// Check if target was discovered during Depth-First Search
	result := discovered.Has(target)

	return result
}

// dfs implements a recursive Depth-First Search algorithm
func (d *Digraph) dfs(discovered *set.Set, target Vertex) {
	// Get the adjacency list for this vertex
	adjList := d.adjList[target]

	// Check all adjacent vertices
	for _, v := range adjList.Adjacent() {
		// Check if vertex has not been discovered
		if !discovered.Has(v) {
			// Mark it as discovered, recursively continue traversal
			discovered.Add(v)
			d.dfs(discovered, v)
		}
	}
}

// EdgeCount returns the number of edges in the digraph
func (d *Digraph) EdgeCount() int {
	d.m.Lock()
	defer d.m.Unlock()
	return d.edgeCount
}

// HasEdge determines if the digraph has an existing edge between source and target,
// returning true if it does, or false if it does not
func (d *Digraph) HasEdge(source Vertex, target Vertex) bool {
	// Lock digraph while checking for edges
	d.m.Lock()
	defer d.m.Unlock()

	// Check if the source vertex exists
	if _, found := d.adjList[source]; !found {
		return false
	}

	// Retrieve adjacency list for this source
	adjList := d.adjList[source]

	// Search for target vertex
	if v := adjList.Search(target); v != nil {
		// Vertex is adjacent, edge exists
		return true
	}

	// No result, edge does not exist
	return false
}

// Print displays a printed "tree" of the digraph to the console
func (d *Digraph) Print(root Vertex, all bool) (string, error) {
	// Lock completely during print process
	d.m.Lock()
	defer d.m.Unlock()

	// Check if the vertex actually exists
	if _, ok := d.adjList[root]; !ok {
		return "", ErrVertexNotExists
	}

	// Track locations which have already been printed
	printed := set.New()

	// Begin recursive printing at the specified root vertex
	tree := d.printRecursive(printed, root, "", all)

	return tree, nil
}

// printRecursive handles the printing of each vertex in "tree" form
func (d *Digraph) printRecursive(printed *set.Set, vertex Vertex, prefix string, all bool) string {
	// Print the current vertex
	str := fmt.Sprintf("%s - %v\n", prefix, vertex)

	// Get the current adjacency list, get adjacent vertices
	adjList := d.adjList[vertex]
	adjacent := adjList.Adjacent()

	// Iterate all adjacent vertices
	for i, v := range adjacent {
		// If not printing all, skip vertices which have already been printed
		if !all {
			if printed.Has(v) {
				continue
			}

			// Mark new ones as printed
			printed.Add(v)
		}

		// If last iteration, don't add a pipe character
		if i == len(adjacent)-1 {
			str = str + d.printRecursive(printed, v, prefix+"    ", all)
		} else {
			// Add pipe character to show multiple items belong to same parent
			str = str + d.printRecursive(printed, v, prefix+"   |", all)
		}
	}

	return str
}

// String retruns a string representation of the digraph, from the first vertex or "root"
func (d *Digraph) String() string {
	out, err := d.Print(d.root, false)
	if err != nil {
		return ""
	}

	return out
}

// VertexCount returns the number of vertices in the digraph
func (d *Digraph) VertexCount() int {
	d.m.Lock()
	defer d.m.Unlock()
	return d.vertexCount
}

// Vertices returns the full map of all vertices
func (d *Digraph) Vertices() map[Vertex]*AdjacencyList {
	return d.adjList
}

func (d *Digraph) Root() Vertex {
	return d.root
}
