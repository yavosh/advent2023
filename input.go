package advent2023

import (
	"embed"
	_ "embed"
	"strings"
)

//go:embed input/*
var input embed.FS

const (
	NewLine = "\n"
)

func Input(d string) (string, error) {
	file, err := input.ReadFile("input/" + d + ".txt")
	if err != nil {
		return "", err
	}

	return string(file), nil
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
			res[key] = append(res[key], trimLines(strings.Split(fields, fieldsSep)))
		}
	}

	return res, nil
}

func Lines(d string) ([]string, error) {
	if data, err := Input(d); err != nil {
		return nil, err
	} else {
		return trimLines(strings.Split(data, NewLine)), nil
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

// trimLines will return only string lines which are not empty
func trimLines(in []string) []string {
	out := make([]string, 0)
	for _, v := range in {
		if strings.TrimSpace(v) != "" {
			out = append(out, strings.TrimSpace(v))
		}
	}
	return out
}
