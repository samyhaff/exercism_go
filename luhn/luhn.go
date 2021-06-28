package luhn

import (
	"strconv"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Valid(s string) bool {
	s = Reverse(strings.ReplaceAll(s, " ", ""))
	if len(s) <= 1 {
		return false
	}

	digits := strings.Split(s, "")

	sum := 0
	for i, digit := range digits {
		digit, err := strconv.Atoi(digit)
		if err != nil {
			return false
		}
		if i%2 != 0 {
			digit = (digit * 2)
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
