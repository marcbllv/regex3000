package regexparser

import "fmt"

func BuildStateMachine(regex string) *State {
	startingState := NewStartingState()
	finalState := NewFinalState()
	return buildStateMachineFromStartAndFinalStates(regex, &startingState, &finalState)
}

func buildStateMachineFromStartAndFinalStates(regex string, startingState *State, finalState *State) *State {
	// TODO:
	//  - -- OK simple concatenation
	//  - pipe operator
	//  - ? * + operators
	//  - braces
	//  - parentheses
	var currentState *State
	var newState *State

	currentState = startingState
	for pos, char := range regex {
		if char == '|' {
			rightSideRegex := regex[pos + 1:]
			buildStateMachineFromStartAndFinalStates(rightSideRegex, startingState, finalState)
			break
		} else {
			charState := NewState(char)
			newState = &charState
			currentState.AppendNextState(newState)
			currentState = newState
		}
	}
	currentState.AppendNextState(finalState)
	return startingState
}

func DisplayStateMachine(stateMachine *State, i int) {
	fmt.Println(i)
	fmt.Printf("State: %p %v\n ", stateMachine, stateMachine)
	if len(stateMachine.NextStates) == 0 {
		fmt.Println("No next states, ending.")
		return
	}

	for _, nextState := range stateMachine.NextStates {
		DisplayStateMachine(nextState, i + 1)
	}
}
