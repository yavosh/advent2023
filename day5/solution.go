package day5

import (
	"fmt"
	"github.com/yavosh/advent2023"
	"log/slog"
	"math"
	"strings"
)

type mapper struct {
	src, dst int64
	c        int64
}

func Solve() error {
	seeds, maps, err := parse("day5")
	if err != nil {
		return err
	}

	result := int64(math.MaxInt64)
	path := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	for _, seed := range seeds {
		res := walk(seed, maps, path)
		//slog.Info("res ", "seed", seed, "res", res)

		if res < result {
			result = res
		}

	}

	slog.Info("day5 solution a", "result", result)
	return nil
}

func SolveB() error {
	seeds, maps, err := parse("day5")
	if err != nil {
		return err
	}

	result := int64(math.MaxInt64)
	path := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
		"water-to-light",
		"light-to-temperature",
		"temperature-to-humidity",
		"humidity-to-location",
	}

	seedRange := make([][]int64, 0)
	for i := 0; i < len(seeds)/2; i++ {
		seedRange = append(seedRange, []int64{seeds[(i * 2)], seeds[(i*2)+1]})
	}

	slog.Info("seeds", "seeds", seeds, "seedRange", seedRange)

	//res := walk(14, maps, path)
	//slog.Info("res 14", "res", res)

	cnt := 0

	visited := make(map[int64]int64)

	for _, rng := range seedRange {
		slog.Info("range", "range", rng)

		for i := rng[0]; i < rng[0]+rng[1]; i++ {

			seed := i

			//slog.Info("seed", "range", rng, "seed", seed)
			cnt++

			if _, ok := visited[seed]; !ok {
				res := walk(seed, maps, path)
				if res < result {
					result = res
				}
			}
		}

	}

	slog.Info("day5 solution b", "seeds", seeds, "cnt", cnt, "result", result)
	return nil
}

func walk(id int64, data map[string][]mapper, maps []string) int64 {
	res := id

	for _, name := range maps {
		mappers, ok := data[name]
		if !ok {
			panic("unknown map " + name)
		}

		for _, mapper := range mappers {
			if res >= mapper.src && res < mapper.src+mapper.c {
				delta := res - mapper.src
				res = mapper.dst + delta
				break // done
			} else {
				// 1-1 mapping no change, equivalent to res=res
			}
		}
	}

	return res
}

func parse(in string) ([]int64, map[string][]mapper, error) {
	data, err := advent2023.Lines(in)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading data: %w", err)
	}

	var seeds []string
	var name string
	values := make([]mapper, 0)
	maps := make(map[string][]mapper)

	for _, line := range data {
		if strings.HasPrefix(line, "seeds:") {
			seeds = strings.Fields(line[6:])
			continue
		}

		if strings.HasSuffix(line, "map:") {
			name = line[0:strings.LastIndex(line, " map:")]
			continue
		}

		if line == "" {
			// close map
			if name != "" {
				//slog.Info("map", "map", name, "values", values)
				maps[name] = values
				name = ""
				values = make([]mapper, 0)
			}

			continue
		}

		fields := advent2023.Int64s(strings.Fields(line)...)
		if len(fields) != 3 {
			return nil, nil, fmt.Errorf("could not parse line %q for map %q", line, name)
		}

		values = append(values, mapper{dst: fields[0], src: fields[1], c: fields[2]})
		//slog.Info("map", "map", name, "values", values)
	}

	return advent2023.Int64s(seeds...), maps, nil
}
