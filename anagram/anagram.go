package anagram

import (
	"reflect"
	"strings"
)

const testVersion = 1

func Detect(subject string, candidates []string) []string {
	var out []string
	sMap := make(map[string]int)
	subject = strings.ToLower(subject)
	for i := 0; i < len(subject); i++ {
		sMap[subject[i:i+1]]++
	}
	for i := 0; i < len(candidates); i++ {
		candidates[i] = strings.ToLower(candidates[i])
		if len(subject) == len(candidates[i]) && subject != candidates[i] {
			cMap := make(map[string]int)
			for ii := 0; ii < len(candidates[i]); ii++ {
				cMap[candidates[i][ii:ii+1]]++
			}
			if reflect.DeepEqual(sMap, cMap) {
				out = append(out, candidates[i])
			}
		}
	}

	return out
}