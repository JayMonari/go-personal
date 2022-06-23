package queenattack

import (
	"errors"
	"regexp"
)

// CanQueenAttack checks if two queens on a chessboard can attack each other.
// strings must be 2 characters satisfying the regexp [a-h]{1}[1-8]{1} else an
// error is returned.
func CanQueenAttack(w, b string) (bool, error) {
	if invalidPieces(w, b) {
		return false, errors.New("invalid pieces")
	}

	wRow, wCol := w[0], w[1]
	bRow, bCol := b[0], b[1]
	return wRow == bRow ||
			wCol == bCol ||
			wRow+wCol == bRow+bCol ||
			wRow-wCol == bRow-bCol,
		nil
}

// invalidPieces checks whether white and black are on the board.
func invalidPieces(w, b string) bool {
	wOK, _ := regexp.MatchString("[a-h]{1}[1-8]{1}", w)
	bOK, _ := regexp.MatchString("[a-h]{1}[1-8]{1}", b)
	if w == b || !wOK || !bOK {
		return true
	}
	return false
}
