package tree

import (
	"fmt"
	"sort"
)

// rootID is the ID that must be in every tree and is the starting ID of the
// tree.
const rootID = 0

// Node is a point in the tree with a unique ID and children Nodes.
type Node struct {
	ID       int
	Children []*Node
}

// Record is a description of a Node containing it's unique ID and its parent
// Node.
type Record struct {
	ID     int
	Parent int
}

// Build makes a tree from a slice of Records. The tree is returned as a
// pointer to the root of the tree with the children nested inside each Node,
// dependent on their Parent's ID.
func Build(records []Record) (*Node, error) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	nodes := make(map[int]*Node, len(records))
	for i, r := range records {
		if r.ID != i {
			return nil, fmt.Errorf("record not in sequence, ID: %d", r.ID)
		} else if r.Parent > r.ID || r.ID != rootID && r.Parent == r.ID {
			return nil, fmt.Errorf("bad parent ID: %d for ID: %d", r.Parent, r.ID)
		}

		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID == rootID {
			continue
		}

		nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
	}
	return nodes[0], nil
}
