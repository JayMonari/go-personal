package main

func MultiStringSearch(big string, smalls []string) []bool {
	trie := Trie{children: map[byte]Trie{}}
	for _, s := range smalls {
		trie.Add(s)
	}
	contained := map[string]bool{}
	for i := range big {
		findSmallStringsIn(big, i, trie, contained)
	}
	out := make([]bool, len(smalls))
	for i, s := range smalls {
		out[i] = contained[s]
	}
	return out
}

func findSmallStringsIn(s string, start int, trie Trie, contained map[string]bool) {
	current := trie
	for i := start; i < len(s); i++ {
		curr := s[i]
		if _, found := current.children[curr]; !found {
			break
		}
		current = current.children[curr]
		if end, found := current.children['*']; found {
			contained[end.word] = true
		}
	}
}

type Trie struct {
	word     string
	children map[byte]Trie
}

func (t Trie) Add(word string) {
	curr := t
	for i := range word {
		letter := word[i]
		if _, found := curr.children[letter]; !found {
			curr.children[letter] = Trie{
				children: map[byte]Trie{},
			}
		}
		curr = curr.children[letter]
	}
	curr.children['*'] = Trie{
		children: map[byte]Trie{},
		word:     word,
	}
}

// Test Case 1
//
// {
//   "bigString": "this is a big string",
//   "smallStrings": ["this", "yo", "is", "a", "bigger", "string", "kappa"]
// }
//
// Test Case 2
//
// {
//   "bigString": "abcdefghijklmnopqrstuvwxyz",
//   "smallStrings": ["abc", "mnopqr", "wyz", "no", "e", "tuuv"]
// }
//
// Test Case 3
//
// {
//   "bigString": "abcdefghijklmnopqrstuvwxyz",
//   "smallStrings": ["abcdefghijklmnopqrstuvwxyz", "abc", "j", "mnopqr", "pqrstuvwxyz", "xyzz", "defh"]
// }
//
// Test Case 4
//
// {
//   "bigString": "hj!)!%Hj1jh8f1985n!)51",
//   "smallStrings": ["%Hj7", "8f198", "!)5", "!)!", "!!", "jh81", "j181hf"]
// }
//
// Test Case 5
//
// {
//   "bigString": "Mary goes to the shopping center every week.",
//   "smallStrings": ["to", "Mary", "centers", "shop", "shopping", "string", "kappa"]
// }
//
// Test Case 6
//
// {
//   "bigString": "adcb akfkw afnmc fkadn vkaca jdaf dacb cdba cbda",
//   "smallStrings": ["abcd", "acbd", "adbc", "dabc", "cbda", "cabd", "cdab"]
// }
//
// Test Case 7
//
// {
//   "bigString": "test testing testings tests testers test-takers",
//   "smallStrings": ["tests", "testatk", "testiing", "trsatii", "test-taker", "test"]
// }
//
// Test Case 8
//
// {
//   "bigString": "ndbajwhfawkjljkfaopwdlaawjk dawkj awjkawkfjhkawk ahjwkjad jadfljawd",
//   "smallStrings": ["abc", "akwbc", "awbc", "abafac", "ajjfbc", "abac", "jadfl"]
// }
//
// Test Case 9
//
// {
//   "bigString": "Is this particular test going to pass or is it going to fail? That is the question.",
//   "smallStrings": ["that", "the", "questions", "goes", "mountain", "passes", "passed", "going", "is"]
// }
//
// Test Case 10
//
// {
//   "bigString": "Everything in this test should fail.",
//   "smallStrings": ["everything", "inn", "that", "testers", "shall", "failure"]
// }
//
// Test Case 11
//
// {
//   "bigString": "this ain't a big string",
//   "smallStrings": ["this", "is", "yo", "a", "bigger"]
// }
//
// Test Case 12
//
// {
//   "bigString": "bbbabb",
//   "smallStrings": ["bbabb"]
// }
