package day6

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
	"strconv"
	"strings"
)

type race struct {
	id       string
	dur      int64
	distance int64
}

func Solve() error {
	races, err := load("day6")
	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	result := int64(1)
	charge := int64(1) // millimeter per millisecond

	/*
		    To guarantee you win the grand prize, you need to make sure you go farther in each race than the
		    current record holder.

			Your toy boat has a starting speed of zero millimeters per millisecond.
			For each whole millisecond you spend at the beginning of the race holding down the button, the boat's
			speed increases by one millimeter per millisecond.
	*/

	for _, r := range races {
		wins := int64(0)
		for i := int64(0); i <= r.dur; i++ {
			// travel is charge times remaining time
			travel := (i * charge) * (r.dur - i)
			//slog.Info(fmt.Sprintf("for %d ms we would travel %d", i, travel), "race", r)

			if travel > r.distance {
				//slog.Info(fmt.Sprintf("WIN for %d ms we would travel %d", i, travel), "race", r)
				wins++
			}
		}

		result *= wins
	}

	//slog.Info("data", "races", races)
	slog.Info("solution a", "result", result)
	return nil
}

func SolveB() error {
	theRace, err := load2("day6")
	if err != nil {
		return err
	}

	result := int64(1)
	charge := int64(1) // millimeter per millisecond

	wins := int64(0)

	// time complexity could be simplified
	for i := int64(0); i <= theRace.dur; i++ {
		// travel is charge times remaining time
		travel := (i * charge) * (theRace.dur - i)
		//slog.Info(fmt.Sprintf("for %d ms we would travel %d", i, travel), "race", r)
		if travel > theRace.distance {
			//slog.Info(fmt.Sprintf("WIN for %d ms we would travel %d", i, travel), "race", r)
			wins++
		}
	}

	result = wins

	//slog.Info("data", "theRace", theRace)
	slog.Info("solution b", "result", result)
	return nil
}

func load(name string) ([]race, error) {
	data, err := advent2023.Lines(name)
	if err != nil {
		return nil, err
	}

	var races []race

	for _, line := range data {
		if strings.HasPrefix(line, "Time:    ") {
			values := strings.Fields(line[9:])
			races = make([]race, len(values))
			for i, v := range values {
				if a, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64); err != nil {
					return nil, fmt.Errorf("invalid time value %q: %v", v, err)
				} else {
					races[i].id = fmt.Sprintf("race-%d", i+1)
					races[i].dur = a
				}
			}
		}

		if strings.HasPrefix(line, "Distance:") {
			values := strings.Fields(line[9:])
			for i, v := range values {
				if a, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64); err != nil {
					return nil, fmt.Errorf("invalid distance value %q: %v", v, err)
				} else {
					races[i].distance = a
				}
			}
		}

	}

	return races, nil
}

func load2(name string) (race, error) {
	data, err := advent2023.Lines(name)
	if err != nil {
		return race{}, err
	}

	r := race{
		id:       "the big one",
		dur:      0,
		distance: 0,
	}

	for _, line := range data {
		if strings.HasPrefix(line, "Time:    ") {
			val := strings.ReplaceAll(line[9:], " ", "")
			if a, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err != nil {
				return race{}, fmt.Errorf("error parsing line %q: %v", line, err)
			} else {
				r.dur = a
			}
		}
		if strings.HasPrefix(line, "Distance:") {
			val := strings.ReplaceAll(line[9:], " ", "")
			if a, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err != nil {
				return race{}, fmt.Errorf("error parsing line %q: %v", line, err)
			} else {
				r.distance = a
			}
		}
	}

	return r, nil
}
