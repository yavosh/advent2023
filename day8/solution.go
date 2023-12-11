package day8

import (
	"fmt"
	"log/slog"
	"math/big"
	"strings"

	"github.com/yavosh/advent2023"
)

type node struct {
	left, right string
}

func load(name string) (string, map[string]node, error) {
	lines, err := advent2023.Lines(name)
	if err != nil {
		return "", nil, fmt.Errorf("error loading data: %w", err)
	}

	instructions := lines[0]
	state := make(map[string]node)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}

		id := line[0:3]
		leafs := strings.Split(line[7:len(line)-1], ",")
		state[id] = node{
			left:  strings.TrimSpace(leafs[0]),
			right: strings.TrimSpace(leafs[1]),
		}
	}

	return instructions, state, nil
}

func Solve() error {
	instructions, state, err := load("day8")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	pos := 0
	nodeId := "AAA"
	for {
		inst := instructions[pos%len(instructions)]
		//for _, inst := range instructions {
		switch inst {
		case 'R':
			if n, ok := state[nodeId]; !ok {
				panic("unknown node id " + nodeId)
			} else {
				nodeId = n.right
			}
		case 'L':
			if n, ok := state[nodeId]; !ok {
				panic("unknown node id " + nodeId)
			} else {
				nodeId = n.left
			}
		default:
			panic("unknown instruction " + string(inst))

		}

		if nodeId == "ZZZ" {
			break
		}

		pos++
	}

	result := int64(pos + 1)
	slog.Info("day8 solution a", "result", result)
	return nil
}

func SolveB() error {
	instructions, state, err := load("day8")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	// since this is computationally infeasible to complete
	// we will establish when a range starts looping and then take the product of
	// the iterations to determine when all loops will converge.
	//
	// somewhat naively we assume that once an end condition is reached (suffix Z) the ranges will loop.
	// for the sample data this seems to hold verified manually
	//

	pos := 0
	heads := make([]string, 0)
	startNodes := make([]string, 0)
	activeNodes := make([]string, 0)
	nodePaths := make([][]string, 0)
	iterations := make([]int64, 0)

	for n := range state {
		if strings.HasSuffix(n, "A") {
			startNodes = append(startNodes, n)
			activeNodes = append(activeNodes, n)
			heads = append(heads, n)
			nodePaths = append(nodePaths, []string{})
			iterations = append(iterations, 0)
		}
	}

	for {
		inst := instructions[pos%len(instructions)]
		checks := 0

		for i, nodeID := range activeNodes {
			if iterations[i] != 0 {
				continue
			}

			checks++
			switch inst {
			case 'R':
				if n, ok := state[nodeID]; !ok {
					panic("unknown node id " + nodeID)
				} else {
					activeNodes[i] = n.right
				}
			case 'L':
				if n, ok := state[nodeID]; !ok {
					panic("unknown node id " + nodeID)
				} else {
					activeNodes[i] = n.left
				}
			default:
				panic("unknown instruction " + string(inst))
			}

			nodePaths[i] = append(nodePaths[i], activeNodes[i])

			if strings.HasSuffix(activeNodes[i], "Z") {
				//fmt.Printf("node path %d head:%s tail:%s %v\n", i, nodePaths[i][0], activeNodes[i], len(nodePaths[i]))
				//fmt.Println("##########################################################################")
				iterations[i] = int64(len(nodePaths[i]))
			}
		}

		if checks == 0 {
			// all paths have reached their loop
			break
		}

		pos++
	}

	result := int64(1)
	lcd := big.NewInt(0)
	for _, iteration := range iterations {
		// result is the product of all iterations
		result *= int64(iteration)
		lcd.GCD(nil, nil, lcd, big.NewInt(int64(iteration)))
	}

	// result is the least common multiplier of the iterations
	result = LcmInts(iterations...)

	slog.Info("day8 solution b", "result", result,
		"startNodes", startNodes,
		"activeNodes", activeNodes,
		"iterations", iterations,
		"lcd", lcd.String())

	return nil
}

func lcm(a, b int64) int64 {
	gcd := big.NewInt(0)
	ab := big.NewInt(a)
	bb := big.NewInt(b)
	gcd.GCD(nil, nil, ab, bb)

	// lcm of two numbers is their product divided by their gcd
	return ab.Mul(ab, bb).Div(ab, gcd).Int64()
}

// LcmInts will calculate the least common multiple of a bunch of integers
func LcmInts(in ...int64) int64 {
	if len(in) < 2 {
		panic("lcm of one number")
	}

	if len(in) == 2 {
		return lcm(in[0], in[1])
	}

	return lcm(in[0], LcmInts(in[1:]...))
}
