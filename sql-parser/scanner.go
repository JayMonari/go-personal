package parser

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

// Scanner represents a lexical scanner.
type Scanner struct{ r *bufio.Reader }

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner { return &Scanner{r: bufio.NewReader(r)} }

// Scan returns the next toekn and literal value.
func (s *Scanner) Scan() (Token, string) {
	r := s.read()
	switch {
	case unicode.IsSpace(r):
		s.unread()
		return s.scanWhitespace()
	case unicode.IsLetter(r):
		s.unread()
		return s.scanIdent()
	}

	switch r {
	case eof:
		return EOF, ""
	case '*':
		return ASTERISK, string(r)
	case ',':
		return COMMA, string(r)
	default:
		return ILLEGAL, string(r)
	}
}

// scanWhitespace consumes the current rune and all contiguous whitespace
func (s *Scanner) scanWhitespace() (Token, string) {
	var sb strings.Builder
	sb.WriteRune(s.read())
	for {
		if r := s.read(); r == eof {
			break
		} else if !unicode.IsSpace(r) {
			s.unread()
			break
		} else {
			sb.WriteRune(r)
		}
	}
	return WS, sb.String()
}

func (s *Scanner) scanIdent() (Token, string) {
	var sb strings.Builder
	sb.WriteRune(s.read())
	for {
		if r := s.read(); r == eof {
			break
		} else if !unicode.In(r, unicode.N, unicode.L) && r != '_' {
			s.unread()
			break
		} else {
			sb.WriteRune(r)
		}
	}
	switch strings.ToUpper(sb.String()) {
	case "SELECT":
		return SELECT, sb.String()
	case "FROM":
		return FROM, sb.String()
	default:
		return IDENT, sb.String()
	}
}

// read reads the next rune from the buffered reader. Returns eof if an error
// occurs or io.EOF is returned.
func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return r
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }
