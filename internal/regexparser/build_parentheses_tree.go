package regexparser


func findMatchingParenthese(s string) int {
    count := 0
    for pos, char := range s {
        if char == '(' {
            count++
        } else if char == ')' {
            if count == 0 {
                return pos
            } else {
                count--
            }
        }
    }
    // error
    return -1
}


func buildParenthesesTree(regex string) *TernaryNode {
    if len(regex) == 0 {
        return nil
    }

    var tree = NewRootTernaryNode(regex)
    for pos, char := range(regex) {
        if char == '(' {
            rightPar := findMatchingParenthese(regex[pos + 1:])

            leftSubstr := regex[:pos]
            middleSubstr := regex[pos + 1:rightPar]
            rightSubstr := regex[rightPar + 1:]

            leftNode := buildParenthesesSubTree(leftSubstr)
            middleNode := buildParenthesesSubTree(middleSubstr)
            rightNode := buildParenthesesSubTree(rightSubstr)

            tree.LeftChild = leftNode
            tree.MiddleChild = middleNode
            tree.RightChild = rightNode
            break
        }
    }
    return &tree
}
