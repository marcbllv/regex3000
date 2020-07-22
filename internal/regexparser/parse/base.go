package parse

import "fmt"

type T interface{}

type Parser interface {
	Parse(str []rune, pos int) (T, int)
}

type IntParser struct {
	optional bool
}

func (intParser IntParser) Parse(str []rune, pos int) (int, int) {
	currentPosition := pos
	intValue := 0
	for {
		if currentPosition >= len(str) {
			break
		}
		char := str[currentPosition]
		if char < 48 || char > 57 {
			break
		}
		intValue = intValue*10 + (int(char) - 48)
		currentPosition++
	}

	if currentPosition == pos && !intParser.optional {
		msg := fmt.Sprintf("Expected integer at position %d", pos)
		panic(msg)
	} else if currentPosition == pos && intParser.optional {
		return -1, currentPosition
	}
	return intValue, currentPosition
}

type RuneParser struct {
	optional bool
	r        rune
}

func (runeParser RuneParser) Parse(str []rune, pos int) (rune, int) {
	if str[pos] == runeParser.r {
		return runeParser.r, pos + 1
	}

	if !runeParser.optional {
		msg := fmt.Sprintf("Expected char '%c' at position %d", runeParser.r, pos)
		panic(msg)
	} else {
		return -1, pos
	}
}
