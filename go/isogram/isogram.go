package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	count := map[byte]int{}
	for i := 0; i < len(word); i++ {
		if word[i] == '-' || word[i] == ' ' {
			continue
		}
		if _, exists := count[word[i]]; exists {
			return false
		} else {
			count[word[i]] = 1
		}
	}
	return true
}
