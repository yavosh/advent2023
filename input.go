package advent2023

import (
	"embed"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input/*
var input embed.FS

const (
	NewLine = "\n"
)

type Set map[string]struct{}

func (s Set) Contains(v string) bool {
	if _, ok := s[v]; ok {
		return true
	}

	return false
}

func ToSet(in []string) Set {
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

func Input(d string) (string, error) {
	file, err := input.ReadFile("input/" + d + ".txt")
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func Lines(d string) ([]string, error) {
	if data, err := Input(d); err != nil {
		return nil, err
	} else {
		return strings.Split(data, NewLine), nil
	}
}

func Grid(d string) ([][]rune, error) {
	lines, err := Lines(d)
	if err != nil {
		return nil, err
	}

	res := make([][]rune, len(lines))
	for i, line := range lines {
		gridLine := make([]rune, len(line))
		for j, c := range line {
			gridLine[j] = c
		}

		res[i] = gridLine
	}

	return res, nil
}

func Sets(d string) (map[string][][]string, error) {
	keySep := ":"
	setsSep := ";"
	fieldsSep := ","

	res := map[string][][]string{}
	lines, err := Lines(d)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		key := line[:strings.Index(line, keySep)]
		rest := line[strings.Index(line, keySep)+1:]
		sets := strings.Split(rest, setsSep)
		res[key] = make([][]string, 0)

		for _, fields := range sets {
			res[key] = append(res[key], trim(strings.Split(fields, fieldsSep)))
		}
	}

	return res, nil
}

func trim(in []string) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = strings.TrimSpace(v)
	}
	return out
}
