package day7

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
	"sort"
	"strconv"
)

type hand struct {
	cards  map[string]int
	all    string
	bid    int
	rank   int64
	pairs  []string
	threes []string
	fours  []string
	fives  []string
}

func Solve() error {
	data, err := advent2023.Lines("day7")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	/*
		If two hands have the same type, a second ordering rule takes effect. Start by comparing the first
		card in each hand. If these cards are different, the hand with the stronger first card is considered
		stronger. If the first card in each hand have the same label, however, then move on to considering the
		second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue
		with the third card in each hand, then the fourth, then the fifth.
	*/

	result := int64(0)
	hands := make([]hand, 0)
	for _, line := range data {
		h := hand{
			//all:   advent2023.SortCards(line[:5]),
			all:    line[:5],
			cards:  make(map[string]int),
			bid:    0,
			pairs:  make([]string, 0),
			threes: make([]string, 0),
			fours:  make([]string, 0),
			fives:  make([]string, 0),
		}

		if b, err := strconv.Atoi(line[6:]); err != nil {
			return fmt.Errorf("reading score: %s", line)
		} else {
			h.bid = b
		}

		for _, r := range line[:5] {
			h.cards[string(r)]++
		}

		for card, cnt := range h.cards {
			switch cnt {
			case 2:
				h.pairs = append(h.pairs, card)
			case 3:
				h.threes = append(h.threes, card)
			case 4:
				h.fours = append(h.fours, card)
			case 5:
				h.fives = append(h.fives, card)
			}
		}

		h.rank = handRank(h)

		hands = append(hands, h)
		fmt.Printf("pairs: %#v\n", h)
	}

	sort.Slice(hands, func(i, j int) bool {
		// reverse order ranks go lower to bigger
		if hands[i].rank == hands[j].rank {
			/*
				If two hands have the same type, a second ordering rule takes effect. Start by comparing
				the first card in each hand. If these cards are different, the hand with the stronger first card
				is considered stronger.
				If the first card in each hand have the same label, however, then move on to considering the second card
				in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the
				third card in each hand, then the fourth, then the fifth.
			*/

			fmt.Printf("compare %q ?? %q \n", hands[i].all, hands[j].all)

			for ci := 0; ci < 5; ci++ {
				// compare
				a := advent2023.CardRank(rune(hands[i].all[ci]))
				b := advent2023.CardRank(rune(hands[j].all[ci]))
				fmt.Printf("compare %q (%d) =? %q (%d) \n", hands[i].all[ci], a, hands[j].all[ci], b)

				if a != b {
					return a < b
				}
			}
		}

		return hands[i].rank < hands[j].rank
	})

	for i, h := range hands {
		topCard, topCnt := top(h)
		fmt.Printf("sorted rank %d hand %v topCard %qx%d \n", i+1, h.all, topCard, topCnt)
		result += int64((i + 1) * h.bid)
	}

	slog.Info("day7 solution a", "hands", hands, "result", result)
	return nil
}

func SolveB() error {
	data, err := advent2023.Lines("day7-sample")
	if err != nil {
		return err
	}

	result := int64(1)
	slog.Info("day7 solution b", "data", data, "result", result)
	return nil
}

func load(name string) ([]hand, error) {
	hands := make([]hand, 0)
	data, err := advent2023.Lines(name)
	if err != nil {
		return hands, fmt.Errorf("error loading data: %w", err)
	}

	for _, line := range data {
		h := hand{
			//all:   advent2023.SortCards(line[:5]),
			all:    line[:5],
			cards:  make(map[string]int),
			bid:    0,
			pairs:  make([]string, 0),
			threes: make([]string, 0),
			fours:  make([]string, 0),
			fives:  make([]string, 0),
		}

		if b, err := strconv.Atoi(line[6:]); err != nil {
			return hands, fmt.Errorf("reading score: %s", line)
		} else {
			h.bid = b
		}

		for _, r := range line[:5] {
			h.cards[string(r)]++
		}
	}

	return hands, nil
}

func handRank(h hand) int64 {

	if len(h.fives) > 0 {
		// five of a kind lol wut
		return int64(7)
	}

	if len(h.fours) > 0 {
		// four of a kind
		return int64(6)
	}

	if len(h.threes) > 0 {
		if len(h.pairs) > 0 {
			// full house
			return int64(5)
		}

		// three of a kind
		return int64(4)
	}

	if len(h.pairs) > 0 {
		if len(h.pairs) > 1 {
			// two pairs
			return int64(3)
		}

		// one pair
		return int64(2)
	}

	// just a bunch of cards
	return 1
}

func top(h hand) (string, int) {
	res := 0
	c := ""

	for card, cnt := range h.cards {
		if res < cnt {
			res = cnt
			c = card
		}
	}

	return c, res
}
