package state

type EpsilonState struct {
	NextStates     []*State
	PreviousStates []*State
}

func NewEpsilonState(char rune) EpsilonState {
	return EpsilonState{nil, nil}
}

func (state EpsilonState) GetNextStates() []*State {
	return state.NextStates
}

func (state EpsilonState) GetPreviousStates() []*State {
	return state.PreviousStates
}

func (state EpsilonState) Match(str string) (bool, string) {
	return true, str
}
