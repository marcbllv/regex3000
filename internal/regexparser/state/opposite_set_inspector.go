package state

type OppositeSetInspector struct {
	oppositeCharSet map[rune]bool
}

func NewOppositeSetInspector(chars []rune) OppositeSetInspector {
	oppositeCharSet := make(map[rune]bool)
	for _, char := range chars {
		oppositeCharSet[char] = true
	}
	return OppositeSetInspector{oppositeCharSet}
}

func (inspector OppositeSetInspector) Match(str []rune, pos int) []int {
	if pos >= len(str) {
		return []int{}
	}
	if inspector.oppositeCharSet[str[pos]] {
		return nil
	}
	return []int{pos + 1}
}

func (inspector OppositeSetInspector) Copy() Inspector {
	newOppositeCharSet := make(map[rune]bool)
	for k := range inspector.oppositeCharSet {
		newOppositeCharSet[k] = true
	}
	return OppositeSetInspector{oppositeCharSet: newOppositeCharSet}
}
