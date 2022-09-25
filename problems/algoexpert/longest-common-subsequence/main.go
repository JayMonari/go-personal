package main

type entry struct {
	letter byte
	length int
	previ  int
	prevj  int
}

func LongestCommonSubsequence(str1 string, str2 string) []byte {
	lcs := make([][]entry, len(str2)+1)
	for i := range lcs {
		lcs[i] = make([]entry, len(str1)+1)
		for j := range lcs[i] {
			lcs[i][j].previ = -1
			lcs[i][j].prevj = -1
		}
	}

	for i := 1; i < len(str2)+1; i++ {
		for j := 1; j < len(str1)+1; j++ {
			if str2[i-1] == str1[j-1] {
				lcs[i][j] = entry{str2[i-1], lcs[i-1][j-1].length + 1, i - 1, j - 1}
				continue
			}

			lcs[i][j] = entry{0, lcs[i][j-1].length, i, j - 1}
			if lcs[i-1][j].length > lcs[i][j-1].length {
				lcs[i][j] = entry{0, lcs[i-1][j].length, i - 1, j}
			}
		}
	}
	return buildSequence(lcs)
}

func buildSequence(lcs [][]entry) []byte {
	sequence := make([]byte, 0)
	i := len(lcs) - 1
	j := len(lcs[0]) - 1
	for i != 0 && j != 0 {
		current := lcs[i][j]
		if current.letter != 0 {
			sequence = append(sequence, current.letter)
		}
		i = current.previ
		j = current.prevj
	}
	return reverse(sequence)
}

func reverse(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

// Test Case 1
// {
//   "str1": "ZXVVYZW",
//   "str2": "XKYKZPW"
// }
// Test Case 2
// {
//   "str1": "",
//   "str2": ""
// }
// Test Case 3
// {
//   "str1": "",
//   "str2": "ABCDEFG"
// }
// Test Case 4
// {
//   "str1": "ABCDEFG",
//   "str2": ""
// }
// Test Case 5
// {
//   "str1": "ABCDEFG",
//   "str2": "ABCDEFG"
// }
// Test Case 6
// {
//   "str1": "ABCDEFG",
//   "str2": "APPLES"
// }
// Test Case 7
// {
//   "str1": "clement",
//   "str2": "antoine"
// }
// Test Case 8
// {
//   "str1": "8111111111111111142",
//   "str2": "222222222822222222222222222222433333333332"
// }
// Test Case 9
// {
//   "str1": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
//   "str2": "CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAG"
// }
// Test Case 10
// {
//   "str1": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
//   "str2": "CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAGTUV"
// }
