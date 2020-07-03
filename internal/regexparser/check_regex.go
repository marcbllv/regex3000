package regexparser

func CheckRegexMatch(regex string, str string) bool {
	stateMachine := BuildStateMachine(regex)
	return checkStateMachine(stateMachine, str)
}

func checkStateMachine(stateMachine *State, str string) bool {
	if stateMachine.StateType == FinalState {
		return matchEmptyString(str)
	} else if stateMachine.StateType == StartingState {
		return forwardFullStringNextStates(stateMachine, str)
	} else if stateMachine.StateType == MatchAnyState {
		if len(str) == 0 {
			return false
		}
		return forwardFullStringNextStates(stateMachine, str[1:])
	} else if stateMachine.StateType == EpsilonState {
		return forwardFullStringNextStates(stateMachine, str)
	} else {
		// Simple concatenation
		return matchFirstCharForwardRest(stateMachine, str)
	}
}

func forwardFullStringNextStates(state *State, str string) bool {
	for _, nextState := range state.NextStates {
		if checkStateMachine(nextState, str) {
			return true
		}
	}
	return false
}

func matchFirstCharForwardRest(state *State, str string) bool {
	if len(str) == 0 {
		return false
	}
	if !matchChar(state, rune(str[0])) {
		return false
	}

	for _, nextState := range state.NextStates {
		if checkStateMachine(nextState, str[1:]) {
			return true
		}
	}
	return false
}

func matchChar(state *State, char rune) bool {
	if state.StateType != CharState {
		return false
	}
	if char == state.Char {
		return true
	}
	exists := state.CharSet[char]
	return exists
}

func matchEmptyString(str string) bool {
	return len(str) == 0
}
