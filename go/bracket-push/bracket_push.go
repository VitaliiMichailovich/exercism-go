package brackets

import (
	"strconv"
)

const testVersion = 4

var bracketOpen = []string{"[", "{", "("}
var bracketClose = []string{"]", "}", ")"}

func Bracket(in string) (bool, error) {
	var out bool
	var err error
	var check string

	for i := 0; i < len(in); i++ {
		for ii := 0; ii < (len(bracketOpen)+len(bracketClose))/2; ii++ {
			if in[i:i+1] == bracketOpen[ii] {
				check += strconv.Itoa(ii)
				break
			} else if in[i:i+1] == bracketClose[ii] {
				if len(check) > 0 && check[len(check)-1:] == strconv.Itoa(ii) {
					check = check[0 : len(check)-1]
					break
				} else {
					return false, err
				}
			}
		}
	}
	if len(check) == 0 {
		out = true
	}
	return out, err
}