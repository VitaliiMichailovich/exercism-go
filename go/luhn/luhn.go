package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 2

func Valid(in string) bool {
	out := false
	cou := 0
	sum := 0
	var err error
	in = strings.Replace(in, " ", "", -1)
	if len(in) <= 1 {
		return out
	}
	cardId := make([]int, len(in), len(in)*2)
	for i := len(in) - 1; i >= 0; i-- {
		if unicode.IsNumber(rune(in[i])) {
			cou++
		}
		check := 0
		if len(in)%2 == 0 {
			check = i
		} else {
			check = i + 1
		}
		if check%2 == 0 {
			s, err := strconv.Atoi(in[i : i+1])
			if err == nil {
				if s*2 > 9 {
					cardId[i] = s*2 - 9
				} else {
					cardId[i] = s * 2
				}
			}
		} else {
			cardId[i], err = strconv.Atoi(in[i : i+1])
			if err != nil {
				return out
			}
		}
		sum += cardId[i]
	}
	if cou != len(in) {
		return out
	}
	if sum%10 == 0 {
		out = true
	}
	return out
}