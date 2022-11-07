package main

import "strings"

type counts struct{ x, y int }

func PatternMatcher(pattern string, s string) []string {
	out := []string{}
	if len(pattern) > len(s) {
		return out
	}
	pattern, switched := flip(pattern)
	count := count(pattern)
	switch count.y {
	case 0:
		if len(s)%count.x != 0 {
			break
		}
		x := s[:len(s)/count.x]
		if s != strings.Repeat(x, len(pattern)) {
			break
		}
		switch switched {
		case true:
			out = []string{"", x}
		case false:
			out = []string{x, ""}
		}
	default:
		for lenx := 1; lenx < len(s); lenx++ {
			totalLeny := len(s) - lenx*count.x
			if len(s) <= lenx*count.x || totalLeny%count.y != 0 {
				continue
			}
			yindex := strings.Index(pattern, "y") * lenx
			x, y := s[:lenx], s[yindex:yindex+totalLeny/count.y]
			if s != placeWords(pattern, x, y) {
				continue
			}
			switch switched {
			case true:
				out = []string{y, x}
			case false:
				out = []string{x, y}
			}
			break
		}
	}
	return out
}

func placeWords(pattern, x, y string) string {
	var sb strings.Builder
	for _, r := range pattern {
		switch r {
		case 'x':
			sb.WriteString(x)
		case 'y':
			sb.WriteString(y)
		}
	}
	return sb.String()
}

func flip(pattern string) (string, bool) {
	if pattern[0] == 'x' {
		return pattern, false
	}
	runes := make([]rune, len(pattern))
	for i := range pattern {
		switch pattern[i] {
		case 'x':
			runes[i] = 'y'
		case 'y':
			runes[i] = 'x'
		}
	}
	return string(runes), true
}

func count(pattern string) counts {
	var c counts
	for _, r := range pattern {
		switch r {
		case 'x':
			c.x++
		case 'y':
			c.y++
		}
	}
	return c
}

// Test Case 1
//
// {
//   "pattern": "xxyxxy",
//   "string": "gogopowerrangergogopowerranger"
// }
//
// Test Case 2
//
// {
//   "pattern": "xyxy",
//   "string": "abab"
// }
//
// Test Case 3
//
// {
//   "pattern": "yxyx",
//   "string": "abab"
// }
//
// Test Case 4
//
// {
//   "pattern": "yxx",
//   "string": "yomama"
// }
//
// Test Case 5
//
// {
//   "pattern": "yyxyyx",
//   "string": "gogopowerrangergogopowerranger"
// }
//
// Test Case 6
//
// {
//   "pattern": "xyx",
//   "string": "thisshouldobviouslybewrong"
// }
//
// Test Case 7
//
// {
//   "pattern": "xxxx",
//   "string": "testtesttesttest"
// }
//
// Test Case 8
//
// {
//   "pattern": "yyyy",
//   "string": "testtesttesttest"
// }
//
// Test Case 9
//
// {
//   "pattern": "xxyxyy",
//   "string": "testtestwrongtestwrongtest"
// }
//
// Test Case 10
//
// {
//   "pattern": "xyxxxyyx",
//   "string": "baddaddoombaddadoomibaddaddoombaddaddoombaddaddoombaddaddoomibaddaddoomibaddaddoom"
// }
//
// Test Case 11
//
// {
//   "pattern": "yxyyyxxy",
//   "string": "baddaddoombaddaddoomibaddaddoombaddaddoombaddaddoombaddaddoomibaddaddoomibaddaddoom"
// }
//
// Test Case 12
//
// {
//   "pattern": "xyxxxyyx",
//   "string": "baddaddoombaddaddoomibaddaddoombaddaddoombaddaddoombaddaddoomibaddaddoomibaddaddoom"
// }
