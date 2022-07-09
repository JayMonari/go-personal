package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	for _, tc := range map[string]struct{ in, want string }{
		"letters":            {"Hello Gophers", "srehpoG olleH"},
		"space":              {" ", " "},
		"special and digits": {"!12345", "54321!"},
	} {
		if rev, _ := Reverse(tc.in); rev != tc.want {
			t.Fatalf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

// go test -fuzz=Fuzz
// go test -run=FuzzReverse/f7d774048ada30d0b90795b843471413d1506dcbc74c4de673e8c919d
func FuzzReverse(f *testing.F) {
	tt := []string{"Hello Gophers", " ", "!12345"}
	for _, tc := range tt {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		dblRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != dblRev {
			t.Fatalf("before: %q, after: %q", orig, dblRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Fatalf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
