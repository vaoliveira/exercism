package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var message string = fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
	return message
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var starLine string = ""
	for i := 0; i < numStarsPerLine; i++ {
		starLine = fmt.Sprintf("%s%s", starLine, "*")
	}
	var message string = fmt.Sprintf("%s\n%s\n%s", starLine, welcomeMsg, starLine)
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}
