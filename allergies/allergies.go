package allergies

import (
	"sort"
)

const testVersion = 1

var allergens = map[string]int{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(in uint) []string {
	var out []string
	var nums []int
	for i := range allergens {
		nums = append(nums, allergens[i])
	}
	sort.Ints(nums)
	if in/uint(allergens["cats"]) > 1 {
		in = in - uint(allergens["cats"])*(in/uint(allergens["cats"]))
		out = append(out, "cats")
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if int64(nums[i]) <= int64(in) {
			in = in - uint(nums[i])
			for iR := range allergens {
				if allergens[iR] == nums[i] {
					out = append(out, iR)
				}
			}
		}
	}
	return out
}

func AllergicTo(i uint, allergen string) bool {
	var out bool = false
	allergenos := Allergies(i)
	for allerg := range allergenos {
		if allergenos[allerg] == allergen {
			return true
		}
	}
	return out
}