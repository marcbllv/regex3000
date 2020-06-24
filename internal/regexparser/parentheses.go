package regexparser

import "fmt"

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


func DisplayParenthesesTree(tree *TernaryNode) {
    fmt.Println("VALUE:", tree.Value)
    fmt.Print("Left child ")
    if tree.LeftChild != nil {
        fmt.Println(tree.LeftChild.Value)
        DisplayParenthesesTree(tree.LeftChild)
    } else {
        fmt.Println("no left child")
    }
    fmt.Print("Middle child ")
    if tree.MiddleChild != nil {
        fmt.Println(tree.MiddleChild.Value)
        DisplayParenthesesTree(tree.MiddleChild)
    } else {
        fmt.Println("no middle child")
    }
    fmt.Print("Right child ")
    if tree.RightChild != nil {
        fmt.Println(tree.RightChild.Value)
        DisplayParenthesesTree(tree.RightChild)
    } else {
        fmt.Println("no right child")
    }
}
