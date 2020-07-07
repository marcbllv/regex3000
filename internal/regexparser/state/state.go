package state

type State struct {
	NextStates     []*State
	PreviousStates []*State
	MatchingState  *State
	StateInspector Inspector
}

func (state State) GetNextStates() []*State {
	return state.NextStates
}

func (state State) GetPreviousStates() []*State {
	return state.PreviousStates
}

func (state *State) AppendNextState(newState *State) {
	state.NextStates = append(state.NextStates, newState)
	newState.PreviousStates = append(newState.PreviousStates, state)
}

func (state *State) AppendStateToItself() {
	state.NextStates = append(state.NextStates, state)
	state.PreviousStates = append(state.PreviousStates, state)
}

func (state *State) Copy() State {
	newInspector := state.StateInspector.Copy()
	return State{NextStates: nil, PreviousStates: nil, MatchingState: nil, StateInspector: newInspector}
}

func (state *State) Match(str string) (bool, string) {
	return state.StateInspector.Match(str)
}

func NewStartingState() State {
	newStartingInspector := StartingInspector{}
	return State{StateInspector: newStartingInspector}
}

func NewFinalState() State {
	newFinalInspector := FinalInspector{mustMatchEndOfString: false}
	return State{StateInspector: newFinalInspector}
}

func NewCharState(char rune) State {
	charInspector := CharInspector{char: char}
	return State{StateInspector: charInspector}
}
