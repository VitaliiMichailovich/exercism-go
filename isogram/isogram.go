package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {
	m := map[rune]int{}
	for _, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) {
			m[r]++
			if m[r] > 1 {
				return false
			}
		}
	}
	return true
}