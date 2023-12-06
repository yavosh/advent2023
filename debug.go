package advent2023

import "encoding/json"

func Dump(in any) string {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return "ERROR:" + err.Error()
	}

	return string(b)
}
