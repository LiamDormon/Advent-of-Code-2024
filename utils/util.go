package utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type IntTuple struct {
	A, B int
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Zip(a, b []int) ([]IntTuple, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([]IntTuple, len(a), len(a))

	for i, e := range a {
		r[i] = IntTuple{e, b[i]}
	}

	return r, nil
}

func CountOccurrences[T comparable](slice []T, value T) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}
