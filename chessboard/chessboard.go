package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains eight Ranks, accessed with values from 'A' to 'H'
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) int {
	cnt := 0
	for _, occupied := range cb[rank] {
		if occupied {
			cnt++
		}
	}
	return cnt
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file > 9 {
		return 0
	}

	cnt := 0
	for _, rank := range cb {
		if rank[file-1] {
			cnt++
		}
	}
	return cnt
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	total := 0
	for _, rank := range cb {
		for range rank {
			total++
		}
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	cnt := 0
	for _, rank := range cb {
		for _, occupied := range rank {
			if occupied {
				cnt++
			}
		}
	}
	return cnt
}
