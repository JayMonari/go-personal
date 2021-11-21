package diamond

import (
	"bytes"
	"fmt"
)

// Gen generates a diamond of uppercase ASCII characters. If b is not between
// the values of A-Z then an error is returned.
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", fmt.Errorf("b needs to be an uppercase letter %b passed in", b)
	}
	area := int((b-'A')*2 + 1)
	lines := fill(area)
	form(&lines)
	return string(bytes.Join(lines, []byte{'\n'})) + "\n", nil
}

// fill creates a slice of slices of bytes that is whitespace padded to the
// area given.
func fill(area int) [][]byte {
	lines := make([][]byte, area)
	for i, l := range lines {
		l = bytes.Repeat([]byte{' '}, area)
		lines[i] = l
	}
	return lines
}

// form modifies the slices of bytes to form a diamond of uppercase ASCII
// values according to the given length of lines.
func form(lines *[][]byte) {
	diamond := *lines
	b, mid, end := byte('A'), len(diamond)/2, len(diamond)-1
	for i := 0; i < mid; i, b = i+1, b+1 {
		diamond[i][mid-i] = b
		diamond[i][mid+i] = b
		diamond[end-i][mid-i] = b
		diamond[end-i][mid+i] = b
	}
	diamond[mid][0] = b
	diamond[mid][end] = b
}
