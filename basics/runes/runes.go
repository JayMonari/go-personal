package runes

import (
	"fmt"
	"unicode/utf8"
)

const helloWorldHindi = "नमस्ते दुनिया"
const helloWorldRussian = "Привет мир"
const helloWorldJapanese = "こんにちは世界"

// ByteAndRuneAreInt shows that we can compare characters to their integer
// counterpart in any common base we want, e.g. binary, octal, decimal,
// hexadecimal, or even Unicode code point. And that `byte` fits in `rune`
// making it a smaller version of a `rune`.
func ByteAndRuneAreInt() {
	var myByte byte = 'a'
	if myByte == 'a' {
		fmt.Println("It's a")
	}
	if myByte == 97 {
		fmt.Printf("dec: %d is %c\n", myByte, myByte)
	}
	if myByte == 0b01100001 {
		fmt.Printf("bin: %b is %c\n", myByte, myByte)
	}
	if myByte == 0o141 {
		fmt.Printf("oct: %o is %c\n", myByte, myByte)
	}
	if myByte == 0x61 {
		fmt.Printf("hex: %x is %c\n", myByte, myByte)
	}
	if myByte == []byte("\u0061")[0] {
		fmt.Printf("Unicode: %#U is %c\n", myByte, myByte)
	}
	myRune := rune(myByte)
	if myRune == 'a' {
		fmt.Println("It's a")
	}
	if myRune == 97 {
		fmt.Printf("dec: %d is %c\n", myRune, myRune)
	}
	if myRune == 0b01100001 {
		fmt.Printf("bin: %b is %c\n", myRune, myRune)
	}
	if myRune == 0o141 {
		fmt.Printf("oct: %o is %c\n", myRune, myRune)
	}
	if myRune == 0x61 {
		fmt.Printf("hex: %x is %c\n", myRune, myRune)
	}
	if myRune == rune([]byte("\u0061")[0]) {
		fmt.Printf("Unicode: %#U is %c\n", myRune, myRune)
	}
}

// ByteCount shows how to get the amount of bytes in a string and using a for
// loop won't work if the bytes do not fall in the ASCII Table.
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
