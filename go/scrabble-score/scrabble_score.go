package scrabble

import "unicode"

func Score(word string) int {
	var points int
	for i := 0; i < len(word); i++ {
		points += Points(word[i])
	}
	return points
}

func Points(letter byte) int {
	letter = byte(unicode.ToLower(rune(letter)))
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}
