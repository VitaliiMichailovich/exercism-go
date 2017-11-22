package accumulate

const testVersion = 1

func Accumulate(aa []string, bb func(string) string) []string {
	var retur []string
	for _, cc := range aa {
		retur = append(retur, bb(cc))
	}
	return retur
}