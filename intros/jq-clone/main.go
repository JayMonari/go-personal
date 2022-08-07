package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/profile"
)

type JSONReader struct {
	read                []byte
	expectString_buffer bytes.Buffer
}

func (jr *JSONReader) readByte(r *bufio.Reader) (byte, error) {
	c, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	jr.read = append(jr.read, c)
	return c, nil
}

func (jr *JSONReader) reset() { jr.read = nil }

func (jr *JSONReader) extractDataFromJSONPath(r *bufio.Reader, path []string) (any, error) {
	if len(path) == 0 {
		return nil, nil
	}
	if err := jr.eatWhitespace(r); err != nil {
		return nil, err
	}
	b, err := jr.readByte(r)
	if err != nil {
		return nil, err
	}
	// Make sure we're actually going into an object
	if b != '{' {
		return nil, fmt.Errorf("expected opening curly brace, got: %q", string(b))
	}
	var result any
	for i := 0; ; i++ {
		if err = jr.eatWhitespace(r); err != nil {
			return nil, err
		}
		bs, err := r.Peek(1)
		if err != nil {
			return nil, err
		}
		b := bs[0]
		if b == '}' {
			r.Discard(1)
			break
		}

		if i > 0 {
			if b != ',' {
				return nil, fmt.Errorf("expected comma between key-value pairs, got: %q", string(b))
			}
			r.Discard(1)
		}
		if err := jr.eatWhitespace(r); err != nil {
			return nil, err
		}
		s, err := jr.expectString(r)
		if err != nil {
			return nil, err
		}
		if err := jr.eatWhitespace(r); err != nil {
			return nil, err
		}
		b, err = jr.readByte(r)
		if err != nil {
			return nil, err
		}
		if b != ':' {
			return nil, fmt.Errorf("expected colon, got: %q", string(b))
		}
		if err := jr.eatWhitespace(r); err != nil {
			return nil, err
		}

		if path[0] != s {
			err = jr.eatValue(r)
			if err != nil {
				return nil, err
			}
			continue
		}
		result, err = jr.expectValue(r, path[1:])
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (jr *JSONReader) extractArrayDataFromJSONPath(r *bufio.Reader, path []string) (any, error) {
	n, err := strconv.Atoi(string(path[0]))
	if err != nil {
		return nil, err
	}
	b, err := jr.readByte(r)
	if err != nil {
		return nil, err
	}

	if b != '[' {
		return nil, fmt.Errorf("expected opening bracket, got: %q", string(b))
	}
	var result any
	for i := 0; ; i++ {
		if err := jr.eatWhitespace(r); err != nil {
			return nil, err
		}
		bs, err := r.Peek(1)
		if err != nil {
			return nil, err
		}
		b := bs[0]
		if b == ']' {
			r.Discard(1)
			break
		}

		if i > 0 {
			if b != ',' {
				return nil, fmt.Errorf("expected comma between key-value pairs, got %q", string(b))
			}
			r.Discard(1)
		}
		if err := jr.eatWhitespace(r); err != nil {
			return nil, err
		}
		if i != n {
			if err = jr.eatValue(r); err != nil {
				return nil, err
			}
			continue
		}
		result, err = jr.expectValue(r, path[1:])
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (jr *JSONReader) eatValue(r *bufio.Reader) error {
	var stack []byte
	inString := false
	var prev byte
	if err := jr.eatWhitespace(r); err != nil {
		return err
	}
	ok, _, err := jr.tryScalar(r)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	// Must be array or object
	length := 0
	first := true
	var bs []byte
	for first || len(stack) > 0 {
		length++
		first = false
		for {
			bs, err = r.Peek(length)
			if err == bufio.ErrBufferFull {
				_, err = r.Discard(length - 1)
				if err != nil {
					return err
				}
				length = 1
				continue
			}
			if err != nil {
				return err
			}
			break
		}
		b := bs[length-1]

		if inString {
			if b == '"' && prev != '\\' {
				inString = false
			}
			// Two backslashes in a row `\\` cancel each other out e.g. `\\n`
			if b == '\\' && prev == '\\' {
				prev = 0
			} else {
				prev = b
			}
			continue
		}
		switch b {
		case '[':
			stack = append(stack, b)
		case ']':
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != '[' {
				return fmt.Errorf("unexpected end of array: %q", string(c))
			}
		case '{':
			stack = append(stack, b)
		case '}':
			c := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != '{' {
				return fmt.Errorf("unexpected end of object: %q", string(c))
			}
		case '"':
			inString = true // Closing quote handled above 👆
		}

		prev = b
	}
	_, err = r.Discard(length)
	return err
}

func (jr *JSONReader) eatWhitespace(r *bufio.Reader) error {
	for {
		bs, err := r.Peek(1)
		if err != nil {
			return err
		}
		if b := bs[0]; !(b == ' ' || b == '\n' || b == '\t' || b == '\r') {
			return nil
		}
		r.Discard(1)
	}
}

func (jr *JSONReader) expectString(r *bufio.Reader) (string, error) {
	jr.expectString_buffer.Reset()
	if err := jr.eatWhitespace(r); err != nil {
		return "", err
	}
	b, err := jr.readByte(r)
	if err != nil {
		return "", err
	}
	if b != '"' {
		return "", fmt.Errorf("expected double quote to start string, got: %q",
			string(b))
	}

	var prev byte
	for {
		b, err = jr.readByte(r)
		if err != nil {
			return "", err
		}
		if b == '\\' && prev == '\\' {
			prev = 0 // Just skip
			continue
		} else if b == '"' {
			if prev == '\\' {
				jr.expectString_buffer.Bytes()[jr.expectString_buffer.Len()-1] = '"'
			} else {
				break // Ending double quote
			}
		}
		jr.expectString_buffer.WriteByte(b)
		prev = b
	}
	return jr.expectString_buffer.String(), nil
}

func (jr *JSONReader) expectValue(r *bufio.Reader, path []string) (any, error) {
	bs, err := r.Peek(1)
	if err != nil {
		return nil, err
	}
	b := bs[0]
	switch b {
	case '{':
		return jr.extractDataFromJSONPath(r, path)
	case '[':
		return jr.extractArrayDataFromJSONPath(r, path)
	}
	if len(path) != 0 {
		// Reached end of this object but more path to explore.
		return nil, nil
	}
	ok, val, err := jr.tryScalar(r)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("expected scalar, got: %q", string(b))
	}
	return val, err
}

func (jr *JSONReader) expectIdentifier(r *bufio.Reader, ident string, v any) (any, error) {
	var bs []byte
	for i := 0; i < len(ident); i++ {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	if s := string(bs); s != ident {
		return nil, fmt.Errorf("unknown value: %q", string(bs))
	}
	return v, nil
}

func (jr *JSONReader) tryNumber(r *bufio.Reader) (bool, any, error) {
	var num []byte
	for {
		bs, err := r.Peek(1)
		if err != nil {
			return false, nil, err
		}
		b := bs[0]

		if !((b >= '0' && b <= '9') || b == 'e' || b == '-') { // Number character
			break
		}
		num = append(num, b)
		r.Discard(1)
	}
	if len(num) == 0 {
		return false, nil, nil
	}
	var n float64
	err := json.Unmarshal(num, &n)
	return true, n, err
}

func (jr *JSONReader) tryScalar(r *bufio.Reader) (bool, any, error) {
	bs, err := r.Peek(1)
	if err != nil {
		return false, nil, err
	}

	switch b := bs[0]; b {
	case '"':
		val, err := jr.expectString(r)
		return true, string(val), err
	case 't':
		val, err := jr.expectIdentifier(r, "true", true)
		return true, val, err
	case 'f':
		val, err := jr.expectIdentifier(r, "false", false)
		return true, val, err
	case 'n':
		val, err := jr.expectIdentifier(r, "null", nil)
		return true, val, err
	}
	return jr.tryNumber(r)
}

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	path := strings.Split(os.Args[1], ".")
	if path[0] == "" {
		path = path[1:]
	}

	b, enc := bufio.NewReader(os.Stdin), json.NewEncoder(os.Stdout)
	var jr JSONReader
	var val any
	var err error
	for {
		jr.reset()
		val, err = jr.extractDataFromJSONPath(b, path)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Read", string(jr.read))
			log.Fatalln(err)
		}
		err = enc.Encode(val)
	}
}
