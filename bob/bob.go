// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isQuestion(s string) bool {
	return strings.HasSuffix(strings.TrimSpace(s), "?")
}

func isYelling(s string) bool {
	return strings.ToUpper(s) == s && strings.IndexFunc(s, unicode.IsLetter) >= 0
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	switch {
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isYelling(remark):
		return "Whoa, chill out!"
	case isEmpty(remark):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
