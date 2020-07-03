package state

import (
	"unicode/utf8"
)

type SetState struct {
	charSet        map[rune]bool
	NextStates     []*State
	PreviousStates []*State
}

func NewSetState(chars []rune) SetState {
	charSet := make(map[rune]bool)
	for _, char := range chars {
		charSet[char] = true
	}
	return SetState{charSet, nil, nil}
}

func (state SetState) GetNextStates() []*State {
	return state.NextStates
}

func (state SetState) GetPreviousStates() []*State {
	return state.PreviousStates
}

func (state SetState) Match(str string) (bool, string) {
	if len(str) == 0 {
		return false, ""
	}
	firstRune, runeSize := utf8.DecodeRuneInString(str)
	if state.charSet[firstRune] {
		return true, str[runeSize:]
	}
	return false, ""
}
