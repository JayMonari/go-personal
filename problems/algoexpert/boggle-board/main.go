package main

func BoggleBoard(board [][]rune, words []string) []string {
	trie := Trie{children: map[rune]Trie{}}
	for _, word := range words {
		trie.Add(word)
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	finalWords := map[string]bool{}
	for i := range board {
		for j := range board[i] {
			explore(i, j, board, trie, visited, finalWords)
		}
	}

	result := []string{}
	for word := range finalWords {
		result = append(result, word)
	}
	return result
}

func explore(i, j int, board [][]rune, trie Trie, visited [][]bool, finalWords map[string]bool) {
	if visited[i][j] {
		return
	}

	letter := board[i][j]
	if _, found := trie.children[letter]; !found {
		return
	}
	visited[i][j] = true

	trie = trie.children[letter]
	if end, found := trie.children['*']; found {
		finalWords[end.word] = true
	}

	for _, n := range getNeighbors(i, j, board) {
		explore(n[0], n[1], board, trie, visited, finalWords)
	}
	visited[i][j] = false
}

func getNeighbors(i, j int, board [][]rune) [][]int {
	neighbors := [][]int{}
	if i > 0 && j > 0 {
		neighbors = append(neighbors, []int{i - 1, j - 1})
	}
	if i > 0 && j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i - 1, j + 1})
	}
	if i < len(board)-1 && j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i + 1, j + 1})
	}
	if i < len(board)-1 && j > 0 {
		neighbors = append(neighbors, []int{i + 1, j - 1})
	}
	if i > 0 {
		neighbors = append(neighbors, []int{i - 1, j})
	}
	if i < len(board)-1 {
		neighbors = append(neighbors, []int{i + 1, j})
	}
	if j > 0 {
		neighbors = append(neighbors, []int{i, j - 1})
	}
	if j < len(board[0])-1 {
		neighbors = append(neighbors, []int{i, j + 1})
	}
	return neighbors
}

type Trie struct {
	word     string
	children map[rune]Trie
}

func (t Trie) Add(word string) {
	currNode := t
	for _, r := range word {
		if _, ok := currNode.children[r]; !ok {
			currNode.children[r] = Trie{children: map[rune]Trie{}}
		}
		currNode = currNode.children[r]
	}
	currNode.children['*'] = Trie{children: map[rune]Trie{}, word: word}
}

// Test Case 1
// {
//   "board": [
//     ["t", "h", "i", "s", "i", "s", "a"],
//     ["s", "i", "m", "p", "l", "e", "x"],
//     ["b", "x", "x", "x", "x", "e", "b"],
//     ["x", "o", "g", "g", "l", "x", "o"],
//     ["x", "x", "x", "D", "T", "r", "a"],
//     ["R", "E", "P", "E", "A", "d", "x"],
//     ["x", "x", "x", "x", "x", "x", "x"],
//     ["N", "O", "T", "R", "E", "-", "P"],
//     ["x", "x", "D", "E", "T", "A", "E"]
//   ],
//   "words": ["this", "is", "not", "a", "simple", "boggle", "board", "test", "REPEATED", "NOTRE-PEATED"]
// }
// Test Case 2
// {
//   "board": [
//     ["y", "g", "f", "y", "e", "i"],
//     ["c", "o", "r", "p", "o", "u"],
//     ["j", "u", "z", "s", "e", "l"],
//     ["s", "y", "u", "r", "h", "p"],
//     ["e", "a", "e", "g", "n", "d"],
//     ["h", "e", "l", "s", "a", "t"]
//   ],
//   "words": ["san", "sana", "at", "vomit", "yours", "help", "end", "been", "bed", "danger", "calm", "ok", "chaos", "complete", "rear", "going", "storm", "face", "epual", "dangerous"]
// }
// Test Case 3
// {
//   "board": [
//     ["a", "b", "c", "d", "e"],
//     ["f", "g", "h", "i", "j"],
//     ["k", "l", "m", "n", "o"],
//     ["p", "q", "r", "s", "t"],
//     ["u", "v", "w", "x", "y"]
//   ],
//   "words": ["agmsy", "agmsytojed", "agmsytojedinhcbgl", "agmsytojedinhcbfl"]
// }
// Test Case 4
// {
//   "board": [
//     ["a", "b"],
//     ["c", "d"]
//   ],
//   "words": ["abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "abca"]
// }
// Test Case 5
// {
//   "board": [
//     ["f", "t", "r", "o", "p", "i", "k", "b", "o"],
//     ["r", "w", "l", "p", "e", "u", "e", "a", "b"],
//     ["j", "o", "t", "s", "e", "l", "f", "l", "p"],
//     ["s", "z", "u", "t", "h", "u", "o", "p", "i"],
//     ["k", "a", "e", "g", "n", "d", "r", "g", "a"],
//     ["h", "n", "l", "s", "a", "t", "e", "t", "x"]
//   ],
//   "words": ["frozen", "rotten", "teleport", "city", "zutgatz", "kappa", "before", "rope", "obligate", "annoying"]
// }
// Test Case 6
// {
//   "board": [
//     ["c", "o", "m"],
//     ["r", "p", "l"],
//     ["c", "i", "t"],
//     ["o", "a", "e"],
//     ["f", "o", "d"],
//     ["z", "r", "b"],
//     ["g", "i", "a"],
//     ["o", "a", "g"],
//     ["f", "s", "z"],
//     ["t", "e", "i"],
//     ["t", "w", "d"]
//   ],
//   "words": ["commerce", "complicated", "twisted", "zigzag", "comma", "foobar", "baz", "there"]
// }
// Test Case 7
// {
//   "board": [
//     ["c", "o", "m"],
//     ["r", "p", "l"],
//     ["c", "i", "t"],
//     ["o", "a", "e"],
//     ["f", "o", "d"],
//     ["z", "r", "b"],
//     ["g", "i", "a"],
//     ["o", "a", "g"],
//     ["f", "s", "z"],
//     ["t", "e", "i"],
//     ["t", "w", "d"]
//   ],
//   "words": ["cr", "oc", "ml", "iao", "opo", "zrb", "big", "fs", "ogiagao", "dwd", "twt"]
// }
// Test Case 8
// {
//   "board": [
//     ["c", "o", "m"],
//     ["r", "p", "l"],
//     ["c", "i", "t"],
//     ["o", "a", "e"],
//     ["f", "o", "d"],
//     ["z", "r", "b"],
//     ["g", "i", "a"],
//     ["o", "a", "g"],
//     ["f", "s", "z"],
//     ["t", "e", "i"],
//     ["t", "w", "d"]
//   ],
//   "words": ["comlpriteacoofziraagsizefttw", "comlpriteacoofzirabagsizefottw", "comlpriteacoofziraagsizefottw", "comlpriteacoofzirabagsizeftttw"]
// }
