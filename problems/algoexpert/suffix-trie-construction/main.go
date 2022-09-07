package main

type SuffixTrie map[byte]SuffixTrie

func NewSuffixTrie() SuffixTrie { return SuffixTrie{} }

func (t SuffixTrie) PopulateSuffixTrieFrom(s string) {
	for i := range s {
		node := t
		for j := i; j < len(s); j++ {
			chr := s[j]
			if _, ok := node[chr]; !ok {
				node[chr] = SuffixTrie{}
			}
			node = node[chr]
		}
		node['*'] = nil
	}
}

func (t SuffixTrie) Contains(s string) bool {
	for i := 0; i < len(s); i++ {
		chr := s[i]
		if _, found := t[chr]; !found {
			return false
		}
		t = t[chr]
	}
	_, found := t['*']
	return found
}

// Test Case 1
// {
//   "string": "babc",
//   "classMethodsToCall": [
//     {
//       "arguments": ["abc"],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 2
// {
//   "string": "test",
//   "classMethodsToCall": [
//     {
//       "arguments": ["t"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["st"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["est"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["test"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["tes"],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 3
// {
//   "string": "invisible",
//   "classMethodsToCall": [
//     {
//       "arguments": ["e"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["le"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ble"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ible"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["sible"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["isible"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["visible"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["nvisible"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["invisible"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["nvisibl"],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 4
// {
//   "string": "1234556789",
//   "classMethodsToCall": [
//     {
//       "arguments": ["9"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["89"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["6789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["56789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["456789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["3456789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["23456789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["123456789"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["45567"],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 5
// {
//   "string": "testtest",
//   "classMethodsToCall": [
//     {
//       "arguments": ["t"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["st"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["est"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["test"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ttest"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["sttest"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["esttest"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["testtest"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["tt"],
//       "method": "contains"
//     }
//   ]
// }
// Test Case 6
// {
//   "string": "ttttttttt",
//   "classMethodsToCall": [
//     {
//       "arguments": ["t"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["tt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["tttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ttttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["tttttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ttttttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["tttttttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["ttttttttt"],
//       "method": "contains"
//     },
//     {
//       "arguments": ["vvv"],
//       "method": "contains"
//     }
//   ]
// }
