package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"unicode"
)

const testVersion = 2 // same as targetTestVersion

func Hey(in string) string {
	out := "Whatever."
	inz := strings.TrimSpace(in)
	if len(inz) == 0 {
		out = "Fine. Be that way!"
		return out
	}
	lenUpCase := 0
	lenLetters := 0
	for _, ina := range inz {
		if unicode.IsLetter(ina) {
			lenLetters++
			if strings.ToUpper(string(ina)) == string(ina) {
				lenUpCase++
			}
		}
	}
	if inz[len(inz)-1:len(inz)-0] == "?" {
		out = "Sure."
	}
	if lenUpCase == lenLetters && lenLetters != 0 {
		out = "Whoa, chill out!"
	}
	return out
}