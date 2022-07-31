package main

func GenerateDocument(characters string, document string) bool {
	counter := map[rune]int{}
	for _, r := range characters {
		counter[r]++
	}
	for _, r := range document {
		if c, ok := counter[r]; !ok || c == 0 {
			return false
		}
		counter[r]--
	}
	return true
}

// Test Case 1
// {
//   "characters": "Bste!hetsi ogEAxpelrt x ",
//   "document": "AlgoExpert is the Best!"
// }
// Test Case 2
// {
//   "characters": "A",
//   "document": "a"
// }
// Test Case 3
// {
//   "characters": "a",
//   "document": "a"
// }
// Test Case 4
// {
//   "characters": "a hsgalhsa sanbjksbdkjba kjx",
//   "document": ""
// }
// Test Case 5
// {
//   "characters": " ",
//   "document": "hello"
// }
// Test Case 6
// {
//   "characters": "     ",
//   "document": "     "
// }
// Test Case 7
// {
//   "characters": "aheaollabbhb",
//   "document": "hello"
// }
// Test Case 8
// {
//   "characters": "aheaolabbhb",
//   "document": "hello"
// }
// Test Case 9
// {
//   "characters": "estssa",
//   "document": "testing"
// }
// Test Case 10
// {
//   "characters": "Bste!hetsi ogEAxpert",
//   "document": "AlgoExpert is the Best!"
// }
// Test Case 11
// {
//   "characters": "helloworld ",
//   "document": "hello wOrld"
// }
// Test Case 12
// {
//   "characters": "helloworldO",
//   "document": "hello wOrld"
// }
// Test Case 13
// {
//   "characters": "helloworldO ",
//   "document": "hello wOrld"
// }
// Test Case 14
// {
//   "characters": "&*&you^a%^&8766 _=-09     docanCMakemthisdocument",
//   "document": "Can you make this document &"
// }
// Test Case 15
// {
//   "characters": "abcabcabcacbcdaabc",
//   "document": "bacaccadac"
// }
