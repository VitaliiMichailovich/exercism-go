package secret

import (
	"strconv"
)

const testVersion = 1

func Handshake(in uint) []string {
	var out []string = make([]string, 0)
	var inBinary string
	binaryIn := string(strconv.FormatInt(int64(in), 2))
	if len(binaryIn) > 4 {
		inBinary = binaryIn[len(binaryIn)-4:]
	} else {
		inBinary = binaryIn
	}
	if len(inBinary) == 4 {
		if string(inBinary[0]) == "1" {
			out = append([]string{"jump"}, out...)

		}
		inBinary = inBinary[1:]
	}
	if len(inBinary) == 3 {
		if string(inBinary[0]) == "1" {
			out = append([]string{"close your eyes"}, out...)

		}
		inBinary = inBinary[1:]
	}
	if len(inBinary) == 2 {
		if string(inBinary[0]) == "1" {
			out = append([]string{"double blink"}, out...)

		}
		inBinary = inBinary[1:]
	}
	if len(inBinary) == 1 {
		if string(inBinary[0]) == "1" {
			out = append([]string{"wink"}, out...)
		}
	}
	if len(binaryIn) > 4 {
		for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
	}
	return out
}