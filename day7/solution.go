package day7

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards  map[string]int
	all    string
	bid    int
	rank   int64
	jokers int
	pairs  []string
	threes []string
	fours  []string
	fives  []string
}

func (h hand) String() string {
	return fmt.Sprintf("%q {b:%d j:%d r:%d pairs:%v threes:%v fours:%v fives:%v}",
		h.all, h.bid, h.jokers, h.rank, h.pairs, h.threes, h.fours, h.fives)
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
		if line == "" {
			continue
		}

		h := hand{
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

			for ci := 0; ci < 5; ci++ {
				// compare
				a := advent2023.CardRank(rune(hands[i].all[ci]))
				b := advent2023.CardRank(rune(hands[j].all[ci]))
				//fmt.Printf("compare %q (%d) =? %q (%d) \n", hands[i].all[ci], a, hands[j].all[ci], b)

				if a != b {
					return a < b
				}
			}
		}

		return hands[i].rank < hands[j].rank
	})

	for i, h := range hands {
		result += int64((i + 1) * h.bid)
	}

	slog.Info("day7 solution a", "result", result)
	return nil
}

func SolveB() error {
	data, err := advent2023.Lines("day7")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	result := int64(0)
	hands := make([]hand, 0)
	for _, line := range data {
		if line == "" {
			continue
		}

		h := hand{
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
			if r == 'J' {
				// don't count jokers as cards
				h.jokers++
				continue
			}

			h.cards[string(r)]++
		}

		for card, cnt := range h.cards {
			if card == "J" {
				// don't count jokers as pairs
				continue
			}

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

		// apply jokers
		for h.jokers > 0 {
			// fmt.Printf("joker hand: %s\n", h)

			if h.all == "JJJJJ" {
				// special case, full jokers = JJJJJ -> AAAAA
				h.fives = []string{"A"}
				h.jokers = 0
				break
			}

			h.jokers--
			if len(h.fours) > 0 {
				// fours become fives, there can be only one, so we do not care to test
				// change it to a five
				h.fives = append(h.fives, h.fours[0])
				h.fours = h.fours[1:]
				continue
			}

			if len(h.threes) > 0 {
				// threes become fours, there can be only one, so we do not care to test which is better
				// change it to a four
				h.fours = append(h.fours, h.threes[0])
				h.threes = h.threes[1:]
				continue
			}

			if len(h.pairs) > 0 {
				// twos become threes, there can be many but the power is coming from card in first position so we also
				// do not care to test which is better, just grab the first pair
				// change it to a four
				h.threes = append(h.threes, h.pairs[0])
				h.pairs = h.pairs[1:]
				continue
			}

			// pick any card to make it a pair
			// except jokers
			cards := strings.ReplaceAll(h.all, "J", "")
			if cards == "" {
				panic("full joker hand")
			} else {
				h.pairs = append(h.pairs, string(cards[0]))
			}

		}

		h.rank = handRank(h)
		hands = append(hands, h)
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

			// fmt.Printf("compare %q ?? %q \n", hands[i].all, hands[j].all)

			for ci := 0; ci < 5; ci++ {
				// compare
				a := advent2023.CardRankB(rune(hands[i].all[ci]))
				b := advent2023.CardRankB(rune(hands[j].all[ci]))
				//fmt.Printf("compare %q (%d) =? %q (%d) \n", hands[i].all[ci], a, hands[j].all[ci], b)

				if a != b {
					return a < b
				}
			}
		}

		return hands[i].rank < hands[j].rank
	})

	for i, h := range hands {
		//fmt.Printf("sorted rank %d hand %s\n", i+1, h)
		result += int64((i + 1) * h.bid)
	}

	slog.Info("day7 solution b", "result", result)
	return nil
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
