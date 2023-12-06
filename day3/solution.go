package day3

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
)

func Solve() error {
	data, err := advent2023.Grid("day3")
	if err != nil {
		return err
	}

	result := 0

	part := ""
	adjacent := false

	var parts = make([]string, 0)
	var notParts = make([]string, 0)

	for y, line := range data {
		//slog.Info("data", "line", line)
		for x, r := range line {
			//slog.Info("rune", "y", y, "x", x, "r", string(r))

			if advent2023.IsDigit(r) {
				part += string(r)
				if adjacentSymbol(data, y, x) {
					adjacent = true
				}
			} else {
				if part != "" {
					//slog.Info("new part", "part", part, "adj", adjacent)
					if adjacent {
						parts = append(parts, part)
					} else {
						notParts = append(notParts, part)
					}
				}
				part = ""
				adjacent = false
			}

		}

	}

	result = advent2023.SumStrings(parts...)

	//slog.Info("data", "grid", data)
	//slog.Info("data", "parts", parts, "not parts", notParts)
	slog.Info("solution a", "result", result)
	return nil
}

func SolveB() error {
	data, err := advent2023.Grid("day3")
	if err != nil {
		return err
	}

	result := 0

	part := ""
	adjacent := false

	var parts = make([]string, 0)
	var notParts = make([]string, 0)
	var partGears = make(map[string]bool)
	var gears = make(map[string][]string)

	for y, line := range data {
		//slog.Info("data", "line", line)
		for x, r := range line {
			//slog.Info("rune", "y", y, "x", x, "r", string(r))

			//if r == '*' {
			//	id := fmt.Sprintf("%d_%d", y, x)
			//	cnt := adjacentCountFunc(data, y, x, advent2023.IsDigit)
			//	slog.Info("gear", "id", id, "cnt", cnt)
			//	gears[id]++
			//}

			if advent2023.IsDigit(r) {
				part += string(r)
				if adjacentCountFunc(data, y, x, advent2023.IsSymbol) > 0 {
					adjacent = true
				}

				if adjacent {
					g := adjacentTilesFunc(data, y, x, advent2023.IsGear)
					//slog.Info("part", "part", part, "gears", g)
					for _, gg := range g {
						partGears[gg] = true
					}
				}

			} else {
				if part != "" {
					//slog.Info("new part", "part", part, "adj", adjacent, "partGears", partGears)
					if adjacent {
						parts = append(parts, part)

						for g := range partGears {
							if _, ok := gears[g]; !ok {
								gears[g] = []string{part}
							} else {
								gears[g] = append(gears[g], part)
							}
						}

					} else {
						notParts = append(notParts, part)
					}
				}
				part = ""
				partGears = make(map[string]bool)
				adjacent = false
			}

		}

	}

	for _, parts := range gears {
		if len(parts) == 2 {
			result += advent2023.MulStrings(parts[0], parts[1])
		}
	}

	//slog.Info("data", "grid", data)
	//slog.Info("data", "parts", parts, "not parts", notParts)
	//slog.Info("gears", "gears", gears)
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
