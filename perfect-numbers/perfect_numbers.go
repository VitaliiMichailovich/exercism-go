package perfect

import (
	"errors"
)

const testVersion = 1

type Classification string

const (
	ClassificationDeficient Classification = "Deficient"
	ClassificationAbundant  Classification = "Abundant"
	ClassificationPerfect   Classification = "Perfect"
)

var (
	ErrOnlyPositive error = errors.New("ErrOnlyPositive")
)

func Classify(in uint64) (Classification, error) {
	var summ, i uint64
	if in <= 0 {
		return "", ErrOnlyPositive
	}
	for i = 1; i < in; i++ {
		if in%i == 0 {
			summ += i
		}
	}
	if in == summ {
		return ClassificationPerfect, nil
	}
	if in > summ {
		return ClassificationDeficient, nil
	}
	return Classification(ClassificationAbundant), nil
}