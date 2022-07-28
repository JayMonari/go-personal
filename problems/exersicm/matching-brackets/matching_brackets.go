package brackets

// pairParts contains all pairs.
var pairParts = [6]rune{'(', ')', '[', ']', '{', '}'}
// closeParts maps all closing parts to their open counterpart.
var closeParts = map[rune]rune{'}': '{', ']': '[', ')': '('}

// Bracket returns true if a line of text with (), [], {} are all properly
// matched.
func Bracket(text string) bool {
	openParts := []rune{}
	for _, part := range text {
		if !isPairPart(part) {
			continue
		}

		if open1, ok := closeParts[part]; ok {
			if len(openParts) == 0 || open1 != openParts[len(openParts)-1] {
				return false
			}

			// "pop" last part from slice
			openParts = openParts[:len(openParts)-1]
		} else {
			openParts = append(openParts, part)
		}
	}
	return len(openParts) == 0
}

// isPairPart returns whether the rune is in pairParts.
func isPairPart(p rune) bool {
	for _, pt := range pairParts {
		if p == pt {
			return true
		}
	}
	return false
}
