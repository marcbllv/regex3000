package regexparser

import (
	"fmt"
)

func BuildStateMachine(regex string) *LegacyState {
	startingState := NewStartingState()
	finalState := NewFinalState()
	return buildStateMachineFromStartAndFinalStates(regex, &startingState, &finalState)
}

func buildStateMachineFromStartAndFinalStates(regex string, startingState *LegacyState, finalState *LegacyState) *LegacyState {
	var currentState *LegacyState
	var newState *LegacyState

	currentState = startingState
	pos := 0
	for pos < len(regex) {
		char := rune(regex[pos])
		switch char {
		case '\\':
			if pos < len(regex)-1 {
				escapedChar := rune(regex[pos+1])
				currentState = createAncConcatNewCharState(currentState, escapedChar)
				pos++
			}
			// todo: raise error if pos == len(regex) - 1
		case '|':
			rightSideRegex := regex[pos+1:]
			buildStateMachineFromStartAndFinalStates(rightSideRegex, startingState, finalState)
			pos = len(regex) // break 'for' loop
		case '?':
			currentState = applyQuestionMarkOperator(currentState)
		case '+':
			currentState = applyPlusOperator(currentState)
		case '*':
			if currentState == nil || currentState.StateType == EpsilonStateType {
				return nil // TODO: handle error properly
			}
			currentState = applyStarOperator(currentState)
		case '.':
			currentState = applyMatchAny(currentState)
		case '(':
			openParState, closingParState := NewParenthesesStates()
			rightParenthesis := findMatchingParenthesis(regex, pos)
			innerContent := regex[pos+1 : rightParenthesis]
			buildStateMachineFromStartAndFinalStates(innerContent, openParState, closingParState)
			currentState.AppendNextState(openParState)
			currentState = closingParState
			pos = rightParenthesis
		case '{':
			var copiedStarting *LegacyState
			var copiedFinal *LegacyState
			var initialState *LegacyState

			rightBrace := findMatchingBrace(regex, pos)
			innerContent := regex[pos+1 : rightBrace]
			min, max := parseBraceContent(innerContent)
			initialState = currentState
			for i := 0; i < max-1; i++ {
				copiedStarting, copiedFinal = duplicateLastState(currentState)
				currentState.AppendNextState(copiedStarting)
				currentState = copiedFinal
				if i < max-min {
					initialState.AppendNextState(currentState)
				}
			}
			pos = rightBrace
		case '[':
			rightBracket := findMatchingBracket(regex, pos)
			innerContent := regex[pos+1 : rightBracket]
			charSet := parseBracket(innerContent)
			pos = rightBracket

			state := NewSetState(charSet)
			newState = &state
			currentState.AppendNextState(newState)
			currentState = newState
		default:
			currentState = createAncConcatNewCharState(currentState, char)
		}
		pos++
	}
	currentState.AppendNextState(finalState)
	return startingState
}

func createAncConcatNewCharState(currentState *LegacyState, char rune) *LegacyState {
	charState := NewState(char)
	newState := &charState
	currentState.AppendNextState(newState)
	return newState
}

func applyQuestionMarkOperator(currentState *LegacyState) *LegacyState {
	epsilonState := NewEpsilonState()
	newState := &epsilonState
	currentState.AppendNextState(newState)

	if currentState.PreviousStates == nil && currentState.matchingState == nil {
		// TODO handle errors properly
		return nil
	}

	var previousStates []*LegacyState
	if currentState.matchingState == nil {
		previousStates = currentState.PreviousStates
	} else {
		previousStates = []*LegacyState{currentState.matchingState}
	}
	for _, previousState := range previousStates {
		previousState.AppendNextState(newState)
	}
	return &epsilonState
}

func applyPlusOperator(currentState *LegacyState) *LegacyState {
	if currentState == nil {
		return nil // TODO: handle error properly
	}

	if currentState.matchingState == nil {
		currentState.AppendNextStateToItself()
	} else {
		currentState.AppendNextState(currentState.matchingState)
	}
	return currentState
}

func applyStarOperator(currentState *LegacyState) *LegacyState {
	applyPlusOperator(currentState)
	epsilonState := applyQuestionMarkOperator(currentState)
	return epsilonState
}

func applyMatchAny(currentState *LegacyState) *LegacyState {
	newState := NewStateMatchAny()
	currentState.AppendNextState(&newState)
	return &newState
}

func duplicateLastState(currentState *LegacyState) (*LegacyState, *LegacyState) {
	// If currentState is ')', it duplicates the whole parentheses block
	// Returns the duplicated starting state & closing state of the block
	if currentState.matchingState == nil {
		copied := CopyState(currentState)
		return copied, copied
	}

	copiedOpeningPar := CopyState(currentState.matchingState)
	duplicateParenthesesBlock(
		currentState.matchingState,
		copiedOpeningPar,
		currentState,
		copiedOpeningPar)
	return copiedOpeningPar, copiedOpeningPar.matchingState
}

func duplicateParenthesesBlock(
	currentState *LegacyState,
	copiedCurrState *LegacyState,
	closingParState *LegacyState,
	copiedOpeningPar *LegacyState) {
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

func DisplayStateMachine(stateMachine *LegacyState, i int) {
	fmt.Println(i)
	fmt.Printf("LegacyState: %p %v\n ", stateMachine, stateMachine)
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
