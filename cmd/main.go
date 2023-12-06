package main

import (
	"github.com/yavosh/advent2023/day1"
	"github.com/yavosh/advent2023/day2"
	"github.com/yavosh/advent2023/day3"
	"log/slog"
)

func main() {

	if err := day1.Solve(); err != nil {
		slog.Error("error solving %s %v", "day1", err)
	}

	if err := day1.SolveB(); err != nil {
		slog.Error("error solving %s %v", "day1b", err)
	}

	if err := day2.Solve(); err != nil {
		slog.Error("error solving %s %v", "day2", err)
	}

	if err := day2.SolveB(); err != nil {
		slog.Error("error solving %s %v", "day2b", err)
	}

	if err := day3.Solve(); err != nil {
		slog.Error("error solving %s %v", "day2", err)
	}

	if err := day3.SolveB(); err != nil {
		slog.Error("error solving %s %v", "day2b", err)
	}

}
