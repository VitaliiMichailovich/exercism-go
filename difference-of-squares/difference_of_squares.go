package diffsquares

import "math"

const testVersion = 1

func SquareOfSums(in int) int {
	var out int
	for i := 1; i <= in; i++ {
		out += i
	}
	return int(math.Pow(float64(out), 2))
}

func SumOfSquares(in int) int {
	var out int
	for i := 1; i <= in; i++ {
		out += int(math.Pow(float64(i), 2))
	}
	return out
}

func Difference(in int) int {
	return SquareOfSums(in) - SumOfSquares(in)
}
