package twofer

import "fmt"

// ShareWith receives a string "name" to return the string "One for <name>, one for me"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
