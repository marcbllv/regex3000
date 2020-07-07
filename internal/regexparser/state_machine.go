package regexparser

import (
	"fmt"

	"github.com/marcbllv/regex3000/internal/regexparser/state"
)

func BuildStateMachine(regex string) *state.State {
	startingState := state.NewStartingState()
	finalState := state.NewFinalState()
	return buildStateMachine(regex, &startingState, &finalState)
}

func buildStateMachine(regex string, startingState *state.State, finalState *state.State) *state.State {
	var currentState *state.State

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
			buildStateMachine(rightSideRegex, startingState, finalState)
			pos = len(regex) // break 'for' loop
		default:
			currentState = createAncConcatNewCharState(currentState, char)
		}
		pos++
	}
	currentState.AppendNextState(finalState)
	return startingState
}

func createAncConcatNewCharState(currentState *state.State, char rune) *state.State {
	charState := state.NewCharState(char)
	currentState.AppendNextState(&charState)
	return &charState
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
