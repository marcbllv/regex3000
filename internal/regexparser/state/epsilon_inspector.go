package state

type EpsilonInspector struct{}

func (inspector EpsilonInspector) Match(str []rune, pos int) []int {
	if pos > len(str) {
		return []int{}
	}
	return []int{pos}
}

func (inspector EpsilonInspector) Copy() Inspector {
	return EpsilonInspector{}
}
