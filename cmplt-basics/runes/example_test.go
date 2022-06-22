package runes_test

import "basics/runes"

func ExampleByteAndRuneAreInt() {
	runes.ByteAndRuneAreInt()
	// Output:
	// It's a
	// dec: 97 is a
	// bin: 1100001 is a
	// oct: 141 is a
	// hex: 61 is a
	// Unicode: U+0061 'a' is a
	// It's a
	// dec: 97 is a
	// bin: 1100001 is a
	// oct: 141 is a
	// hex: 61 is a
	// Unicode: U+0061 'a' is a
}

func ExampleByteCount() {
	runes.ByteCount()
	// Output:
	// Len नमस्ते दुनिया: 37
	// Len Привет мир: 19
	// Len こんにちは世界: 21
	//
	// Byte in hexadecimal: e3 and decimal: 227 and as string: ã
	// Byte in hexadecimal: 81 and decimal: 129 and as string: 
	// Byte in hexadecimal: 93 and decimal: 147 and as string: 
	// Byte in hexadecimal: e3 and decimal: 227 and as string: ã
	// Byte in hexadecimal: 82 and decimal: 130 and as string: 
	// Byte in hexadecimal: 93 and decimal: 147 and as string: 
	// Byte in hexadecimal: e3 and decimal: 227 and as string: ã
	// Byte in hexadecimal: 81 and decimal: 129 and as string: 
	// Byte in hexadecimal: ab and decimal: 171 and as string: «
	// Byte in hexadecimal: e3 and decimal: 227 and as string: ã
	// Byte in hexadecimal: 81 and decimal: 129 and as string: 
	// Byte in hexadecimal: a1 and decimal: 161 and as string: ¡
	// Byte in hexadecimal: e3 and decimal: 227 and as string: ã
	// Byte in hexadecimal: 81 and decimal: 129 and as string: 
	// Byte in hexadecimal: af and decimal: 175 and as string: ¯
	// Byte in hexadecimal: e4 and decimal: 228 and as string: ä
	// Byte in hexadecimal: b8 and decimal: 184 and as string: ¸
	// Byte in hexadecimal: 96 and decimal: 150 and as string: 
	// Byte in hexadecimal: e7 and decimal: 231 and as string: ç
	// Byte in hexadecimal: 95 and decimal: 149 and as string: 
	// Byte in hexadecimal: 8c and decimal: 140 and as string: 
}

func ExampleRuneCount() {
	runes.RuneCount()
	// Output:
	// Rune count in नमस्ते दुनिया: 13
	// Rune count in Привет мир: 10
	// Rune count in こんにちは世界: 7
	//
	// Rune verbose unicode value: U+0928 'न' Its index: 0 As a string "न"
	// Rune verbose unicode value: U+092E 'म' Its index: 3 As a string "म"
	// Rune verbose unicode value: U+0938 'स' Its index: 6 As a string "स"
	// Rune verbose unicode value: U+094D '्' Its index: 9 As a string "्"
	// Rune verbose unicode value: U+0924 'त' Its index: 12 As a string "त"
	// Rune verbose unicode value: U+0947 'े' Its index: 15 As a string "े"
	// Rune verbose unicode value: U+0020 ' ' Its index: 18 As a string " "
	// Rune verbose unicode value: U+0926 'द' Its index: 19 As a string "द"
	// Rune verbose unicode value: U+0941 'ु' Its index: 22 As a string "ु"
	// Rune verbose unicode value: U+0928 'न' Its index: 25 As a string "न"
	// Rune verbose unicode value: U+093F 'ि' Its index: 28 As a string "ि"
	// Rune verbose unicode value: U+092F 'य' Its index: 31 As a string "य"
	// Rune verbose unicode value: U+093E 'ा' Its index: 34 As a string "ा"
}

func ExampleByteToRuneForLoop() {
	runes.ByteToRuneForLoop()
	// Output:
	// Rune verbose unicode value: U+041F 'П' Its index: 0 As a string "П"
	// Rune verbose unicode value: U+0440 'р' Its index: 2 As a string "р"
	// Rune verbose unicode value: U+0438 'и' Its index: 4 As a string "и"
	// Rune verbose unicode value: U+0432 'в' Its index: 6 As a string "в"
	// Rune verbose unicode value: U+0435 'е' Its index: 8 As a string "е"
	// Rune verbose unicode value: U+0442 'т' Its index: 10 As a string "т"
	// Rune verbose unicode value: U+0020 ' ' Its index: 12 As a string " "
	// Rune verbose unicode value: U+043C 'м' Its index: 13 As a string "м"
	// Rune verbose unicode value: U+0438 'и' Its index: 15 As a string "и"
	// Rune verbose unicode value: U+0440 'р' Its index: 17 As a string "р"
}
