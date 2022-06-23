package railfence

import "strings"

// Encode takes in text and a number of rails to encode the string into a
// ciphered message.
func Encode(text string, nRails int) string {
	p := processor{
		fence:  create(len(text), nRails),
		result: &strings.Builder{},
	}
	walkFence(p, text, nRails,
		func(p processor, row int, col *int, b byte) {
			p.fence[row][*col] = byte(b)
		})

	for _, line := range p.fence {
		for _, b := range line {
			if b != '.' {
				p.result.WriteByte(b)
			}
		}
	}
	return p.result.String()
}

// Decode takes in a ciphered text and with the number of rails to the fence,
// provides the decoded message.
func Decode(cipher string, nRails int) string {
	p := processor{
		fence:  create(len(cipher), nRails),
		result: &strings.Builder{},
	}
	walkFence(p, strings.Repeat("*", len(cipher)), nRails,
		func(p processor, row int, col *int, b byte) {
			p.fence[row][*col] = byte(b)
		})

	i := 0
	for row := 0; row < nRails; row++ {
		for col := 0; col < len(cipher); col++ {
			if p.fence[row][col] == '*' && i < len(cipher) {
				p.fence[row][col] = cipher[i]
				i++
			}
		}
	}
	walkFence(p, cipher, nRails,
		func(p processor, row int, col *int, b byte) {
			if f := p.fence[row][*col]; f != '*' {
				p.result.WriteByte(f)
				*col++
			}
		})
	return p.result.String()
}

// create initializes a matrix of bytes to filler '.' characters. sLen will
// produce the amount of rows in the matrix and nRails will produces the amount
// of columns.
func create(sLen, nRails int) (fence [][]byte) {
	fence = make([][]byte, nRails)
	for i := 0; i < nRails; i++ {
		fence[i] = make([]byte, sLen)
		for j := 0; j < sLen; j++ {
			fence[i][j] = '.'
		}
	}
	return fence
}

// walkFence will follow the rail fence and do work depending on the passed in
// logic. The logic passed in should change the processor.fence or
// processor.result
func walkFence(p processor, s string, nRails int,
	logic func(processor, int, *int, byte)) {

	goDown := false
	row := 0
	for col, b := range s {
		logic(p, row, &col, byte(b))
		switch row {
		case 0:
			goDown = true
		case nRails - 1:
			goDown = false
		}
		if goDown {
			row++
		} else {
			row--
		}
	}
}

// processor builds the result from the fence
type processor struct {
	fence  [][]byte
	result *strings.Builder
}
