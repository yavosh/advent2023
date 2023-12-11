package day9

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/yavosh/advent2023"
)

// zeros  returns true if all values in `in` are equal to 0
func zeros(in []int) bool {
	res := true
	for _, v := range in {
		res = res && v == 0
	}

	return res
}

func Solve() error {
	lines, err := advent2023.Lines("day9-sample")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		measures := advent2023.Ints(strings.Fields(line)...)
		deltas := make([]int, len(measures)-1)
		for i := 0; i < len(measures)-1; i++ {
			deltas[i] = measures[i+1] - measures[i]
		}

		//slog.Info("iteration", "measures", measures, "deltas", deltas)
	}

	result := int64(9)
	slog.Info("day9 solution a", "result", result)
	return nil
}

func SolveB() error {
	_, err := advent2023.Lines("day9-sample")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	result := int64(1)

	slog.Info("day9 solution b", "result", result)

	return nil
}
