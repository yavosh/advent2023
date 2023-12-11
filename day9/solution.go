package day9

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/yavosh/advent2023"
)

func reverse(in []int) []int {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	return in
}

func prior(in []int) int {
	deltas := make([]int, len(in)-1)
	for i := len(in) - 1; i > 0; i-- {
		deltas[i-1] = in[i] - in[i-1]
	}

	allZeros := true
	for _, v := range deltas {
		allZeros = allZeros && v == 0
	}

	if allZeros {
		//slog.Info("zeros", "p", in[0])
		return in[0]
	}

	p := prior(deltas)
	//slog.Info("prior", "p", p, "deltas", deltas, "measures", in)
	return in[0] - p
}

func next(in []int) int {
	deltas := make([]int, 0)
	for i := 0; i < len(in)-1; i++ {
		deltas = append(deltas, in[i+1]-in[i])
	}

	allZeros := true
	for _, v := range deltas {
		allZeros = allZeros && v == 0
	}

	if allZeros {
		return in[len(in)-1]
	}

	return in[len(in)-1] + next(deltas)
}

func Solve() error {
	lines, err := advent2023.Lines("day9")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	result := int64(0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		measures := advent2023.Ints(strings.Fields(line)...)
		result += int64(next(measures))
	}

	slog.Info("day9 solution a", "result", result)
	return nil
}

func SolveB() error {
	lines, err := advent2023.Lines("day9")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	result := int64(0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		measures := advent2023.Ints(strings.Fields(line)...)
		//slog.Info("solve ", "p", prior(measures), "measures", measures)
		result += int64(prior(measures))
	}

	slog.Info("day9 solution b", "result", result)

	return nil
}
