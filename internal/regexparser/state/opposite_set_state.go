package state

import "unicode/utf8"

type OppositeSetState struct {
	oppositeCharSet map[rune]bool
	NextStates      []*State
	PreviousStates  []*State
}

func NewOppositeSetState(chars []rune) SetState {
	charSet := make(map[rune]bool)
	for _, char := range chars {
		charSet[char] = true
	}
	return SetState{charSet, nil, nil}
}

func (state OppositeSetState) GetNextStates() []*State {
	return state.NextStates
}

func (state OppositeSetState) GetPreviousStates() []*State {
	return state.PreviousStates
}

func (state OppositeSetState) Match(str string) (bool, string) {
	if len(str) == 0 {
		return false, ""
	}
	firstRune, runeSize := utf8.DecodeRuneInString(str)
	if state.oppositeCharSet[firstRune] {
		return false, ""
	}
	return true, str[runeSize:]
}
