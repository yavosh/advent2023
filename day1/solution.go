package day1

import (
	"log/slog"
	"strings"

	"github.com/yavosh/advent2023"
)

func Solve() error {
	data, err := advent2023.Lines("day1")
	if err != nil {
		return err
	}

	acc := 0
	for _, line := range data {
		var first, last rune
		for _, r := range line {
			if advent2023.IsDigit(r) {
				if first == 0 {
					first = r
				}
				last = r
			}

		}

		num := advent2023.ToDigit(first)*10 + advent2023.ToDigit(last)
		acc += num
	}

	slog.Info("day1 solution a", "result", acc)
	return err
}

func SolveB() error {
	data, err := advent2023.Lines("day1")
	if err != nil {
		return err
	}

	acc := 0
	for _, line := range data {
		line = clean(strings.ToLower(line))
		var first, last rune
		for _, r := range line {
			if advent2023.IsDigit(r) {
				if first == 0 {
					first = r
				}
				last = r
			}

		}

		num := advent2023.ToDigit(first)*10 + advent2023.ToDigit(last)
		acc += num
	}

	slog.Info("day1 solution b", "result", acc)
	return err
}

// clean will replace word values for digit names with actual digits
func clean(s string) string {
	subs := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for i := 0; i < len(s); i++ {
		for k, v := range subs {
			if strings.HasPrefix(s[i:], k) {
				s = s[:i] + v + s[i+len(k):]
			}
		}
	}

	return s
}
