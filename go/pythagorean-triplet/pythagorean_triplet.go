package pythagorean

import (
	"math"
)

type Triplet struct {
	a, b, c int
}

const testVersion = 1

func Range(min, max int) []Triplet {
	var out []Triplet
	for i := min; i <= max-2; i++ {
		aa := int(math.Pow(float64(i), 2.0))
		for ii := i + 1; ii <= max-1; ii++ {
			bb := int(math.Pow(float64(ii), 2.0))
			for iii := ii + 1; iii <= max; iii++ {
				cc := int(math.Pow(float64(iii), 2.0))
				if (aa + bb) == cc {
					out = append(out, Triplet{i, ii, iii})
				}
				if (aa + bb) < cc {
					break
				}
			}
		}
	}
	return out
}

func Sum(p int) []Triplet {
	var out []Triplet
	for i := 1; i > 0; i++ {
		aa := i * i
		for ii := i + 1; ii > 0; ii++ {
			bb := ii * ii
			for iii := ii + 1; iii > 0; iii++ {
				cc := iii * iii
				if (aa+bb) == cc && i+ii+iii == p {
					out = append(out, Triplet{i, ii, iii})
				}
				if i+ii+iii > p || aa+bb < cc {
					break
				}
			}
			if i+ii*2+1 > p {
				break
			}
		}
		if i*3+3 > p {
			break
		}
	}
	return out
}