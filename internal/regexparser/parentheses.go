package regexparser

func findMatchingParenthese(s string, openParPos int) int {
    count := 0
    substr := s[openParPos+ 1:]

    for pos, char := range substr {
        if char == '(' {
            count++
        } else if char == ')' {
            if count > 0 {
                count--
            } else {
                return openParPos + pos + 1
            }
        }
    }
    // error
    // todo: Raise error properly
    return -1
}


func BuildParenthesesTree(regex string) *TernaryNode {
    var tree = NewRootTernaryNode(regex)

    for pos, char := range regex {
        if char == '(' {
            closingParenthesis := findMatchingParenthese(regex, pos)
            tree.LeftChild = BuildParenthesesTree(regex[:pos])
            tree.MiddleChild = BuildParenthesesTree(regex[pos + 1: closingParenthesis])
            tree.RightChild = BuildParenthesesTree(regex[closingParenthesis+ 1:])
            break
        }
    }
    return &tree
}
