package main

import (
	"fmt"
	"strings"
)

func CaesarCipherEncryptor(str string, key int) string {
	var sb strings.Builder
	key %= 26
	for _, r := range str {
		if r = (r + rune(key)) % ('z' + 1); r < 'a' {
			r += 'a'
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func main() {
	fmt.Println(CaesarCipherEncryptor("abc", 52))
}

// Test Case 1
// {
//   "string": "xyz",
//   "key": 2
// }
// Test Case 2
// {
//   "string": "abc",
//   "key": 0
// }
// Test Case 3
// {
//   "string": "abc",
//   "key": 3
// }
// Test Case 4
// {
//   "string": "xyz",
//   "key": 5
// }
// Test Case 5
// {
//   "string": "abc",
//   "key": 26
// }
// Test Case 6
// {
//   "string": "abc",
//   "key": 52
// }
// Test Case 7
// {
//   "string": "abc",
//   "key": 57
// }
// Test Case 8
// {
//   "string": "xyz",
//   "key": 25
// }
// Test Case 9
// {
//   "string": "iwufqnkqkqoolxzzlzihqfm",
//   "key": 25
// }
// Test Case 10
// {
//   "string": "ovmqkwtujqmfkao",
//   "key": 52
// }
// Test Case 11
// {
//   "string": "mvklahvjcnbwqvtutmfafkwiuagjkzmzwgf",
//   "key": 7
// }
// Test Case 12
// {
//   "string": "kjwmntauvjjnmsagwgawkagfuaugjhawgnawgjhawjgawbfawghesh",
//   "key": 15
// }
