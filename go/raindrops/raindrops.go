package raindrops

import "fmt"

// Convert converts a number into a string that contains raindrop sounds
func Convert(number int) string {
	var sound string
	if number%3 == 0 {
		sound += "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = fmt.Sprintf("%v", number)
	}
	return sound
}
