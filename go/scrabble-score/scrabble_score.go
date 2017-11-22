package scrabble

import (
	"strings"
)

const testVersion = 4

var (
	score = map[int][]string{
		1:  {"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"},
		2:  {"d", "g"},
		3:  {"b", "c", "m", "p"},
		4:  {"f", "h", "v", "m", "y"},
		5:  {"k"},
		8:  {"j", "x"},
		10: {"q", "z"},
	}
)

func Score(in string) int {
	var out int
	scoreMap := map[string]int{}
	for i, mapa := range score {
		for _, sMap := range mapa {
			scoreMap[strings.ToLower(sMap)] = i
		}
	}
	for _, inR := range in {
		out += scoreMap[strings.ToLower(string(inR))]
	}
	return out
}
