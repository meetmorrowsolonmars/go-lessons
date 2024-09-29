package domain

import "golang.org/x/exp/constraints"

func IsEvenNumber[T constraints.Integer](number T) bool {
	if number%2 == 0 {
		return true
	}

	return false
}
