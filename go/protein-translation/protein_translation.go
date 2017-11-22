package protein

const testVersion = 1

var codons = [][2]string{
	{"AUG", "Methionine"},
	{"UUU", "Phenylalanine"},
	{"UUC", "Phenylalanine"},
	{"UUA", "Leucine"},
	{"UUG", "Leucine"},
	{"UCU", "Serine"},
	{"UCC", "Serine"},
	{"UCA", "Serine"},
	{"UCG", "Serine"},
	{"UAU", "Tyrosine"},
	{"UAC", "Tyrosine"},
	{"UGU", "Cysteine"},
	{"UGC", "Cysteine"},
	{"UGG", "Tryptophan"},
	{"UAA", "STOP"},
	{"UAG", "STOP"},
	{"UGA", "STOP"},
}

func FromCodon(in string) string {
	var out string
	for i := 0; i < len(codons); i++ {
		if codons[i][0] == in {
			return codons[i][1]
		}
	}
	return out
}

func FromRNA(in string) []string {
	var out []string
	for i := 0; i < len(in)/3; i++ {
		appendR := FromCodon(in[i*3 : i*3+3])
		if appendR == "STOP" {
			break
		} else {
			out = append(out, FromCodon(in[i*3:i*3+3]))
		}
	}
	return out
}