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


func buildParenthesesSubTree(regex string, pointer int) (int, *Node) {
    closingParenthesis := findMatchingParenthese(regex, pointer)
    innerContent := regex[pointer + 1:closingParenthesis]
    node := buildSubRegexTree(innerContent)
    return closingParenthesis, node
}
