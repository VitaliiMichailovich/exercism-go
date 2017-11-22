package summultiples

const testVersion = 1

func SumMultiples(limit int, divisors ...int) int {
	var out int
	for i := 0; i < limit; i++ {
		div := false
		for ii := 0; ii < len(divisors); ii++ {
			if i%divisors[ii] == 0 {
				div = true
			}
		}
		if div == true {
			out += i
		}
	}
	return out
}