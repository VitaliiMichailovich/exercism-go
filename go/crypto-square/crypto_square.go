package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion = 2

func Encode(in string) string {
	var out, norm string
	for i := 0; i < len(in); i++ {
		if unicode.IsNumber(rune(in[i])) || unicode.IsLetter(rune(in[i])) {
			norm += strings.ToLower(in[i : i+1])
		}
	}
	var lenNorm int
	lenN := math.Sqrt(float64(len(norm)))
	if lenN-float64(int(lenN)) == 0 {
		lenNorm = int(lenN)
	} else {
		lenNorm = int(lenN) + 1
	}

	for i := 0; i < lenNorm; i++ {
		for ii := 0; ii < lenNorm; ii++ {
			if (ii*lenNorm + i + 1) > len(norm) {
				break
			}
			out += norm[ii*lenNorm+i : ii*lenNorm+i+1]
		}
		if (i + 1) != lenNorm {
			out += " "
		}
	}

	return out
}