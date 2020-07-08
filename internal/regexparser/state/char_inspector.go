package state

import "unicode/utf8"

type CharInspector struct {
	char rune
}

func (inspector CharInspector) Match(str string) (bool, string) {
	firstRune, runeSize := utf8.DecodeRuneInString(str)
	if inspector.char == firstRune {
		return true, str[runeSize:]
	}
	return false, ""
}

func (inspector CharInspector) Copy() Inspector {
	return CharInspector{char: inspector.char}
}
