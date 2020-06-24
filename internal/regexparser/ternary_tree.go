package regexparser


type TernaryNode struct {
	Value string
	Start int
	End int
	LeftChild *TernaryNode
	MiddleChild *TernaryNode
	RightChild *TernaryNode
}


func NewTernaryNode(value string, start int, end int) TernaryNode {
	return TernaryNode{value, start, end, nil, nil, nil}
}


func NewRootTernaryNode(value string) TernaryNode {
	return TernaryNode{value, 0, len(value) - 1, nil, nil, nil}
}


func NewLeafTernaryNode(value string, pos int) TernaryNode {
	return TernaryNode{value, pos, pos + 1, nil, nil, nil}
}


func isLeafTernaryNode(node TernaryNode) bool {
	return (node.LeftChild == nil) && (node.MiddleChild == nil) && (node.RightChild == nil)
}


func getChildren(tree TernaryNode) []*TernaryNode {
	return []*TernaryNode{tree.LeftChild, tree.MiddleChild, tree.RightChild}
}