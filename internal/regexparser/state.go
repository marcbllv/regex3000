package regexparser

type State struct {
	Char           rune
	CharSet        map[rune]bool
	StateType      int
	NextStates     []*State
	PreviousStates []*State
	matchingState  *State
}

const (
	StartingState = iota
	FinalState    = iota
	CharState     = iota
	MatchAnyState = iota
	EpsilonState  = iota
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
	charSet := make(map[rune]bool)
	charSet[char] = true
	return State{char, charSet, CharState, nil, nil, nil}
}

func NewStartingState() State {
	return State{'!', nil, StartingState, nil, nil, nil}
}

func NewFinalState() State {
	return State{'_', nil, FinalState, nil, nil, nil}
}

func NewStateMatchAny() State {
	return State{'.', nil, MatchAnyState, nil, nil, nil}
}

func NewEpsilonState() State {
	return State{0, nil, EpsilonState, nil, nil, nil}
}

func NewSetState(charSet []rune) State {
	charSetMap := make(map[rune]bool)
	for _, r := range charSet {
		charSetMap[r] = true
	}
	return State{0, charSetMap, CharState, nil, nil, nil}
}

func NewStateCustomType(char rune, charSet map[rune]bool, stateType int) State {
	return State{char, charSet, stateType, nil, nil, nil}
}

func CopyState(state *State) *State {
	var charSet map[rune]bool
	if state.CharSet != nil {
		charSet = make(map[rune]bool)
		for k, v := range state.CharSet {
			charSet[k] = v
		}
	}
	newState := State{state.Char, charSet, state.StateType, nil, nil, nil}
	return &newState
}

func NewParenthesesStates() (*State, *State) {
	var charSet map[rune]bool
	openParState := State{'(', charSet, EpsilonState, nil, nil, nil}
	closingParState := State{')', charSet, EpsilonState, nil, nil, nil}
	openParState.matchingState = &closingParState
	closingParState.matchingState = &openParState
	return &openParState, &closingParState
}
