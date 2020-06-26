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
	//  - -- OK pipe operator
	//  - -- OK ? * + operators
	//  - -- OK parentheses
	//  - braces
	var currentState *State
	var newState *State

	currentState = startingState
	pos := 0
	for pos < len(regex) {
		char := rune(regex[pos])
		if char == '|' {
			rightSideRegex := regex[pos + 1:]
			buildStateMachineFromStartAndFinalStates(rightSideRegex, startingState, finalState)
			break
		} else if char == '?' {
			if currentState == nil || currentState.StateType == EpsilonState {
				return nil  // TODO: handle error properly
			}
			currentState = applyQuestionMarkOperator(currentState)
		} else if char == '+' {
			if currentState == nil || currentState.StateType == EpsilonState {
				return nil  // TODO: handle error properly
			}
			currentState = applyPlusOperator(currentState)
		} else if char == '*' {
			if currentState == nil || currentState.StateType == EpsilonState {
				return nil  // TODO: handle error properly
			}
			currentState = applyStarOperator(currentState)
		} else if char == '(' {
			openParState, closingParState := NewParenthesesStates()
			rightParenthesis := findMatchingParenthesis(regex, pos)
			innerContent := regex[pos + 1:rightParenthesis]
			buildStateMachineFromStartAndFinalStates(innerContent, openParState, closingParState)
			currentState.AppendNextState(openParState)
			currentState = closingParState
			pos = rightParenthesis
		} else {
			charState := NewState(char)
			newState = &charState
			currentState.AppendNextState(newState)
			currentState = newState
		}
		pos++
	}
	currentState.AppendNextState(finalState)
	return startingState
}


func applyQuestionMarkOperator(currentState *State) *State {
	epsilonState := NewEpsilonState()
	newState := &epsilonState
	currentState.AppendNextState(newState)

	var previousStates []*State
	if currentState.matchingState == nil {
		previousStates = currentState.PreviousStates
	} else {
		previousStates = []*State{currentState.matchingState}
	}
	for _, previousState := range previousStates {
		previousState.AppendNextState(newState)
	}
	return &epsilonState
}


func applyPlusOperator(currentState *State) *State {
	if currentState.matchingState == nil {
		currentState.AppendNextStateToItself()
	} else {
		currentState.AppendNextState(currentState.matchingState)
	}
	return currentState
}


func applyStarOperator(currentState *State) *State {
	applyPlusOperator(currentState)
	epsilonState := applyQuestionMarkOperator(currentState)
	return epsilonState
}



func DisplayStateMachine(stateMachine *State, i int) {
	fmt.Println(i)
	fmt.Printf("State: %p %v\n ", stateMachine, stateMachine)
	if len(stateMachine.NextStates) == 0 {
		fmt.Println("No next states, ending.")
		return
	}

	for _, nextState := range stateMachine.NextStates {
		if nextState != stateMachine {
			DisplayStateMachine(nextState, i+1)
		}
	}
}
