package raindrops

import (
	"strconv"
)

const testVersion = 2

func Convert(in int) string {
	retur := ""
	if in%3 == 0 {
		retur += "Pling"
	}
	if in%5 == 0 {
		retur += "Plang"
	}
	if in%7 == 0 {
		retur += "Plong"
	}
	if retur == "" {
		retur = strconv.Itoa(in)
	}
	return retur
}
