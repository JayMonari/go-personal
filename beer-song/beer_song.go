package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verse returns a verse from the Beer Song. If the verse is out of range of
// 0-99 an error is returned.
func Verse(line int) (string, error) {
	switch {
	case line > 99 || line < 0:
		return "", errors.New("verse does not exist")
	case line == 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	case line == 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case line == 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	default:
		return fmt.Sprintf(`%d bottles of beer on the wall, %d bottles of beer.
Take one down and pass it around, %d bottles of beer on the wall.
`, line, line, line-1), nil
	}
}

// Verses returns the verses specified by start and end. If end > start or
// start > 99 or end < 0 then an error is returned.
func Verses(start, end int) (string, error) {
	if start > 99 || end < 0 || start < end {
		return "", errors.New("Not within range")
	}
	var verses []string
	for i := start; i >= end; i-- {
		verse, _ := Verse(i)
		verses = append(verses, verse)
	}
	return strings.Join(verses, "\n") + "\n", nil
}

// Song returns the entire Beer song from 99 to 0.
func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}
