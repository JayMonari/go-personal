package main

func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}
	cache := make([][]*bool, len(one)+1)
	for i := 0; i < len(one)+1; i++ {
		cache[i] = make([]*bool, len(two)+1)
	}
	return areInterwoven(one, two, three, 0, 0, cache)
}

func areInterwoven(one, two, three string, i, j int, cache [][]*bool) bool {
	if cache[i][j] != nil {
		return *cache[i][j]
	}

	k := i + j
	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		result := areInterwoven(one, two, three, i+1, j, cache)
		cache[i][j] = &result
		if result {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		result := areInterwoven(one, two, three, i, j+1, cache)
		cache[i][j] = &result
		return result
	}

	result := false
	cache[i][j] = &result
	return result
}

// Test Case 1
// {
//   "one": "algoexpert",
//   "two": "your-dream-job",
//   "three": "your-algodream-expertjob"
// }
// Test Case 2
// {
//   "one": "a",
//   "two": "b",
//   "three": "ab"
// }
// Test Case 3
// {
//   "one": "a",
//   "two": "b",
//   "three": "ba"
// }
// Test Case 4
// {
//   "one": "a",
//   "two": "b",
//   "three": "ac"
// }
// Test Case 5
// {
//   "one": "abc",
//   "two": "def",
//   "three": "abcdef"
// }
// Test Case 6
// {
//   "one": "abc",
//   "two": "def",
//   "three": "adbecf"
// }
// Test Case 7
// {
//   "one": "abc",
//   "two": "def",
//   "three": "deabcf"
// }
// Test Case 8
// {
//   "one": "aabcc",
//   "two": "dbbca",
//   "three": "aadbbcbcac"
// }
// Test Case 9
// {
//   "one": "aabcc",
//   "two": "dbbca",
//   "three": "aadbbbaccc"
// }
// Test Case 10
// {
//   "one": "algoexpert",
//   "two": "your-dream-job",
//   "three": "ayloguore-xdpreeratm-job"
// }
// Test Case 11
// {
//   "one": "aaaaaaa",
//   "two": "aaaabaaa",
//   "three": "aaaaaaaaaaaaaab"
// }
// Test Case 12
// {
//   "one": "aaaaaaa",
//   "two": "aaaaaaa",
//   "three": "aaaaaaaaaaaaaa"
// }
// Test Case 13
// {
//   "one": "aacaaaa",
//   "two": "aaabaaa",
//   "three": "aaaabacaaaaaaa"
// }
// Test Case 14
// {
//   "one": "aacaaaa",
//   "two": "aaabaaa",
//   "three": "aaaacabaaaaaaa"
// }
// Test Case 15
// {
//   "one": "aacaaaa",
//   "two": "aaabaaa",
//   "three": "aaaaaacbaaaaaa"
// }
// Test Case 16
// {
//   "one": "algoexpert",
//   "two": "your-dream-job",
//   "three": "1your-algodream-expertjob"
// }
// Test Case 17
// {
//   "one": "algoexpert",
//   "two": "your-dream-job",
//   "three": "your-algodream-expertjob1"
// }
// Test Case 18
// {
//   "one": "algoexpert",
//   "two": "your-dream-job",
//   "three": "your-algodream-expertjo"
// }
// Test Case 19
// {
//   "one": "ae",
//   "two": "e",
//   "three": "see"
// }
// Test Case 20
// {
//   "one": "algo",
//   "two": "frog",
//   "three": "fralgogo"
// }
