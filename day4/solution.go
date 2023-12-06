package day4

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
	"strconv"
	"strings"
)

func Solve() error {
	lines, err := advent2023.Lines("day4")
	if err != nil {
		return err
	}

	result := 0

	for _, line := range lines {
		if !strings.HasPrefix(line, "Card") {
			// skip empty
			continue
		}

		val := line[strings.Index(line, ":")+1:]
		winners := strings.Fields(val[0:strings.Index(val, "|")])
		numbers := strings.Fields(val[strings.Index(val, "|")+1:])

		wins := advent2023.ToSet(winners)

		score := 0

		for _, n := range numbers {
			if wins.Contains(n) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}

		result += score
	}

	slog.Info("day4 solution a", "result", result)
	return nil
}

func SolveB() error {
	lines, err := advent2023.Lines("day4")
	if err != nil {
		return err
	}

	result := 0
	type scratchCard struct {
		ID      int
		Score   int
		Winners advent2023.Set
		Numbers []string
	}

	all := make([]int, 0)
	cards := make(map[int]scratchCard)

	for _, line := range lines {
		if !strings.HasPrefix(line, "Card") {
			// skip empty
			continue
		}

		card, err := strconv.Atoi(strings.TrimSpace(line[5:strings.Index(line, ":")]))
		if err != nil {
			return fmt.Errorf("invalid card number %q: %v", line, err)
		}

		val := line[strings.Index(line, ":")+1:]
		winners := strings.Fields(val[0:strings.Index(val, "|")])
		numbers := strings.Fields(val[strings.Index(val, "|")+1:])
		wins := advent2023.ToSet(winners)

		score := 0
		for _, n := range numbers {
			if wins.Contains(n) {
				score++
			}
		}

		all = append(all, card)
		cards[card] = scratchCard{
			ID:      card,
			Score:   score,
			Winners: wins,
			Numbers: numbers,
		}
	}

	for {
		if len(all) == 0 {
			// no more cards
			break
		}

		result++
		c := all[0]
		all = all[1:]

		card, ok := cards[c]
		if !ok {
			panic(fmt.Sprintf("missing card %d", c))
		}

		//fmt.Println("c", c, "all", all)
		for i := c + 1; i <= c+card.Score; i++ {
			all = append(all, i)
		}
		//fmt.Println("c", c, "all", all)
	}

	fmt.Printf("cards \n%s\n", advent2023.Dump(cards))
	slog.Info("day4 solution b", "result", result)
	return nil
}
