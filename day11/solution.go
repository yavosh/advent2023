package day11

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/yavosh/advent2023"
)

type plan [][]rune
type pos [2]int

func (p plan) String() string {
	b := strings.Builder{}
	for y, row := range p {
		b.WriteString(fmt.Sprintf("%04d ", y))
		for _, r := range row {
			b.WriteRune(r)
		}

		b.WriteRune('\n')
	}

	return b.String()
}

func empty(space []rune) bool {
	res := true
	for _, r := range space {
		res = res && r == '.'
	}
	return res
}

func emptyCol(p plan, x int) bool {
	res := true
	for y := 0; y < len(p); y++ {
		res = res && p[y][x] == '.'
	}
	return res
}

func computeExpansions(p plan) ([]int, []int) {
	var emptyRows = make([]int, 0)
	var emptyCols = make([]int, 0)

	for i, row := range p {
		if empty(row) {
			emptyRows = append(emptyRows, i)
		}
	}

	for x := 0; x < len(p[0]); x++ {
		if emptyCol(p, x) {
			emptyCols = append(emptyCols, x)
		}
	}

	return emptyCols, emptyRows
}

func expand(p plan) plan {
	emptyCols, emptyRows := computeExpansions(p)

	offset := 0
	for _, e := range emptyRows {
		p = slices.Insert(p, e+offset, []rune(strings.Repeat(".", len(p[0]))))
		offset++
	}

	offset = 0
	for _, ex := range emptyCols {
		for y, line := range p {
			line = slices.Insert(line, ex+offset, '.')
			p[y] = line
		}
		offset++
	}

	return p
}

func galaxies(p plan) []pos {
	res := make([]pos, 0)
	id := 0
	for y := range p {
		for x := range p[y] {
			if p[y][x] == '#' {
				id++
				res = append(res, pos{y, x})
			}
		}
	}

	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func distance(p1, p2 pos) int {
	return abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
}

func distanceB(p1, p2 pos, cols, rows []int) int {
	//fmt.Printf("distanceB p1=%v p2=%v cols=%v rows=%v \n", p1, p2, cols, rows)

	xd := 0
	yd := 0

	factor := 1000000 - 1

	//  0 - 292
	//  1 - 374   -   82
	//?10 - 1112  -  820
	// 10 - 1030  -  738
	//100 - 8410  - 8118

	for _, r := range rows {
		if p1[0] < r && p2[0] > r {
			//fmt.Printf(" -- cross y %v -> %v %d \n", p1, p2, r)
			yd += factor
		}

		if p1[0] > r && p2[0] < r {
			//fmt.Printf(" -+ cross y %v -> %v %d \n", p1, p2, r)
			yd += factor
		}

	}

	for _, c := range cols {
		if p1[1] > c && p2[1] < c {
			//fmt.Printf(" ~+ cross x %v -> %v %d \n", p1, p2, c)
			//xd += 1 // factor 1
			xd += factor
		}

		if p1[1] < c && p2[1] > c {
			//fmt.Printf(" ~- cross x %v -> %v %d \n", p1, p2, c)
			//xd += 1 // factor 1
			xd += factor
		}
	}

	return (abs(p1[0]-p2[0]) + yd) + (abs(p1[1]-p2[1]) + xd)
}

func Solve() error {
	grid, err := advent2023.Grid("day11")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	p := expand(grid)

	result := int64(0)
	all := galaxies(p)
	g1 := all[0]
	rem := all[1:]
	pairs := 0
	for len(rem) > 0 {
		for _, g2 := range rem {
			result += int64(distance(g1, g2))
			pairs++
			//fmt.Println(" G ", g1, " => O ", g2, "distance=", distance(g1, g2))
		}
		g1 = rem[0]
		rem = rem[1:]
	}

	slog.Info("day11 solution a", "result", result, "pairs", pairs)
	return nil
}

func SolveB() error {
	grid, err := advent2023.Grid("day11")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	emptyCols, emptyRows := computeExpansions(grid)
	result := int64(0)
	all := galaxies(grid)
	g1 := all[0]
	rem := all[1:]
	pairs := 0
	for len(rem) > 0 {
		for _, g2 := range rem {
			result += int64(distanceB(g1, g2, emptyCols, emptyRows))
			pairs++
			//fmt.Println(" G ", g1, " => O ", g2, "distance=", distanceB(g1, g2, emptyCols, emptyRows))
		}
		g1 = rem[0]
		rem = rem[1:]
	}

	slog.Info("day11 solution b", "result", result, "pairs", pairs)
	return nil
}
