package regexparser


type BinaryNode struct {
    Value string
    LeftChild *BinaryNode
    RightChild *BinaryNode
}

func NewBinaryNode(value string) BinaryNode {
    return BinaryNode{value, nil, nil}
}

func isLeafBinaryNode(tree BinaryNode) bool {
    return (tree.LeftChild == nil) && (tree.RightChild == nil)
}