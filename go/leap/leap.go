package leap

// Given a year, return true if it is a leap year and false if it is not
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
