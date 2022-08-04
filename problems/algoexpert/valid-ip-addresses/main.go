package main

import (
	"strconv"
	"strings"
)

const (
	groupCount = 4
	limit      = 4
)

func ValidIPAddresses(str string) []string {
	if len(str) < 4 {
		return []string{}
	}
	ipAddresses := []string{}
	for i := 1; i < min(len(str), limit); i++ {
		groups := make([]string, groupCount)
		if groups[0] = str[:i]; !isValid(groups[0]) {
			continue
		}
		for j := i + 1; j < i+min(len(str)-i, limit); j++ {
			if groups[1] = str[i:j]; !isValid(groups[1]) {
				continue
			}
			for k := j + 1; k < j+min(len(str)-j, limit); k++ {
				groups[2], groups[3] = str[j:k], str[k:]
				if !isValid(groups[2]) || !isValid(groups[3]) {
					continue
				}
				ipAddresses = append(ipAddresses, strings.Join(groups, "."))
			}
		}
	}
	return ipAddresses
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValid(s string) bool {
	i, err := strconv.Atoi(s)
	switch {
	case err != nil:
		return false
	case i > 255:
		return false
	}
	return len(s) == len(strconv.Itoa(i))
}

// Test Case 1
// {
//   "string": "1921680"
// }
// Test Case 2
// {
//   "string": "3700100"
// }
// Test Case 3
// {
//   "string": "9743"
// }
// Test Case 4
// {
//   "string": "97430"
// }
// Test Case 5
// {
//   "string": "997430"
// }
// Test Case 6
// {
//   "string": "255255255255"
// }
// Test Case 7
// {
//   "string": "255255255256"
// }
// Test Case 8
// {
//   "string": "99999999"
// }
// Test Case 9
// {
//   "string": "33133313"
// }
// Test Case 10
// {
//   "string": "00010"
// }
// Test Case 11
// {
//   "string": "100100"
// }
// Test Case 12
// {
//   "string": "1072310"
// }
// Test Case 13
// {
//   "string": "1"
// }
// Test Case 14
// {
//   "string": "11"
// }
// Test Case 15
// {
//   "string": "111"
// }
// Test Case 16
// {
//   "string": "00001"
// }
