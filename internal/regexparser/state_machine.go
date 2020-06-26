package regexparser

import (
	"fmt"
)

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
		} else if char == '{' {
			rightBrace := findMatchingBrace(regex, pos)
			innerContent := regex[pos + 1:rightBrace]
			min, max := parseBraceContent(innerContent)

			var copiedStarting *State
			var copiedFinal *State
			var previousCurrentState *State
			for i := 0; i < min - 1; i++ {
				copiedStarting, copiedFinal = duplicateLastState(currentState)
				currentState.AppendNextState(copiedStarting)
				previousCurrentState = currentState
				currentState = copiedFinal
			}
			for i := min; i < max; i++ {
				copiedStarting, copiedFinal = duplicateLastState(currentState)
				previousCurrentState.AppendNextState(copiedStarting)
				for _, prevState := range copiedFinal.PreviousStates {
					prevState.AppendNextState(currentState)
				}
			}
			pos = rightBrace
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


func duplicateLastState(currentState *State) (*State, *State){
	// If currentState is ')', it duplicates the whole parentheses block
	// Returns the duplicated starting state & closing state of the block
	if currentState.matchingState == nil {
		copied := CopyState(currentState)
		return copied, copied
	} else {
		copiedClosingPar := CopyState(currentState.matchingState)
		return duplicateParenthesesBlock(currentState, copiedClosingPar)
	}
}


func duplicateParenthesesBlock(currentState *State, copiedClosingPar *State) (*State, *State){
	if currentState == copiedClosingPar {
		return copiedClosingPar, copiedClosingPar
	}

	char := currentState.Char
	stateType := currentState.StateType
	newState := NewStateCustomType(char, stateType)

	for _, nextState := range currentState.NextStates {
		duplicatedNextState, _ := duplicateParenthesesBlock(nextState, copiedClosingPar)
		newState.AppendNextState(duplicatedNextState)
	}
	return &newState, copiedClosingPar
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
