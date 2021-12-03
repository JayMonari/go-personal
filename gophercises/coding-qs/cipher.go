package main

import "unicode"

func cipher(r rune, delta int) rune {
	if unicode.IsLower(r) {
		return rotate(r, 'a', delta)
	} else if unicode.IsUpper(r) {
		return rotate(r, 'A', delta)
	}
	return r
}

func rotate(r rune, base, delta int) rune {
	tmp := ((int(r) - base) + delta) % 26
	return rune(tmp + base)
}
