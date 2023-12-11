package advent2023

import (
	"fmt"
	"strconv"
)

type Set map[string]struct{}

func (s Set) Contains(v string) bool {
	if _, ok := s[v]; ok {
		return true
	}

	return false
}

func ToSet(in ...string) Set {
	res := make(Set)
	for _, v := range in {
		res[v] = struct{}{}
	}

	return res
}

func IntKeys(in map[int]int) []int {
	res := make([]int, 0)
	for k := range in {
		res = append(res, k)
	}

	return res
}

func StrKeys(in map[string]int) []string {
	res := make([]string, 0)
	for k := range in {
		res = append(res, k)
	}

	return res
}

func Ints(in ...string) []int {
	res := make([]int, len(in))
	for i, v := range in {
		if a, err := strconv.ParseInt(v, 10, 32); err != nil {
			panic(fmt.Errorf("invalid value %q: %v", v, err))
		} else {
			res[i] = int(a)
		}
	}

	return res
}

func Int64s(in ...string) []int64 {
	res := make([]int64, len(in))
	for i, v := range in {
		if a, err := strconv.ParseInt(v, 10, 64); err != nil {
			panic(fmt.Errorf("invalid value %q: %v", v, err))
		} else {
			res[i] = a
		}
	}

	return res
}
