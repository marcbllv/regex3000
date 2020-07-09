package state

type CharInspector struct {
	char rune
}

func (inspector CharInspector) Match(str []rune, pos int) []int {
	if pos >= len(str) {
		return []int{}
	}
	if inspector.char == str[pos] {
		return []int{pos + 1}
	}
	return []int{}
}

func (inspector CharInspector) Copy() Inspector {
	return CharInspector{char: inspector.char}
}
