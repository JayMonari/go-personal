package main

import "strings"

const (
	openers = "{[("
	closers = ")]}"
)

var matching = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

func BalancedBrackets(s string) bool {
	var stack []rune
	for _, r := range s {
		if strings.ContainsRune(openers, r) {
			stack = append(stack, r)
			continue
		}
		if !strings.ContainsRune(closers, r) {
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if stack[len(stack)-1] != matching[r] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

// Test Case 1
// {
//   "string": "([])(){}(())()()"
// }
// Test Case 2
// {
//   "string": "()[]{}{"
// }
// Test Case 3
// {
//   "string": "(((((({{{{{[[[[[([)])]]]]]}}}}}))))))"
// }
// Test Case 4
// {
//   "string": "()()[{()})]"
// }
// Test Case 5
// {
//   "string": "(()())((()()()))"
// }
// Test Case 6
// {
//   "string": "{}()"
// }
// Test Case 7
// {
//   "string": "()([])"
// }
// Test Case 8
// {
//   "string": "((){{{{[]}}}})"
// }
// Test Case 9
// {
//   "string": "((({})()))"
// }
// Test Case 10
// {
//   "string": "(([]()()){})"
// }
// Test Case 11
// {
//   "string": "(((((([[[[[[{{{{{{{{{{{{()}}}}}}}}}}}}]]]]]]))))))((([])({})[])[])[]([]){}(())"
// }
// Test Case 12
// {
//   "string": "{[[[[({(}))]]]]}"
// }
// Test Case 13
// {
//   "string": "[((([])([]){}){}){}([])[]((())"
// }
// Test Case 14
// {
//   "string": ")[]}"
// }
// Test Case 15
// {
//   "string": "(a)"
// }
// Test Case 16
// {
//   "string": "(a("
// }
// Test Case 17
// {
//   "string": "(141[])(){waga}((51afaw))()hh()"
// }
// Test Case 18
// {
//   "string": "aafwgaga()[]a{}{gggg"
// }
// Test Case 19
// {
//   "string": "(((((({{{{{safaf[[[[[([)]safsafsa)]]]]]}}}gawga}}))))))"
// }
// Test Case 20
// {
//   "string": "()(agawg)[{()gawggaw})]"
// }
// Test Case 21
// {
//   "string": "(()agwg())((()agwga()())gawgwgag)"
// }
// Test Case 22
// {
//   "string": "{}gawgw()"
// }
// Test Case 23
// {
//   "string": "(agwgg)([ghhheah%&@Q])"
// }
