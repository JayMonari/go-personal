package pov

import "fmt"

// Graph is a tree of nodes that have arcs to child nodes
type Graph map[string]string

// New ...
func New() *Graph { return &Graph{} }

// AddNode does nothing because map[string]string doesn't need to be
// initialized in anyway, though the tests fail without it.
func (g *Graph) AddNode(nodeLabel string) {}

// AddArc adds an arc to the graph
func (g *Graph) AddArc(from, to string) { (*g)[to] = from }

// ArcList returns a slice of arcs formatted as: "from -> to"
func (g *Graph) ArcList() []string {
	list := make([]string, 0, len(*g))
	for to, from := range *g {
		list = append(list, fmt.Sprintf("%s -> %s", from, to))
	}
	return list
}

// ChangeRoot reroots the Graph from the oldRoot to the newRoot.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	newG := Graph{}
	for to, from := range *g {
		newG[to] = from
	}
	for curr := newRoot; curr != oldRoot; curr = (*g)[curr] {
		parent := (*g)[curr]
		newG[parent] = curr
		if curr == newRoot {
			delete(newG, curr)
		}
	}
	return &newG
}

