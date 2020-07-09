package state

type FinalInspector struct {
	mustMatchEndOfString bool
}

func (inspector FinalInspector) Match(str []rune, pos int) []int {
	if pos > len(str) {
		return []int{}
	}
	if inspector.mustMatchEndOfString && pos < len(str) {
		return []int{}
	}
	return []int{pos}
}

func (inspector FinalInspector) Copy() Inspector {
	return FinalInspector{}
}
