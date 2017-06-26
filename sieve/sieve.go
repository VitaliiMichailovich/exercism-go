package sieve

const testVersion = 1

func Sieve(in int) []int {
	var out []int
	tuo := make([]int, in-1)
	for i := range tuo {
		tuo[i] = i + 2
	}

	for i := 0; i < len(tuo); i++ {
		if tuo[i] != 0 || i == 0 {
			out = append(out, tuo[i])
		}
		for ii := i + 1; ii < len(tuo); ii++ {
			if tuo[i] != 0 {
				if tuo[ii]%tuo[i] == 0 || tuo[ii] == 0 {
					tuo[ii] = 0
				}
			}
		}
	}
	return out
}