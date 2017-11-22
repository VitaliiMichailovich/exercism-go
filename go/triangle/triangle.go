package triangle

import "math"

type Kind struct {
	Iso, NaT, Sca, Equ bool
}

var Iso = Kind{Iso: true}
var NaT = Kind{NaT: true}
var Sca = Kind{Sca: true}
var Equ = Kind{Equ: true}

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	out := Kind{}
	switch {
	case a <= 0 || b <= 0 || c <= 0 ||
		a+b < c || a+c < b || b+c < a ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		out.NaT = true
	case a == b && b == c && a > 0:
		out.Equ = true
	case a != b && b != c && c != a:
		out.Sca = true
	case (a == b && a != c) || (b == c && a != c) || (b != c && a == c):
		out.Iso = true
	}
	return out
}