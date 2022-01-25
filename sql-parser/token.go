package parser

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS // whitespace

	// Literals
	IDENT // fields, table_name

	// Misc characters
	ASTERISK // *
	COMMA // ,

	// Keywords
	SELECT
	FROM
)

var eof = rune(0)
