package state

import (
	"unicode/utf8"
)

type CharState struct {
	char           rune
	NextStates     []*State
	PreviousStates []*State
}

func NewCharState(char rune) CharState {
	return CharState{char, nil, nil}
}

func (state CharState) GetNextStates() []*State {
	return state.NextStates
}

func (state CharState) GetPreviousStates() []*State {
	return state.PreviousStates
}

func (state CharState) Match(str string) (bool, string) {
	if len(str) == 0 {
		return false, ""
	}
	firstRune, runeSize := utf8.DecodeRuneInString(str)
	if firstRune == state.char {
		return true, str[runeSize:]
	}
	return false, ""
}
