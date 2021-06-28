package scrabble

import "strings"

var values = map[int]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

func Score(word string) int {
	score := 0
	word = strings.ToUpper(word)

	for _, c := range word {
		for k, v := range values {
			if strings.Contains(v, string(c)) {
				score += k
			}
		}
	}

	return score
}
