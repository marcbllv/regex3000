package state

import "unicode/utf8"

type OppositeSetInspector struct {
	oppositeCharSet map[rune]bool
}

func NewOppositeSetInspector(chars []rune) OppositeSetInspector {
	oppositeCharSet := make(map[rune]bool)
	for _, char := range chars {
		oppositeCharSet[char] = true
	}
	return OppositeSetInspector{oppositeCharSet}
}

func (inspector OppositeSetInspector) Match(str string) (bool, string) {
	if len(str) == 0 {
		return false, ""
	}
	firstRune, runeSize := utf8.DecodeRuneInString(str)
	if inspector.oppositeCharSet[firstRune] {
		return false, ""
	}
	return true, str[runeSize:]
}

func (inspector OppositeSetInspector) Copy() Inspector {
	newOppositeCharSet := make(map[rune]bool)
	for k := range inspector.oppositeCharSet {
		newOppositeCharSet[k] = true
	}
	return SetInspector{charSet: newOppositeCharSet}
}
