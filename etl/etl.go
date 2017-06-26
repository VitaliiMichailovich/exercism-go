package etl

import (
	"strings"
)

const testVersion = 1

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for i, mapa := range in {
		for _, mapi := range mapa {
			out[strings.ToLower(mapi)] = i
		}
	}

	return out
}