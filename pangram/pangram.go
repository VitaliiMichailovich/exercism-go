package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(in string) bool {
	arrr := []string{"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z"}
	ins := strings.ToLower(in)
	count := 26
	for _, inp := range arrr {
		if strings.Contains(ins, inp) {
			count --
		}
	}
	if count == 0 {
		return true
	} else {
		return false
	}
}