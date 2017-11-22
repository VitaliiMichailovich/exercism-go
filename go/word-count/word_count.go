package wordcount

import (
	"strings"
)

const testVersion = 3

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	out := make(Frequency)
	phrase = strings.ToLower(phrase)
	repl := strings.NewReplacer(",", " ",
		".", " ",
		"!", "",
		"@", "",
		"#", "",
		"$", "",
		"%", "",
		"^", "",
		"&", "",
		"*", "",
		":", "")
	phrase = repl.Replace(phrase)
	phrase = strings.Replace(phrase, "  ", " ", -1)
	phrase = strings.TrimPrefix(phrase, " ")
	phrase = strings.TrimSuffix(phrase, " ")
	word := strings.Split(phrase, " ")
	for i := 0; i < len(word); i++ {
		w := strings.TrimSpace(word[i])
		w = strings.TrimPrefix(w, "'")
		w = strings.TrimSuffix(w, "'")
		out[w]++
	}
	return out
}