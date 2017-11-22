package say

const testVersion = 1

var llion = []string{"hundred", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion"}
var ty = []string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var teen = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var ten = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func Say(in uint64) string {
	var out string
	if in == 0 {
		return ten[0]
	}
	var llionI int
	llionIIn := in
	for i := 1; i >= 0; i++ {
		if llionIIn < 1000 {
			break
		}
		llionIIn = llionIIn / 1000
		llionI++
	}
	for i := llionI; i >= 0; i-- {
		var dev uint64 = 1
		for ii := i; ii >= 1; ii-- {
			dev = dev * 1000
		}
		var inn int = int(in / dev)
		sayThree := SayThree(inn)
		if sayThree != "" {
			out += sayThree
			if dev > 1 {
				out += " " + llion[i]
			}
			in = in - uint64(inn)*dev
			if in > 0 {
				out += " "
			}
		}
	}
	return out
}

func SayThree(in int) string {
	var out string
	if in >= 100 {
		out += ten[in/100] + " " + llion[0]
		in = in - (in/100)*100
		if in > 0 {
			out += " "
		}
	}
	if in >= 20 {
		out += ty[in/10]
		in = in - (in/10)*10
		if in > 0 {
			out += "-"
		}
	}
	if in < 20 && in > 10 {
		out += teen[in-10]
	}
	if in > 0 && in < 10 {
		out += ten[in]
	}
	return out
}
