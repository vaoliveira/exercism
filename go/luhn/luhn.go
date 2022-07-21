package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	var sum int
	id = strings.ReplaceAll(id, " ", "")
	if len(id) <= 1 {
		return false
	}
	for i := len(id) - 1; i >= 0; i-- {
		num, err := strconv.Atoi(string(id[i]))
		/* if err != nil {
			return false
		} else if (len(id)-i)%2 == 0 && num*2 > 9 {
			sum += (num*2 - 9)
		} else if (len(id)-i)%2 == 0 {
			sum += num * 2
		} else {
			sum += num
		} */
		switch {
		case err != nil:
			return false
		case (len(id)-i)%2 == 0 && num*2 > 9:
			sum += (num*2 - 9)
		case (len(id)-i)%2 == 0:
			sum += num * 2
		default:
			sum += num
		}
	}
	return sum%10 == 0
}
