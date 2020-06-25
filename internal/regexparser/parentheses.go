package regexparser

func findMatchingClosing(s string, openParPos int, open rune, closing rune) int {
    count := 0
    substr := s[openParPos+ 1:]

    for pos, char := range substr {
        if char == open {
            count++
        } else if char == closing {
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


func findMatchingParenthese(s string, openParPos int) int {
    return findMatchingClosing(s, openParPos, '(', ')')
}


func findMatchingBrace(s string, openParPos int) int {
    return findMatchingClosing(s, openParPos, '{', '}')
}


func buildParenthesesSubTree(regex string, pointer int) (int, *Node) {
    closingParenthesis := findMatchingParenthese(regex, pointer)
    innerContent := regex[pointer + 1:closingParenthesis]
    node := buildSubRegexTree(innerContent)
    return closingParenthesis, node
}
