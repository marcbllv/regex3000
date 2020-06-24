package regexparser

func buildDisjunctionSubTree(regex string) *BinaryNode {
	// Build tree for operator '|'
	tree := NewBinaryNode(regex)
	if len(regex) == 0 {
		return &tree
	}

	for pos, char := range regex {
		if char == '|' {
			tree.LeftChild = buildDisjunctionSubTree(regex[:pos])
			tree.RightChild = buildDisjunctionSubTree(regex[pos + 1:])
			break
		}
	}
	return &tree
}


func buildDisjunctionTree(regex string) BinaryNode {
	return *buildDisjunctionSubTree(regex)
}