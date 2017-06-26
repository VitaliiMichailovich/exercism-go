package house

const testVersion = 1

var info = []struct {
	first  string
	second string
}{
	{"lay", "house that Jack built."},
	{"ate", "malt"},
	{"killed", "rat"},
	{"worried", "cat"},
	{"tossed", "dog"},
	{"milked", "cow with the crumpled horn"},
	{"kissed", "maiden all forlorn"},
	{"married", "man all tattered and torn"},
	{"woke", "priest all shaven and shorn"},
	{"kept", "rooster that crowed in the morn"},
	{"belonged to", "farmer sowing his corn"},
	{"", "horse and the hound and the horn"},
}

func Verse(in int) string {
	var out string
	for i := in - 1; i >= 0; i-- {
		if i == (in - 1) {
			out += "This is the " + info[i].second
			if in != 1 {
				out += "\n"
			}
		} else if i == 0 {
			out += "that " + info[i].first + " in the " + info[i].second
		} else {
			out += "that " + info[i].first + " the " + info[i].second + "\n"
		}
	}
	return out
}

func Song() string {
	var out string
	for i := 1; i <= len(info); i++ {
		if i != len(info) {
			out += Verse(i) + "\n\n"
		} else {
			out += Verse(i)
		}
	}
	return out
}