package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	count := map[rune]struct{}{}
	for _, value := range word {
		if value == '-' || value == ' ' {
			continue
		}
		if _, exists := count[value]; exists {
			return false
		}
		count[value] = struct{}{}
	}
	return true
}
