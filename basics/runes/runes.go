package runes

import (
	"fmt"
	"unicode/utf8"
)

const helloWorldHindi = "नमस्ते दुनिया"
const helloWorldRussian = "Привет мир"
const helloWorldJapanese = "こんにちは世界"

// ByteCount shows the usual way of thinking about looping through a string
// will not work if they do not fall in the ASCII Table.
// For the complete table you can visit https://asciitable.com
func ByteCount() {
	fmt.Printf("Len %s: %d\n", helloWorldHindi, len(helloWorldHindi))
	fmt.Printf("Len %s: %d\n", helloWorldRussian, len(helloWorldRussian))
	fmt.Printf("Len %s: %d\n", helloWorldJapanese, len(helloWorldJapanese))
	fmt.Println()
	for i := 0; i < len(helloWorldJapanese); i++ {
		fmt.Printf("Byte in hexadecimal: %x and decimal: %d and as string: %s\n",
			helloWorldJapanese[i], helloWorldJapanese[i], string(helloWorldJapanese[i]))
	}
}

// RuneCount shows how to get the actual count of characters in a string if
// they do not fall in the ASCII Table.
func RuneCount() {
	fmt.Printf("Rune count in %s: %d\n", helloWorldHindi, utf8.RuneCountInString(helloWorldHindi))
	fmt.Printf("Rune count in %s: %d\n", helloWorldRussian, utf8.RuneCountInString(helloWorldRussian))
	fmt.Printf("Rune count in %s: %d\n", helloWorldJapanese, utf8.RuneCountInString(helloWorldJapanese))
}

// RuneRange shows you how to loop through a string and get the runes using the
// range keyword or by counting the amount of bytes and moving forward by that
// width.
func RuneRange() {
	for i, r := range helloWorldHindi {
		fmt.Printf("Rune verbose unicode value: %#U Its index: %d As a string %q\n",
			r, i, string(r))
	}

	for i, width := 0, 0; i < len(helloWorldRussian); i += width {
		r, w := utf8.DecodeRuneInString(helloWorldRussian[i:])
		fmt.Printf("Rune verbose unicode value: %#U Its index: %d As a string %q\n",
			r, i, string(r))
		width = w
	}
}
