package regexparser

type State struct {
	Char rune
	StateType int
	NextStates []*State
	PreviousStates []*State
}


const (
	StartingState = iota
	FinalState = iota
	ConcatState = iota
	EpsilonState = iota
)


func (state *State) AppendNextState(nextState *State) {
	state.NextStates = append(state.NextStates, nextState)
	nextState.PreviousStates = append(nextState.PreviousStates, state)
}


func (state *State) AppendNextStateToItself() {
	state.NextStates = append(state.NextStates, state)
	state.PreviousStates = append(state.PreviousStates, state)
}



func (state *State) AppendNextStates(nextStates []*State) {
	for _, s := range nextStates {
		state.AppendNextState(s)
	}
}


func NewState(char rune) State {
	return State{char, ConcatState, nil, nil}
}


func NewStartingState() State {
	return State{'!', StartingState, nil, nil}
}


func NewFinalState() State {
	return State{'_', FinalState, nil, nil}
}


func NewEpsilonState() State {
	return State{0, EpsilonState, nil, nil}
}
