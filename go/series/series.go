package slice

const testVersion = 1

func All(n int, s string) []string {
	var out []string

	if n > len(s) {
		out = nil
	} else {
		var count int = len(s) - n + 1
		for i := 0; i < count; i++ {
			out = append(out, s[i:n+i])
		}
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}
