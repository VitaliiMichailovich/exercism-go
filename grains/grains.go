package grains

import (
	"errors"
)

const testVersion = 1

func Square(in int) (uint64, error) {
	var prev uint64 = 1
	var err error
	if in <= 0 || in > 64 {
		return 0, errors.New("Out of range!")
	}
	for i := 2; i <= in; i++ {
		prev = prev * 2
	}
	return prev, err
}

func Total() uint64 {
	var out uint64 = 1
	last, _ := Square(64)
	for i := last; i > 1; i = i / 2 {
		out += i
	}
	return out
}