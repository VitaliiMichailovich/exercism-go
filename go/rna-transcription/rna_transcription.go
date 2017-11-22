package strand

const testVersion = 3

var repl = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

func ToRNA(in string) string {
	var out string
	for i := range in {
		for ii := range repl {
			if in[i:i+1] == ii {
				out += repl[ii]
			}
		}
	}
	return out
}