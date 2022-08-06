package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	seen := map[string]struct{}{}
	curr := descendantOne
	for curr != nil {
		seen[curr.Name] = struct{}{}
		curr = curr.Ancestor
	}

	curr = descendantTwo
	for curr != nil {
		if _, found := seen[curr.Name]; found {
			return curr
		}
		curr = curr.Ancestor
	}
	return nil
}

func GetYoungestCommonAncestorOpt(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	depth1, depth2 := depthInTree(descendantOne), depthInTree(descendantTwo)
	if depth1 > depth2 {
		return backtrackTree(descendantOne, descendantTwo, depth1-depth2)
	}
	return backtrackTree(descendantTwo, descendantOne, depth2-depth1)
}

func backtrackTree(lower, higher *AncestralTree, diff int) *AncestralTree {
	for diff > 0 {
		lower = lower.Ancestor
		diff--
	}
	for lower != higher {
		lower = lower.Ancestor
		higher = higher.Ancestor
	}
	return lower
}

func depthInTree(descendant *AncestralTree) int {
	depth := 0
	for descendant != nil {
		descendant = descendant.Ancestor
		depth++
	}
	return depth
}

// Test Case 1
// {
//   "topAncestor": "A",
//   "descendantOne": "E",
//   "descendantTwo": "I",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "B", "id": "D", "name": "D"},
//       {"ancestor": "B", "id": "E", "name": "E"},
//       {"ancestor": "C", "id": "F", "name": "F"},
//       {"ancestor": "C", "id": "G", "name": "G"},
//       {"ancestor": "D", "id": "H", "name": "H"},
//       {"ancestor": "D", "id": "I", "name": "I"}
//     ]
//   }
// }
// Test Case 2
// {
//   "topAncestor": "A",
//   "descendantOne": "A",
//   "descendantTwo": "B",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 3
// {
//   "topAncestor": "A",
//   "descendantOne": "B",
//   "descendantTwo": "F",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 4
// {
//   "topAncestor": "A",
//   "descendantOne": "G",
//   "descendantTwo": "M",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 5
// {
//   "topAncestor": "A",
//   "descendantOne": "U",
//   "descendantTwo": "S",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 6
// {
//   "topAncestor": "A",
//   "descendantOne": "Z",
//   "descendantTwo": "M",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 7
// {
//   "topAncestor": "A",
//   "descendantOne": "O",
//   "descendantTwo": "I",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 8
// {
//   "topAncestor": "A",
//   "descendantOne": "T",
//   "descendantTwo": "Z",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 9
// {
//   "topAncestor": "A",
//   "descendantOne": "T",
//   "descendantTwo": "V",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 10
// {
//   "topAncestor": "A",
//   "descendantOne": "T",
//   "descendantTwo": "H",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 11
// {
//   "topAncestor": "A",
//   "descendantOne": "W",
//   "descendantTwo": "V",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 12
// {
//   "topAncestor": "A",
//   "descendantOne": "Z",
//   "descendantTwo": "B",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 13
// {
//   "topAncestor": "A",
//   "descendantOne": "Q",
//   "descendantTwo": "W",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
// Test Case 14
// {
//   "topAncestor": "A",
//   "descendantOne": "A",
//   "descendantTwo": "Z",
//   "ancestralTree": {
//     "nodes": [
//       {"ancestor": null, "id": "A", "name": "A"},
//       {"ancestor": "A", "id": "B", "name": "B"},
//       {"ancestor": "A", "id": "C", "name": "C"},
//       {"ancestor": "A", "id": "D", "name": "D"},
//       {"ancestor": "A", "id": "E", "name": "E"},
//       {"ancestor": "A", "id": "F", "name": "F"},
//       {"ancestor": "B", "id": "G", "name": "G"},
//       {"ancestor": "B", "id": "H", "name": "H"},
//       {"ancestor": "B", "id": "I", "name": "I"},
//       {"ancestor": "C", "id": "J", "name": "J"},
//       {"ancestor": "D", "id": "K", "name": "K"},
//       {"ancestor": "D", "id": "L", "name": "L"},
//       {"ancestor": "F", "id": "M", "name": "M"},
//       {"ancestor": "F", "id": "N", "name": "N"},
//       {"ancestor": "H", "id": "O", "name": "O"},
//       {"ancestor": "H", "id": "P", "name": "P"},
//       {"ancestor": "H", "id": "Q", "name": "Q"},
//       {"ancestor": "H", "id": "R", "name": "R"},
//       {"ancestor": "K", "id": "S", "name": "S"},
//       {"ancestor": "P", "id": "T", "name": "T"},
//       {"ancestor": "P", "id": "U", "name": "U"},
//       {"ancestor": "R", "id": "V", "name": "V"},
//       {"ancestor": "V", "id": "W", "name": "W"},
//       {"ancestor": "V", "id": "X", "name": "X"},
//       {"ancestor": "V", "id": "Y", "name": "Y"},
//       {"ancestor": "X", "id": "Z", "name": "Z"}
//     ]
//   }
// }
