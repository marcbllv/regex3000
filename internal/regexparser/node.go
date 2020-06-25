package regexparser

type Node struct {
	Value string
	Type string  // One of: | or empty string for concatenation
	RepeatMin int
	RepeatMax int
	Children []*Node
}


func NewNode(value string, nodeType string) Node {
	return Node{value, nodeType, 1, 1, nil}
}


func IsLeafNode(node Node) bool {
	for _, child := range node.Children {
		if child != nil {
			return false
		}
	}
	return true
}
