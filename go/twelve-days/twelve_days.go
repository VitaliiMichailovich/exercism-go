package twelve

const testVersion = 1

var daySmth = []struct {
	day  string
	smth string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}

func Verse(day int) string {
	var verseSmth string
	for i := day - 1; i >= 0; i-- {
		if i == 0 && day != 1 {
			verseSmth += "and " + daySmth[i].smth + "."
		} else if i == 0 && day == 1 {
			verseSmth += daySmth[i].smth + "."
		} else {
			verseSmth += daySmth[i].smth + ", "
		}
	}
	verse := "On the " + daySmth[day-1].day + " day of Christmas my true love gave to me, " + verseSmth
	return verse
}