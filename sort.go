package advent2023

import (
	"sort"
)

var (
	//A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	cardRanks = map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		// lol not a real card
		'1': 1,
	}
)

func CardRank(c rune) int {
	if r, ok := cardRanks[c]; !ok {
		panic("not a card " + string(c))
	} else {
		return r
	}
}

func SortString(in string) string {
	s := []rune(in)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func SortCards(in string) string {
	s := []rune(in)
	sort.Slice(s, func(i int, j int) bool { return CardRank(s[i]) > CardRank(s[j]) })
	return string(s)
}
