package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

func Gen(in byte) (string, error) {
	var outUp, outDown string
	if in < 'A' || in > 'Z' {
		return "", errors.New("Enter correct byte.")
	}
	for i := 1; i <= int(in)-0x41+1; i++ {
		line := strings.Repeat(" ", int(in)-0x41+1-i)
		if i == 1 {
			line = line + string(0x41-1+i) + line
		} else {
			line = line + string(0x41-1+i) + strings.Repeat(" ", i*2-3) + string(0x41-1+i) + line
		}
		outUp += line + "\n"
		if i == 1 && in != 0x41 {
			outDown = line + outDown + "\n"
		} else if i != int(in)-0x41+1 {
			outDown = line + "\n" + outDown
		}
	}
	return outUp + outDown, nil
}
