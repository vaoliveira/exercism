package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(\~|\*|\=|\-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`".*(P|p)(A|a)(S|s){2}(W|w)(O|o)(R|r)(D|d).*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User(\s+)(\w+)`)
	for i, line := range lines {
		if re.MatchString(line) {
			user := re.FindStringSubmatch(line)
			lines[i] = fmt.Sprintf("[USR] %s %s", user[1], line)
		}
	}
	return lines
}
