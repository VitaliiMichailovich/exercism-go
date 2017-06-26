package queenattack

import (
	"errors"
	"math"
	"strconv"
)

const testVersion = 2

func CanQueenAttack(w, b string) (bool, error) {
	var whiteCol, blackCol int
	var out bool
	var ok error
	var field []string = []string{"", "a", "b", "c", "d", "e", "f", "g", "h"}
	if len(w) == 0 || len(b) == 0 || len(w) > 2 || len(b) > 2 {
		return false, errors.New("Error0")
	}
	for i, id := range field {
		if id == string(w[0]) {
			whiteCol = i
		}
		if id == string(b[0]) {
			blackCol = i
		}
	}
	whiteRow, errWR := strconv.Atoi(string(w[1]))
	blackRow, errBR := strconv.Atoi(string(b[1]))
	if (whiteRow == blackRow && whiteCol == blackCol) ||
		errWR != nil || errBR != nil ||
		whiteRow > 8 || blackRow > 8 ||
		whiteRow == 0 || blackRow == 0 ||
		whiteCol > 8 || blackCol > 8 ||
		whiteCol == 0 || blackCol == 0 {
		return false, errors.New("Error1")
	}
	if math.Abs(float64(blackRow-whiteRow)) == math.Abs(float64(blackCol-whiteCol)) ||
		blackCol == whiteCol || blackRow == whiteRow {
		out = true
	} else {
		out = false
	}
	return out, ok
}
