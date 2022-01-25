package parser_test

import (
	p "parser"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	var tcs = []struct {
		s   string
		tok p.Token
		lit string
	}{
		{s: ``, tok: p.EOF},
		{s: `#`, tok: p.ILLEGAL, lit: `#`},
		{s: ` `, tok: p.WS, lit: ` `},
		{s: "\t", tok: p.WS, lit: "\t"},
		{s: "\n", tok: p.WS, lit: "\n"},
		{s: `*`, tok: p.ASTERISK, lit: "*"},
		{s: `foo`, tok: p.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: p.IDENT, lit: `Zx12_3U_`},

		{s: `FROM`, tok: p.FROM, lit: `FROM`},
		{s: `SELECT`, tok: p.SELECT, lit: `SELECT`},
		{s: `from`, tok: p.FROM, lit: `from`},
		{s: `select`, tok: p.SELECT, lit: `select`},
	}
	for i, tt := range tcs {
		s := p.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: want=%q, got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: want=%q, got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
