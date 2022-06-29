package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var squares int
	for _, value := range cb[rank] {
		if value {
			squares++
		}
	}
	return squares
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}
	var squares int
	for _, value := range cb {
		if value[file-1] {
			squares++
		}
	}
	return squares
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var squares int
	for _, ranks := range cb {
		for range ranks {
			squares++
		}
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var squares int
	for _, ranks := range cb {
		for _, value := range ranks {
			if value {
				squares++
			}
		}
	}
	return squares
}
