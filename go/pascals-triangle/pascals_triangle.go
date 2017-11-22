package pascal

const testVersion = 1

func Triangle(in int) [][]int {
	var out [][]int = make([][]int, 0)
	for row := 0; row < in; row++ {
		var rowa []int = make([]int, 0)
		if row == 0 {
			rowa = append(rowa, 1)
		} else if row == 1 {
			rowa = append(rowa, 1)
			rowa = append(rowa, 1)
		} else {
			for i := 0; i <= row; i++ {
				if i == 0 || i == row {
					rowa = append(rowa, 1)
				} else {
					rowa = append(rowa, out[row-1][i-1]+out[row-1][i])
				}
			}
		}
		out = append(out, rowa)
	}
	return out
}