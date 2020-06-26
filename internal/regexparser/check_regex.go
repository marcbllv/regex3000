package regexparser

import "fmt"

func CheckRegexMatch(regex string, str string) bool {
	stateMachine := BuildStateMachine(regex)
	return checkStateMachine(stateMachine, str)
}


func checkStateMachine(stateMachine *State, str string) bool {
	if stateMachine.StateType == FinalState {
		return len(str) == 0
	} else if stateMachine.StateType == StartingState {
		for _, nextState := range stateMachine.NextStates {
			if checkStateMachine(nextState, str) {
				return true
			}
		}
		return false
	} else {
		// Simple concatenation
		if len(str) == 0 || stateMachine.Char != rune(str[0]) {
			return false
		}
		for _, nextState := range stateMachine.NextStates {
			return checkStateMachine(nextState, str[1:])
		}
	}
	fmt.Println("error")
	return false
}
