package state

type StartingInspector struct {
	mustMatchBeginningOfString bool
}

func (inspector StartingInspector) Match(str []rune, pos int) []int {
	if pos != 0 {
		return []int{}
	}

	if inspector.mustMatchBeginningOfString {
		return []int{0}
	}
	matchingPositions := make([]int, len(str))
	for i := 0; i < len(str); i++ {
		matchingPositions[i] = i
	}
	return matchingPositions
}

func (inspector StartingInspector) Copy() Inspector {
	return StartingInspector{}
}
