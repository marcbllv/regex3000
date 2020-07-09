package state

type SetInspector struct {
	charSet map[rune]bool
}

func NewSetInspector(chars []rune) SetInspector {
	charsSet := make(map[rune]bool)
	for _, char := range chars {
		charsSet[char] = true
	}
	return SetInspector{charsSet}
}

func (inspector SetInspector) Match(str []rune, pos int) []int {
	if pos >= len(str) {
		return []int{}
	}
	if inspector.charSet[str[pos]] {
		return []int{pos + 1}
	}
	return []int{}
}

func (inspector SetInspector) Copy() Inspector {
	newCharSet := make(map[rune]bool)
	for k := range inspector.charSet {
		newCharSet[k] = true
	}
	return SetInspector{charSet: newCharSet}
}
