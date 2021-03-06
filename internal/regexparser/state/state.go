package state

type State struct {
	NextStates     []*State
	PreviousStates []*State
	MatchingState  *State
	StateInspector Inspector
	IsFinalState   bool
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

func (state *State) CopyState() *State {
	newInspector := state.StateInspector.Copy()
	return &State{NextStates: nil, PreviousStates: nil, MatchingState: nil, StateInspector: newInspector}
}

func (state *State) Copy() *State {
	var copiedState *State
	if state.MatchingState == nil {
		copiedState = state.CopyState()
	} else {
		openParenthesisState := state.MatchingState
		copiedStatesCache := make(map[*State]*State)
		copiedState = recursiveStatesCopy(openParenthesisState, copiedStatesCache, state)
	}
	return copiedState
}

func recursiveStatesCopy(state *State, copiedStatesCache map[*State]*State, endState *State) *State {
	if copiedStatesCache[state] == nil {
		copiedStatesCache[state] = state.CopyState()
	}
	copiedState := copiedStatesCache[state]
	if state == endState {
		return copiedState
	}

	for _, nextState := range state.NextStates {
		copiedNextState := recursiveStatesCopy(nextState, copiedStatesCache, endState)
		copiedState.AppendNextState(copiedNextState)
	}
	if state.MatchingState != nil {
		copiedMatching := copiedStatesCache[state.MatchingState]
		if copiedMatching == nil {
			// TODO: return error, the matching state is supposed to have been copied
		} else {
			copiedState.MatchingState = copiedMatching
			copiedMatching.MatchingState = copiedState
		}
	}
	return copiedState
}

func (state *State) Match(str []rune, pos int) bool {
	matchPositions := state.StateInspector.Match(str, pos)
	if len(matchPositions) == 0 {
		return false
	}

	nextStates := state.GetNextStates()
	if len(nextStates) == 0 {
		return state.IsFinalState
	}
	for _, nextPos := range matchPositions {
		for _, nextState := range nextStates {
			if nextState.Match(str, nextPos) {
				return true
			}
		}
	}
	return false
}

func NewStartingState(matchBeginning bool) State {
	newStartingInspector := StartingInspector{mustMatchBeginningOfString: matchBeginning}
	return State{StateInspector: newStartingInspector}
}

func NewFinalState(matchEnd bool) State {
	newFinalInspector := FinalInspector{mustMatchEndOfString: matchEnd}
	return State{StateInspector: newFinalInspector, IsFinalState: true}
}

func NewEpsilonState() State {
	inspector := EpsilonInspector{}
	return State{StateInspector: inspector}
}

func NewParenthesesStates() (*State, *State) {
	openParState := NewEpsilonState()
	closingParState := NewEpsilonState()
	openParState.MatchingState = &closingParState
	closingParState.MatchingState = &openParState
	return &openParState, &closingParState
}

func NewCharState(char rune) State {
	charInspector := CharInspector{char: char}
	return State{StateInspector: charInspector}
}

func NewSetState(charsSet []rune) State {
	setInspector := NewSetInspector(charsSet)
	return State{StateInspector: setInspector}
}

func NewOppositeSetState(oppositeChars []rune) State {
	oppositeSetInspector := NewOppositeSetInspector(oppositeChars)
	return State{StateInspector: oppositeSetInspector}
}
