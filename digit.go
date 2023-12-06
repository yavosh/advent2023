package advent2023

import "fmt"

func IsDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}

func IsSpace(r rune) bool {
	return r == '.'
}

func IsGear(r rune) bool {
	return r == '*'
}

func IsSymbol(r rune) bool {
	if IsDigit(r) {
		return false
	}

	if r == '.' {
		return false
	}

	return true
}

func ToDigit(r rune) int {
	switch r {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		panic(fmt.Sprintf("not a digit %c", r))
	}
}
