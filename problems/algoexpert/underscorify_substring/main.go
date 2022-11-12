package main

import "strings"

type (
	position  struct{ left, right int }
	positions []*position
)

func UnderscorifySubstring(s string, ss string) string {
	locations := getLocations(s, ss)
	locations = locations.collapse()
	return underscorify(s, locations)
}

func getLocations(str, substring string) positions {
	result := positions{}
	for start := 0; start < len(str); {
		nextIndex := strings.Index(str[start:], substring)
		if nextIndex == -1 {
			break
		}
		nextIndex += start
		result = append(result, &position{nextIndex, nextIndex + len(substring)})
		start = nextIndex + 1
	}
	return result
}

func (n positions) collapse() positions {
	if len(n) == 0 {
		return n
	}

	result := positions{n[0]}
	prev := n[0]
	for _, curr := range n[1:] {
		if curr.left <= prev.right {
			prev.right = curr.right
			continue
		}
		result = append(result, curr)
		prev = curr
	}
	return result
}

func underscorify(s string, posns positions) string {
	if len(posns) == 0 {
		return s
	}

	result := make([]rune, len(s)+2*len(posns))
	resIdx, pIdx := 0, 0
	for i, r := range s {
		p := posns[pIdx]
		switch i {
		case p.left:
			result[resIdx] = '_'
			resIdx++
		case p.right:
			result[resIdx] = '_'
			resIdx++
			if pIdx+1 < len(posns) {
				pIdx++
			}
		}
		result[resIdx] = r
		resIdx++
	}

	if posns[pIdx].right == len(s) {
		result[len(result)-1] = '_'
	}
	return string(result)
}

// Test Case 1
//
// {
//   "string": "testthis is a testtest to see if testestest it works",
//   "substring": "test"
// }
//
// Test Case 2
//
// {
//   "string": "this is a test to see if it works",
//   "substring": "test"
// }
//
// Test Case 3
//
// {
//   "string": "test this is a test to see if it works",
//   "substring": "test"
// }
//
// Test Case 4
//
// {
//   "string": "testthis is a test to see if it works",
//   "substring": "test"
// }
//
// Test Case 5
//
// {
//   "string": "testthis is a testest to see if testestes it works",
//   "substring": "test"
// }
//
// Test Case 6
//
// {
//   "string": "this is a test to see if it works and test",
//   "substring": "test"
// }
//
// Test Case 7
//
// {
//   "string": "this is a test to see if it works and test",
//   "substring": "bfjawkfja"
// }
//
// Test Case 8
//
// {
//   "string": "ttttttttttttttbtttttctatawtatttttastvb",
//   "substring": "ttt"
// }
//
// Test Case 9
//
// {
//   "string": "tzttztttz",
//   "substring": "ttt"
// }
//
// Test Case 10
//
// {
//   "string": "abababababababababababababaababaaabbababaa",
//   "substring": "a"
// }
//
// Test Case 11
//
// {
//   "string": "abcabcabcabcabcabcabcabcabcabcabcabcabcabc",
//   "substring": "abc"
// }
