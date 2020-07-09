package state

type FinalInspector struct {
	matchEnd bool
}

func (inspector FinalInspector) Match(str []rune, pos int) []int {
	if pos > len(str) {
		return []int{}
	}
	if inspector.matchEnd && pos < len(str) {
		return []int{}
	}
	return []int{pos}
}

func (inspector FinalInspector) Copy() Inspector {
	return FinalInspector{}
}
