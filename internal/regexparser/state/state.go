package state

type State interface {
	Match(str string) (bool, string)
	GetNextStates() []*State
	GetPreviousStates() []*State
	AppendNextState(state *State)
	AppendStateToItself()
	Copy() *State
}

// Starting state
type StartingState struct {
	mustMatchFirstChar bool
	NextStates         []*State
}

func NewStartingState() StartingState {
	return StartingState{nil}
}

func (state StartingState) GetNextStates() []*State {
	return state.NextStates
}

func (state StartingState) GetPreviousStates() []*State {
	return []*State{}
}

// Final state
type FinalState struct {
	PreviousStates []*State
}

func NewFinalState() FinalState {
	return FinalState{nil}
}

func (state FinalState) GetNextStates() []*State {
	return []*State{}
}

func (state FinalState) GetPreviousStates() []*State {
	return state.PreviousStates
}
