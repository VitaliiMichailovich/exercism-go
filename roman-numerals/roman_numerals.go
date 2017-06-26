package romannumerals

import "errors"

const testVersion = 3

var deci = []int{1000, 500, 100, 50, 10, 5, 1}
var roma = []string{"M", "D", "C", "L", "X", "V", "I"}

func ToRomanNumeral(in int) (string, error) {
	var out string
	if in <= 0 || in >= 4000 {
		return out, errors.New("Be positive =)")
	}
	for i := 0; i < len(deci); i++ {
		count := in / deci[i]
		if count >= 1 {
			if in >= 900 && in <= 999 {
				out += "CM"
				in = in - 900
			} else if in >= 400 && in <= 499 {
				out += "CD"
				in = in - 400
			} else if in >= 90 && in <= 99 {
				out += "XC"
				in = in - 90
			} else if in >= 40 && in <= 49 {
				out += "XL"
				in = in - 40
			} else if in == 9 {
				out += "IX"
				in = in - 9
			} else if in >= 4 && in <= 4 {
				out += "IV"
				in = in - 4
			} else {
				for ii := 1; ii <= count; ii++ {
					out += roma[i]
				}
				in = in - deci[i]*count
			}
		}
	}
	return out, nil
}