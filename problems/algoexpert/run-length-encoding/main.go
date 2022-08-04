package main

import "strings"

func RunLengthEncoding(s string) string {
	count := 1
	var sb strings.Builder
	for i := 0; i < len(s)-1; i++ {
		if curr := s[i]; curr != s[i+1] || count == 9 {
			sb.WriteByte(byte(count) + '0')
			sb.WriteByte(curr)
			count = 0
		}
		count++
	}
	sb.WriteByte(byte(count) + '0')
	sb.WriteByte(s[len(s)-1])
	return sb.String()
}

// Test Case 1
// {
//   "string": "AAAAAAAAAAAAABBCCCCDD"
// }
// Test Case 2
// {
//   "string": "aA"
// }
// Test Case 3
// {
//   "string": "122333"
// }
// Test Case 4
// {
//   "string": "************^^^^^^^$$$$$$%%%%%%%!!!!!!AAAAAAAAAAAAAAAAAAAA"
// }
// Test Case 5
// {
//   "string": "aAaAaaaaaAaaaAAAABbbbBBBB"
// }
// Test Case 6
// {
//   "string": "                          "
// }
// Test Case 7
// {
//   "string": "1  222 333    444  555"
// }
// Test Case 8
// {
//   "string": "1A2BB3CCC4DDDD"
// }
// Test Case 9
// {
//   "string": "........______=========AAAA   AAABBBB   BBB"
// }
// Test Case 10
// {
//   "string": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
// }
// Test Case 11
// {
//   "string": "        aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
// }
// Test Case 12
// {
//   "string": " "
// }
// Test Case 13
// {
//   "string": "[(aaaaaaa,bbbbbbb,ccccc,dddddd)]"
// }
// Test Case 14
// {
//   "string": ";;;;;;;;;;;;''''''''''''''''''''1233333332222211112222111s"
// }
// Test Case 15
// {
//   "string": "AAAAAAAAAAAAABBCCCCDDDDDDDDDDD"
// }
