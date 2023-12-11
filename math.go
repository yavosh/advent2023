package advent2023

import "strconv"

func Sum(in []int) int {
	acc := 0
	for _, a := range in {
		acc += a
	}

	return acc
}

// MulStrings will convert the provided strings to `int` and multiply them
func MulStrings(in ...string) int {
	acc := 1
	for _, v := range in {
		if a, err := strconv.ParseInt(v, 10, 64); err != nil {
			panic("invalid int value " + v)
		} else {
			acc *= int(a)
		}
	}

	return acc
}

// SumStrings will convert the provided strings to `int` and add them
func SumStrings(in ...string) int {
	acc := 0
	for _, v := range in {
		if a, err := strconv.ParseInt(v, 10, 64); err != nil {
			panic("invalid int value " + v)
		} else {
			acc += int(a)
		}
	}

	return acc
}
