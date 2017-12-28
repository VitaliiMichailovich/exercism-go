package sublist

type Relation string

func Sublist(listOne, listTwo []int) Relation {
	if len(listOne) == len(listTwo) {
		if len(listOne) == 0 {
			return "equal"
		}
		for k, v := range listOne {
			if listTwo[k] != v {
				return "unequal"
			}
		}
		return "equal"
	} else if len(listOne) > len(listTwo) {
		if len(listTwo) == 0 {
			return "superlist"
		}
		var firsts []int
		for k, v := range listOne {
			if v == listTwo[0] {
				firsts = append(firsts, k)
			}
		}
		if len(firsts) == 0 {
			return "unequal"
		} else {
			cont := 0
			for _, valFirsts := range firsts {
				for k, v := range listTwo {
					if valFirsts+k > len(listTwo) {
						break
					}
					if listOne[valFirsts+k] != v {
						cont++
						break
					}
				}
			}
			if cont != len(firsts) {
				return "superlist"
			}
			return "unequal"
		}
	} else if len(listOne) < len(listTwo) {
		if len(listOne) == 0 {
			return "sublist"
		}
		var firsts []int
		for k, v := range listTwo {
			if v == listOne[0] {
				firsts = append(firsts, k)
			}
		}
		if len(firsts) == 0 {
			return "unequal"
		} else {
			cont := 0
			for _, valFirsts := range firsts {
				for k, v := range listOne {
					if valFirsts+k > len(listOne) {
						break
					}
					if listTwo[valFirsts+k] != v {
						cont++
						break
					}
				}

			}
			if cont != len(firsts) {
				return "sublist"
			}
			return "unequal"
		}
	}
	return "---"
}
