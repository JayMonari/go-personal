package main

// PhoneNumberMnemonics
func PhoneNumberMnemonics(phoneNo string) []string {
	mnemonic := make([]byte, len(phoneNo))
	var mnemonics []string
	phoneNumberMnemonics(0, phoneNo, mnemonic, &mnemonics)
	return mnemonics
}

func phoneNumberMnemonics(idx int, phoneNo string, curr []byte, all *[]string) {
	if idx == len(phoneNo) {
		*all = append(*all, string(curr))
		return
	}
	for _, ltr := range DigitLetters[phoneNo[idx]] {
		curr[idx] = ltr
		phoneNumberMnemonics(idx+1, phoneNo, curr, all)
	}
}

var DigitLetters = map[byte][]byte{
	'0': {'0'},
	'1': {'1'},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

// Test Case 1
// {
//   "phoneNumber": "1905"
// }
// Test Case 2
// {
//   "phoneNumber": "1111"
// }
// Test Case 3
// {
//   "phoneNumber": "002"
// }
// Test Case 4
// {
//   "phoneNumber": "444"
// }
// Test Case 5
// {
//   "phoneNumber": "9056661234"
// }
// Test Case 6
// {
//   "phoneNumber": "4163420000"
// }
// Test Case 7
// {
//   "phoneNumber": "1"
// }
// Test Case 8
// {
//   "phoneNumber": "0"
// }
// Test Case 9
// {
//   "phoneNumber": "23"
// }
// Test Case 10
// {
//   "phoneNumber": "1212"
// }
// Test Case 11
// {
//   "phoneNumber": "97"
// }
// Test Case 12
// {
//   "phoneNumber": "980016"
// }
