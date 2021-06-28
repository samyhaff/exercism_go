package isogram

import "strings"

func IsIsogram(s string) bool {
	count := make(map[string](int))

	for _, c := range s {
		c := strings.ToUpper(string(c))
		if c != "-" && c != " " {
			count[c]++
			if count[c] > 1 {
				return false
			}
		}
	}
	return true
}
