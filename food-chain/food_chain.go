package foodchain

const testVersion = 3

var verse = []struct {
	thing     string
	action    string
	addition2 string
}{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", ""},
	{"spider", "It wriggled and jiggled and tickled inside her.", "that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!", ""},
	{"cat", "Imagine that, to swallow a cat!", ""},
	{"dog", "What a hog, to swallow a dog!", ""},
	{"goat", "Just opened her throat and swallowed a goat!", ""},
	{"cow", "I don't know how she swallowed a cow!", ""},
	{"horse", "She's dead, of course!", ""},
}

func Verse(in int) string {
	var out string
	out += "I know an old lady who swallowed a " + verse[in-1].thing + ".\n"
	if in >= 2 && in != 8 {
		out += verse[in-1].action + "\n"
		for i := in; i >= 2; i-- {
			if i == 3 {
				out += "She swallowed the " +
					verse[i-1].thing + " to catch the " +
					verse[i-2].thing + " " +
					verse[i-2].addition2 + ".\n"
			} else {
				out += "She swallowed the " +
					verse[i-1].thing + " to catch the " +
					verse[i-2].thing + ".\n"
			}
		}
	}
	if in != 8 {
		out += verse[0].action
	} else {
		out += verse[7].action
	}
	return out
}

func Verses(from, to int) string {
	var out string
	for i := from; i <= to; i++ {
		out += Verse(i)
		if i != to {
			out += "\n\n"
		}
	}
	return out
}

func Song() string {
	out := Verses(1, 8)
	return out
}