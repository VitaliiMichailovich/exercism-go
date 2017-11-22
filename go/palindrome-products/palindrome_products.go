package palindrome

import (
	"errors"
	"math"
	"sort"
	"strconv"
)

const testVersion = 1

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	var pmin, pmax Product
	var fArr [][3]int
	var sortFArr []int
	var err error
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return pmin, pmax, err
	}
	for min1 := fmin; min1 <= fmax; min1++ {
		for min2 := min1; min2 <= fmax; min2++ {
			mult := int(math.Abs(float64(min1 * min2)))
			if PalindromeChecker(mult) == true {
				fArr = append(fArr, [3]int{min1 * min2, min1, min2})
				sortFArr = append(sortFArr, min1*min2)
			}
		}
	}
	if len(sortFArr) == 0 {
		err = errors.New("no palindromes")
		return pmin, pmax, err
	}
	sort.Ints(sortFArr)
	pmin.Product = sortFArr[0]
	pmax.Product = sortFArr[len(sortFArr)-1]
	for i := 0; i < len(fArr); i++ {
		if pmin.Product == fArr[i][0] {
			pmin.Factorizations = append(pmin.Factorizations, [2]int{fArr[i][1], fArr[i][2]})
		}
		if pmax.Product == fArr[i][0] {
			pmax.Factorizations = append(pmax.Factorizations, [2]int{fArr[i][1], fArr[i][2]})
		}
	}
	return pmin, pmax, err
}

func PalindromeChecker(in int) bool {
	inS := strconv.Itoa(in)
	halfLenInS := int(len(inS) / 2)
	count := 0
	for i := 0; i < halfLenInS; i++ {
		if inS[i:i+1] == inS[len(inS)-1-i:len(inS)-i] {
			count++
		}
	}
	if count == halfLenInS {
		return true
	}
	return false
}