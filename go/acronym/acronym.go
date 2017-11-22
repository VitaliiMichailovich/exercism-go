package acronym

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Abbreviate(in string)string {
	out := in[0:1]
	for i, a := range in {
		if i == 0 {
			continue
		}
		aa := strings.ToUpper(string(a))
		prev := rune(in[i-1])
		curr := rune(in[i])
		if (unicode.IsLetter(curr) && !unicode.IsLetter(prev)) || (unicode.IsUpper(curr) && unicode.IsLower(prev)) {
			out += aa
		}
	}
	return out
}