package regexparser

type LegacyState struct {
	Char           rune
	CharSet        map[rune]bool
	StateType      int
	NextStates     []*LegacyState
	PreviousStates []*LegacyState
	matchingState  *LegacyState
}

const (
	StartingStateType = iota
	FinalStateType    = iota
	CharStateType     = iota
	MatchAnyStateType = iota
	EpsilonStateType  = iota
)

func (state *LegacyState) AppendNextState(nextState *LegacyState) {
	state.NextStates = append(state.NextStates, nextState)
	nextState.PreviousStates = append(nextState.PreviousStates, state)
}

func (state *LegacyState) AppendNextStateToItself() {
	state.NextStates = append(state.NextStates, state)
	state.PreviousStates = append(state.PreviousStates, state)
}

func (state *LegacyState) AppendNextStates(nextStates []*LegacyState) {
	for _, s := range nextStates {
		state.AppendNextState(s)
	}
}

func NewState(char rune) LegacyState {
	charSet := make(map[rune]bool)
	charSet[char] = true
	return LegacyState{char, charSet, CharStateType, nil, nil, nil}
}

func NewStartingState() LegacyState {
	return LegacyState{'!', nil, StartingStateType, nil, nil, nil}
}

func NewFinalState() LegacyState {
	return LegacyState{'_', nil, FinalStateType, nil, nil, nil}
}

func NewStateMatchAny() LegacyState {
	return LegacyState{'.', nil, MatchAnyStateType, nil, nil, nil}
}

func NewStateCustomType(char rune, charSet map[rune]bool, stateType int) LegacyState {
	return LegacyState{char, charSet, stateType, nil, nil, nil}
}

func CopyState(state *LegacyState) *LegacyState {
	var charSet map[rune]bool
	if state.CharSet != nil {
		charSet = make(map[rune]bool)
		for k, v := range state.CharSet {
			charSet[k] = v
		}
	}
	newState := LegacyState{state.Char, charSet, state.StateType, nil, nil, nil}
	return &newState
}
