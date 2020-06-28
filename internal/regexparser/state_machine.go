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
	//  - WIP braces
	//  - brackets
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
			var copiedStarting *State
			var copiedFinal *State
			var initialState *State

			rightBrace := findMatchingBrace(regex, pos)
			innerContent := regex[pos + 1:rightBrace]
			min, max := parseBraceContent(innerContent)
			initialState = currentState
			for i := 0; i < max - 1; i++ {
				copiedStarting, copiedFinal = duplicateLastState(currentState)
				currentState.AppendNextState(copiedStarting)
				currentState = copiedFinal
				if i < max - min {
					initialState.AppendNextState(currentState)
				}
			}
			pos = rightBrace
		} else if char == '['{
			rightBracket := findMatchingBracket(regex, pos)
			innerContent := regex[pos + 1:rightBracket]
			charSet := parseBracket(innerContent)
			pos = rightBracket

			state := NewSetState(charSet)
			newState = &state
			currentState.AppendNextState(newState)
			currentState = newState
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
		copiedOpeningPar := CopyState(currentState.matchingState)
		duplicateParenthesesBlock(
			currentState.matchingState,
			copiedOpeningPar,
			currentState,
			copiedOpeningPar)
		return copiedOpeningPar, copiedOpeningPar.matchingState
	}
}


func duplicateParenthesesBlock(
	currentState *State,
	copiedCurrState *State,
	closingParState *State,
	copiedOpeningPar *State) {
	if currentState == closingParState {
		// Plug closing par with opening par
		copiedOpeningPar.matchingState = copiedCurrState
		copiedCurrState.matchingState = copiedOpeningPar
	}

	for _, nextState := range currentState.NextStates {
		char := nextState.Char
		charSet := nextState.CharSet
		stateType := nextState.StateType
		copiedNextState := NewStateCustomType(char, charSet, stateType)
		copiedCurrState.AppendNextState(&copiedNextState)
		duplicateParenthesesBlock(nextState, &copiedNextState, closingParState, copiedOpeningPar)
	}
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
