package day6

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
)

func Solve() error {
	data, err := advent2023.Grid("day6-sample")
	if err != nil {
		return err
	}

	result := 0
	speed := 0.0
	charge := 1.0 // millimeter per millisecond

	/*
		    To guarantee you win the grand prize, you need to make sure you go farther in each race than the
		    current record holder.

			Your toy boat has a starting speed of zero millimeters per millisecond.
			For each whole millisecond you spend at the beginning of the race holding down the button, the boat's
			speed increases by one millimeter per millisecond.
	*/

	slog.Info("data", "grid", data)
	slog.Info("x", "speed", speed, "charge", charge)
	slog.Info("solution a", "result", result)
	return nil
}

func SolveB() error {
	data, err := advent2023.Grid("day6-sample")
	if err != nil {
		return err
	}

	result := 0

	slog.Info("data", "grid", data)
	slog.Info("solution b", "result", result)
	return nil
}

func adjacentSymbol(plan [][]rune, y, x int) bool {
	if len(plan) == 0 || len(plan[0]) == 0 {
		panic("empty plan")
	}

	ybounds := len(plan) - 1
	xbounds := len(plan[0]) - 1

	//slog.Info("plan", "ybounds", ybounds, "xbounds", xbounds)

	if y > 0 && x > 0 {
		// up left
		if advent2023.IsSymbol(plan[y-1][x-1]) {
			return true
		}
	}

	if y > 0 {
		// up
		if advent2023.IsSymbol(plan[y-1][x]) {
			return true
		}
	}

	if y > 0 && x < xbounds {
		// up right
		if advent2023.IsSymbol(plan[y-1][x+1]) {
			return true
		}
	}

	if x < xbounds {
		// right
		if advent2023.IsSymbol(plan[y][x+1]) {
			return true
		}
	}

	if y < ybounds && x < xbounds {
		// right down
		if advent2023.IsSymbol(plan[y+1][x+1]) {
			return true
		}
	}

	if y < ybounds {
		// down
		if advent2023.IsSymbol(plan[y+1][x]) {
			return true
		}
	}

	if y < ybounds && x > 0 {
		// down left
		if advent2023.IsSymbol(plan[y+1][x-1]) {
			return true
		}
	}

	if x > 0 {
		// left
		if advent2023.IsSymbol(plan[y][x-1]) {
			return true
		}
	}

	return false
}

func adjacentCountFunc(plan [][]rune, y, x int, check func(rune) bool) int {
	if len(plan) == 0 || len(plan[0]) == 0 {
		panic("empty plan")
	}

	ybounds := len(plan) - 1
	xbounds := len(plan[0]) - 1

	//slog.Info("plan", "ybounds", ybounds, "xbounds", xbounds)

	acc := 0

	if y > 0 && x > 0 {
		// up left
		if check(plan[y-1][x-1]) {
			acc++
		}
	}

	if y > 0 {
		// up
		if check(plan[y-1][x]) {
			acc++
		}
	}

	if y > 0 && x < xbounds {
		// up right
		if check(plan[y-1][x+1]) {
			acc++
		}
	}

	if x < xbounds {
		// right
		if check(plan[y][x+1]) {
			acc++
		}
	}

	if y < ybounds && x < xbounds {
		// right down
		if check(plan[y+1][x+1]) {
			acc++
		}
	}

	if y < ybounds {
		// down
		if check(plan[y+1][x]) {
			acc++
		}
	}

	if y < ybounds && x > 0 {
		// down left
		if check(plan[y+1][x-1]) {
			acc++
		}
	}

	if x > 0 {
		// left
		if check(plan[y][x-1]) {
			acc++
		}
	}

	return acc
}

func adjacentTilesFunc(plan [][]rune, y, x int, check func(rune) bool) []string {
	if len(plan) == 0 || len(plan[0]) == 0 {
		panic("empty plan")
	}

	ybounds := len(plan) - 1
	xbounds := len(plan[0]) - 1

	//slog.Info("plan", "ybounds", ybounds, "xbounds", xbounds)

	acc := 0
	tiles := make([]string, 0)

	if y > 0 && x > 0 {
		// up left
		if check(plan[y-1][x-1]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y-1, x-1))
		}
	}

	if y > 0 {
		// up
		if check(plan[y-1][x]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y-1, x))
		}
	}

	if y > 0 && x < xbounds {
		// up right
		if check(plan[y-1][x+1]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y-1, x+1))
		}
	}

	if x < xbounds {
		// right
		if check(plan[y][x+1]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y, x+1))
		}
	}

	if y < ybounds && x < xbounds {
		// right down
		if check(plan[y+1][x+1]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y+1, x+1))
		}
	}

	if y < ybounds {
		// down
		if check(plan[y+1][x]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y+1, x))
		}
	}

	if y < ybounds && x > 0 {
		// down left
		if check(plan[y+1][x-1]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y+1, x-1))
		}
	}

	if x > 0 {
		// left
		if check(plan[y][x-1]) {
			acc++
			tiles = append(tiles, fmt.Sprintf("%d_%d", y, x-1))
		}
	}

	return tiles
}
