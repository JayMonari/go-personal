package encode

import (
	"bytes"
	"strconv"
	"strings"
)

func RunLengthEncode(s string) string {
	enc := strings.Builder{}
	for len(s) > 0 {
		char := s[0]
		untrimLen := len(s)
		s = strings.TrimLeft(s, string(char))
		if cnt := untrimLen - len(s); cnt > 1 {
			enc.WriteString(strconv.Itoa(cnt))
		}
		enc.WriteByte(char)
	}
	return enc.String()
}

func RunLengthDecode(s string) string {
	dec := strings.Builder{}
	nextIdx := func(i int) int {
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			i++
		}
		return i
	}

	for i :=  0; i < len(s); i++ {
		j := nextIdx(i)
		count, err := strconv.Atoi(s[i:j])
		i = j
		if err != nil {
			count = 1
		}

		dec.Write(bytes.Repeat([]byte{s[j]}, count))
	}
	return dec.String()
}
