package state

func NewStartingState() State {
	newStartingInspector := StartingInspector{}
	return State{StateInspector: newStartingInspector}
}

func NewFinalState() State {
	newFinalInspector := FinalInspector{}
	return State{StateInspector: newFinalInspector}
}

func NewCharState(char rune) State {
	charInspector := CharInspector{char: char}
	return State{StateInspector: charInspector}
}
