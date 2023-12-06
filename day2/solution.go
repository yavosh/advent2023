package day2

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
	"strconv"
	"strings"
)

func Solve() error {
	data, err := advent2023.Sets("day2")
	if err != nil {
		return err
	}

	/*
		Determine which games would have been possible if the bag had been loaded with
		only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the
		IDs of those games?
	*/

	possibleGames := make([]int, 0)
	result := 0

	gameSet := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for game, sets := range data {
		check := true
		for _, set := range sets {
			check = check && possible(gameSet, toGame(set))
		}

		if check {
			possibleGames = append(possibleGames, mustParseGame(game))
		}
	}

	for _, g := range possibleGames {
		result += g
	}

	slog.Info("solution a", "result", result)
	return nil
}

func SolveB() error {
	data, err := advent2023.Sets("day2")
	if err != nil {
		return err
	}

	result := 0
	for _, sets := range data {
		gameSet := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range sets {
			minGame(gameSet, toGame(set))
		}

		result += gameSet["red"] * gameSet["green"] * gameSet["blue"]
	}

	slog.Info("solution b", "result", result)
	return nil
}

// toGame will convert a string array into a map of key/val where the val is converted into
// an int
func toGame(in []string) map[string]int {
	res := make(map[string]int)

	for _, v := range in {
		pair := strings.Split(v, " ")
		if len(pair) != 2 {
			panic(fmt.Sprintf("malformed pair %q", v))
		}

		if val, err := strconv.ParseInt(strings.TrimSpace(pair[0]), 10, 32); err != nil {
			panic(fmt.Sprintf("invalid pair value pair=%q val=%q", v, pair[0]))
		} else {
			res[strings.TrimSpace(pair[1])] = int(val)
		}
	}

	return res
}

// possible will test that all values in set can be satisfied by game
func possible(game map[string]int, set map[string]int) bool {
	res := true
	for k, v := range set {
		res = res && game[k] >= v
	}
	return res
}

// minGame establish what is the minimu required cubes for a game
func minGame(game map[string]int, set map[string]int) {
	for k, want := range set {
		if curr, ok := game[k]; ok {
			if curr < want {
				game[k] = want
			}
		} else {
			game[k] = want
		}
	}
}

func mustParseGame(g string) int {
	if id, err := strconv.ParseInt(strings.TrimSpace(g[4:]), 10, 32); err != nil {
		panic(fmt.Sprintf("invalid game %q", g))
	} else {
		return int(id)
	}
}
