package main

import "strings"

func ShortenPath(path string) string {
	tokens := []string{}
	for _, tok := range strings.Split(path, "/") {
		if len(tok) > 0 && tok != "." {
			tokens = append(tokens, tok)
		}
	}

	var stack []string
	if path[0] == '/' {
		stack = append(stack, "")
	}
	for _, tok := range tokens {
		if tok == ".." {
			switch {
			case len(stack) == 0 || stack[len(stack)-1] == "..":
				stack = append(stack, tok)
			case stack[len(stack)-1] != "":
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, tok)
	}

	if len(stack) == 1 && stack[0] == "" {
		return "/"
	}
	return strings.Join(stack, "/")
}

// Test Case 1
//
// {
//   "path": "/foo/../test/../test/../foo//bar/./baz"
// }
//
// Test Case 2
//
// {
//   "path": "/foo/bar/baz"
// }
//
// Test Case 3
//
// {
//   "path": "foo/bar/baz"
// }
//
// Test Case 4
//
// {
//   "path": "/../../foo/bar/baz"
// }
//
// Test Case 5
//
// {
//   "path": "../../foo/bar/baz"
// }
//
// Test Case 6
//
// {
//   "path": "/../../foo/../../bar/baz"
// }
//
// Test Case 7
//
// {
//   "path": "../../foo/../../bar/baz"
// }
//
// Test Case 8
//
// {
//   "path": "/foo/./././bar/./baz///////////test/../../../kappa"
// }
//
// Test Case 9
//
// {
//   "path": "../../../this////one/./../../is/../../going/../../to/be/./././../../../just/eight/double/dots/../../../../../.."
// }
//
// Test Case 10
//
// {
//   "path": "/../../../this////one/./../../is/../../going/../../to/be/./././../../../just/a/forward/slash/../../../../../.."
// }
//
// Test Case 11
//
// {
//   "path": "../../../this////one/./../../is/../../going/../../to/be/./././../../../just/eight/double/dots/../../../../../../foo"
// }
//
// Test Case 12
//
// {
//   "path": "/../../../this////one/./../../is/../../going/../../to/be/./././../../../just/a/forward/slash/../../../../../../foo"
// }
//
// Test Case 13
//
// {
//   "path": "foo/bar/.."
// }
//
// Test Case 14
//
// {
//   "path": "./foo/bar"
// }
//
// Test Case 15
//
// {
//   "path": "foo/../.."
// }
//
// Test Case 16
//
// {
//   "path": "/"
// }
//
// Test Case 17
//
// {
//   "path": "./.."
// }
