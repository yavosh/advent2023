package day10

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/yavosh/advent2023"
)

type plan [][]rune

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

func (p plan) Y() int {
	return len(p)
}

func (p plan) X() int {
	return len(p[0])
}

func (p plan) At(y, x int) rune {
	return p[y][x]
}

func (p plan) AtPos(p1 pos) rune {
	return p[p1.Y()][p1.X()]
}

type pos [2]int

func (p pos) Y() int {
	return p[0]
}

func (p pos) X() int {
	return p[1]
}

func (p pos) move(p1 pos) pos {
	return pos{p[0] + p1[0], p[1] + p1[1]}
}

func (p pos) String() string {
	return fmt.Sprintf("(%d,%d)", p[0], p[1])
}

func start(grid plan) pos {
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'S' {
				return pos{y, x}
			}
		}
	}

	panic("invalid map")
}

// maps possible connections to a pipe

var (
	/*
		| is a vertical pipe connecting north and south.
		- is a horizontal pipe connecting east and west.
		L is a 90-degree bend connecting north and east.
		J is a 90-degree bend connecting north and west.
		7 is a 90-degree bend connecting south and west.
		F is a 90-degree bend connecting south and east.
		. is ground; there is no pipe in this tile.
		S is the starting position of the animal;
	*/

	N = pos{-1, 0} // N ⬆️
	S = pos{+1, 0} // S ⬇️
	W = pos{0, -1} // W ⬅️
	E = pos{0, +1} // E ➡️

	S2N = advent2023.ToSet("|", "7", "F") // check south to north N ⬆️ (aka up)
	N2S = advent2023.ToSet("|", "J", "L") // N to S ⬇️
	W2E = advent2023.ToSet("-", "7", "J") // W to E ➡️
	E2W = advent2023.ToSet("-", "F", "L") // E to W ⬅️

)

func directionName(d pos) string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	default:
		panic(fmt.Sprintf("invalid direction %v", d))
	}
}

func possible(plan plan, p pos) []pos {
	res := make([]pos, 0)

	ybounds := plan.Y() - 1
	xbounds := plan.X() - 1

	y := p[0]
	x := p[1]

	if y > 0 {
		// check south to north N ⬆️ (aka up)
		pipe := string(plan[y-1][x])
		if S2N.Contains(pipe) {
			res = append(res, pos{y - 1, x})
		}
	}

	if y < ybounds {
		// N to S ⬇️
		pipe := string(plan[y+1][x])
		if N2S.Contains(pipe) {
			res = append(res, pos{y + 1, x})
		}
	}

	if x < xbounds {
		// W to E ➡️
		pipe := string(plan[y][x+1])
		if W2E.Contains(pipe) {
			res = append(res, pos{y, x + 1})
		}
	}

	if x > 0 {
		// E to W ⬅️
		pipe := string(plan[y][x-1])
		if E2W.Contains(pipe) {
			res = append(res, pos{y, x - 1})
		}
	}

	return res
}

func follow(plan plan, from, to, end pos) (int, plan, error) {
	pplan := make([][]rune, len(plan))
	for x := 0; x < len(plan); x++ {
		pplan[x] = []rune(strings.Repeat(".", len(plan[0])))
	}

	pplan[from[0]][from[1]] = plan[from[0]][from[1]]
	pplan[to[0]][to[1]] = plan[to[0]][to[1]]

	moves := 0
	for {

		pipe := string(plan.AtPos(to))
		direction := pos{to[0] - from[0], to[1] - from[1]}
		//fmt.Println(" follow", pipe,
		//	"from", from, string(plan.AtPos(from)), " => ",
		//	to, string(plan.AtPos(to)),
		//	"direction", direction, directionName(direction),
		//	"distance", moves)

		from = to
		switch pipe {
		case "|":
			if direction == N {
				to = from.move(N)
			} else if direction == S {
				to = from.move(S)
			} else {
				panic(fmt.Sprintf("invalid connection %s => %s (direction %s)", from, to, directionName(direction)))
			}
		case "-":
			if direction == E {
				to = from.move(E)
			} else if direction == W {
				to = from.move(W)
			} else {
				panic(fmt.Sprintf("invalid connection %s => %s (direction %s)", from, to, directionName(direction)))
			}
		case "L":
			if direction == S {
				to = from.move(E)
			} else if direction == W {
				to = from.move(N)
			} else {
				panic(fmt.Sprintf("invalid connection %s => %s (direction %s)", from, to, directionName(direction)))
			}
		case "J":
			if direction == S {
				to = from.move(W)
			} else if direction == E {
				to = from.move(N)
			} else {
				panic(fmt.Sprintf("invalid connection %s => %s (direction %s)", from, to, directionName(direction)))
			}
		case "7":
			if direction == N {
				to = from.move(W)
			} else if direction == E {
				to = from.move(S)
			} else {
				panic(fmt.Sprintf("invalid connection %s => %s (direction %s)", from, to, directionName(direction)))
			}
		case "F":
			if direction == N {
				to = from.move(E)
			} else if direction == W {
				to = from.move(S)
			} else {
				panic(fmt.Sprintf("invalid connection %s => %s (direction %s)", from, to, directionName(direction)))
			}

		default:
			panic("can't follow " + pipe)
		}

		pipe = string(plan.AtPos(to))
		moves++

		pplan[to[0]][to[1]] = plan[to[0]][to[1]]

		// end condition
		if to == end {
			break
		}
	}

	return moves, pplan, nil
}

func Solve() error {
	grid, err := advent2023.Grid("day10")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	result := int64(0)

	s := start(grid)
	for _, c := range possible(grid, s) {
		moves, _, err := follow(grid, s, c, s)
		if err != nil {
			fmt.Printf("not valid loop  %s -> %s (%v) \n", s, c, err)
		} else {
			//fmt.Printf("solution with moves %d\n", moves)
			result = int64(moves) / 2
			if result%2 != 0 {
				// add one more if odd number of moves
				result++
			}
		}
	}

	slog.Info("day10 solution a", "result", result)
	return nil
}

func SolveB() error {
	grid, err := advent2023.Grid("day10-sample-a")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	//fmt.Printf("%#v", plan)

	result := int64(1)

	fmt.Println(plan(grid))

	s := start(grid)
	startPos := possible(grid, s)
	if len(startPos) != 2 {
		return fmt.Errorf("pipe loop needs two possible starting points only")
	}

	curr := startPos[0]

	moves, pplan, err := follow(grid, s, curr, s)
	if err != nil {
		fmt.Printf("not valid loop  %s -> %s (%v) \n", s, curr, err)
	} else {
		//fmt.Printf("solution with moves %d\n", moves)
		result = int64(moves) / 2
		if result%2 != 0 {
			// add one more if odd number of moves
			result++
		}
	}

	fmt.Println(pplan)

	iplan := make([][]rune, len(grid))
	for x := 0; x < len(grid); x++ {
		iplan[x] = []rune(strings.Repeat(" ", len(grid[0])))
	}

	tiles := 0
	inside := 0
	for y, line := range pplan {
		for x := range line {

			if x == len(line)-1 {
				continue
			}

			fmt.Printf("%c", pplan[y][x])

			if inside%2 == 1 {
				if pplan[y][x] == '.' {
					// count tiles
					tiles++
					iplan[y][x] = 'I'
				} else {
					inside++
				}
				continue
			} else {
				if pplan[y][x] != '.' {
					inside++
				}
			}
		}

		inside = 0
		fmt.Println()
	}

	fmt.Println(plan(iplan))

	result = int64(tiles)
	slog.Info("day10 solution b", "result", result)

	return nil
}
