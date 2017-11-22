package lsproduct

import (
	"errors"
	"strconv"
)

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var max, curr, dig, productCurr int64
	var err error
	digitsArr := make([]int64, len(digits))
	if span > len(digits) || span < 0 {
		err = errors.New("1")
		return -1, err
	}
	if digits == "" || span == 0 {
		return 1, err
	}
	for i := 0; i < len(digits); i++ {
		dig, err = strconv.ParseInt(digits[i:i+1], 10, 64)
		if err == nil {
			digitsArr[i] = dig
		} else {
			return 0, err
		}
	}
	for i := 0; i <= len(digitsArr)-span; i++ {
		curr = 0
		productCurr = 0
		for ii := i; ii < i+span; ii++ {
			if ii == i {
				curr = digitsArr[ii]
				productCurr = digitsArr[ii]
			} else {
				curr *= digitsArr[ii]
				productCurr = productCurr*10 + digitsArr[ii]
			}
		}
		if max <= curr {
			max = curr
		}
	}
	return max, err
}