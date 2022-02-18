package xno

import (
	"strconv"
	"strings"
)

type Row [3]string
type Board []Row

func (b Board) String() string {
	var sb strings.Builder
	for i, r := range b {
		sb.WriteString(" | ")
		for j, v := range r {
			if v == "" {
				sb.WriteString(strconv.Itoa(1 + j + (i * 3)))
			} else {
				sb.WriteString(v)
			}
			sb.WriteString(" | ")
		}
		sb.WriteByte('\n')
		if i == len(b)-1 {
			break
		}
		sb.WriteString(strings.Repeat("-", 15))
		sb.WriteByte('\n')
	}
	return sb.String()
}
