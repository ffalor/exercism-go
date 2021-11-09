package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	occupied := 0

	for _, v := range cb[rank] {
		if v {
			occupied++
		}
	}

	return occupied
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	occupied := 0

	if file > len(cb) {
		return occupied
	}

	for _, v := range cb {
		if v[file-1] {
			occupied++
		}
	}

	return occupied
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for range v {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	occupied := 0

	for _, v := range cb {
		for _, v2 := range v {
			if v2 {
				occupied++
			}
		}
	}

	return occupied
}
