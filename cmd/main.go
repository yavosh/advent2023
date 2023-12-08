package main

import (
	"github.com/yavosh/advent2023/day1"
	"github.com/yavosh/advent2023/day2"
	"github.com/yavosh/advent2023/day3"
	"github.com/yavosh/advent2023/day4"
	"github.com/yavosh/advent2023/day5"
	"github.com/yavosh/advent2023/day6"
	"github.com/yavosh/advent2023/day7"
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

	if err := day4.Solve(); err != nil {
		slog.Error("error solving %s %v", "day2", err)
	}

	if err := day4.SolveB(); err != nil {
		slog.Error("error solving %s %v", "day2b", err)
	}

	if err := day5.Solve(); err != nil {
		slog.Error("error solving %s %v", "day2", err)
	}

	slog.Info("day5 solution b ***SKIP***")
	//if err := day5.SolveB(); err != nil {
	//	slog.Error("error solving %s %v", "day2b", err)
	//}

	if err := day6.Solve(); err != nil {
		slog.Error("error solving %s %v", "day2", err)
	}

	if err := day6.SolveB(); err != nil {
		slog.Error("error solving %s %v", "day2b", err)
	}

	if err := day7.Solve(); err != nil {
		slog.Error("error solving %s %v", "day2", err)
	}

	if err := day7.SolveB(); err != nil {
		slog.Error("error solving %s %v", "day2b", err)
	}

}
