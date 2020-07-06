package state

import "unicode/utf8"

type SetInspector struct {
	charSet map[rune]bool
}

func (inspector SetInspector) Match(str string) (bool, string) {
	firstRune, runeSize := utf8.DecodeRuneInString(str)
	if inspector.charSet[firstRune] {
		return true, str[runeSize:]
	}
	return false, ""
}

func (inspector SetInspector) Copy() Inspector {
	newCharSet := make(map[rune]bool)
	for k := range inspector.charSet {
		newCharSet[k] = true
	}
	return SetInspector{charSet: newCharSet}
}
