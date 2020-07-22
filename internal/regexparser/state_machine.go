package regexparser

import (
	"fmt"

	"github.com/marcbllv/regex3000/internal/regexparser/parse"

	"github.com/marcbllv/regex3000/internal/regexparser/state"
)

func BuildStateMachine(regex string) *state.State {
	regexRunes := []rune(regex)
	matchBeginning, regexRunes := getMatchBeginning(regexRunes)
	matchEnd, regexRunes := getMatchEnd(regexRunes)
	startingState := state.NewStartingState(matchBeginning)
	finalState := state.NewFinalState(matchEnd)
	return buildStateMachine(regexRunes, &startingState, &finalState)
}

func buildStateMachine(regex []rune, startingState *state.State, finalState *state.State) *state.State {
	var currentState *state.State

	currentState = startingState
	pos := 0
	for pos < len(regex) {
		char := regex[pos]
		switch char {
		case '\\':
			currentState, pos = applyBackslashEscaping(regex, pos, currentState)
		case '|':
			buildStateMachine(regex[pos+1:], startingState, finalState)
			pos = len(regex) // to end the for loop
		case '?':
			currentState = applyQuestionMarkOperator(currentState)
			pos++
		case '+':
			currentState = applyPlusOperator(currentState)
			pos++
		case '*':
			currentState = applyStarOperator(currentState)
			pos++
		case '.':
			currentState = applyMatchAny(currentState)
			pos++
		case '(':
			currentState, pos = buildParenthesesContentStates(currentState, regex, pos)
		case '{':
			currentState, pos = buildNewBracesStates(currentState, regex, pos)
		case '[':
			currentState, pos = buildNewSetState(currentState, regex, pos)
		default:
			currentState = createAncConcatNewCharState(currentState, char)
			pos++
		}
	}
	currentState.AppendNextState(finalState)
	return startingState
}

func getMatchBeginning(regex []rune) (bool, []rune) {
	if regex[0] == '^' {
		return true, regex[1:]
	}
	return false, regex
}

func getMatchEnd(regex []rune) (bool, []rune) {
	if regex[len(regex)-1] == '$' {
		return true, regex[:len(regex)-1]
	}
	return false, regex
}

func createAncConcatNewCharState(currentState *state.State, char rune) *state.State {
	charState := state.NewCharState(char)
	currentState.AppendNextState(&charState)
	return &charState
}

func buildParenthesesContentStates(currentState *state.State, regex []rune, openParPosition int) (*state.State, int) {
	openParState, closingParState := state.NewParenthesesStates()
	rightParenthesis := findMatchingParenthesis(regex, openParPosition)
	innerContent := regex[openParPosition+1 : rightParenthesis]
	buildStateMachine(innerContent, openParState, closingParState)
	currentState.AppendNextState(openParState)
	return closingParState, rightParenthesis + 1
}

func buildNewBracesStates(currentState *state.State, regex []rune, openBracePosition int) (*state.State, int) {
	rangeQuantifier, currentPosition := parse.ParseBraces(regex, openBracePosition)
	initialState := currentState
	for i := 0; i < rangeQuantifier.High-1; i++ {
		copiedStarting := currentState.Copy()
		currentState.AppendNextState(copiedStarting)
		if copiedStarting.MatchingState == nil {
			currentState = copiedStarting
		} else {
			currentState = copiedStarting.MatchingState
		}

		if i < rangeQuantifier.High-rangeQuantifier.Low {
			for _, prevInitState := range initialState.GetPreviousStates() {
				prevInitState.AppendNextState(currentState)
			}
		}
	}
	return currentState, currentPosition
}

func applyQuestionMarkOperator(currentState *state.State) *state.State {
	epsilonState := state.NewEpsilonState()
	newState := &epsilonState
	currentState.AppendNextState(newState)

	if currentState.PreviousStates == nil && currentState.MatchingState == nil {
		// TODO handle errors properly
		return nil
	}

	var previousStates []*state.State
	if currentState.MatchingState == nil {
		previousStates = currentState.PreviousStates
	} else {
		previousStates = []*state.State{currentState.MatchingState}
	}
	for _, previousState := range previousStates {
		previousState.AppendNextState(newState)
	}
	return &epsilonState
}

func applyPlusOperator(currentState *state.State) *state.State {
	if currentState == nil {
		return nil // TODO: handle error properly
	}

	if currentState.MatchingState == nil {
		currentState.AppendStateToItself()
	} else {
		currentState.AppendNextState(currentState.MatchingState)
	}
	return currentState
}

func applyStarOperator(currentState *state.State) *state.State {
	applyPlusOperator(currentState)
	epsilonState := applyQuestionMarkOperator(currentState)
	return epsilonState
}

func DisplayStateMachine(stateMachine *state.State, i int) {
	fmt.Println(i)
	fmt.Printf("State: %p %v\n ", stateMachine, stateMachine)
	nextStates := stateMachine.GetNextStates()
	if len(nextStates) == 0 {
		fmt.Println("No next states, ending.")
		return
	}

	for _, nextState := range nextStates {
		if nextState != stateMachine {
			DisplayStateMachine(nextState, i+1)
		}
	}
}

func applyMatchAny(currentState *state.State) *state.State {
	newState := state.NewOppositeSetState([]rune{})
	currentState.AppendNextState(&newState)
	return &newState
}

func buildNewSetState(currentState *state.State, regex []rune, pos int) (*state.State, int) {
	rightBracket := findMatchingBracket(regex, pos)
	innerContent := regex[pos+1 : rightBracket]
	charSet, isOppositeSet := parseBracket(innerContent)

	var newState state.State
	if isOppositeSet {
		newState = state.NewOppositeSetState(charSet)
	} else {
		newState = state.NewSetState(charSet)

	}
	currentState.AppendNextState(&newState)
	currentState = &newState
	return currentState, rightBracket + 1
}

func applyBackslashEscaping(regex []rune, pos int, currentState *state.State) (*state.State, int) {
	if pos >= len(regex)-1 {
		panic("Unexpected EOF after \\.")
	}
	pos++
	escapedChar := regex[pos]
	pos++
	currentState = createAncConcatNewCharState(currentState, escapedChar)
	return currentState, pos
}
