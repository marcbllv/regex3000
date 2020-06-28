package regexparser

import (
    "strconv"
    "strings"
)


const(
    alphaChars="abcdefghijklmnopqrstuvwxyz"
    numberChars="0123456789"
)


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


func findMatchingParenthesis(s string, openParPos int) int {
    return findMatchingClosing(s, openParPos, '(', ')')
}


func findMatchingBrace(s string, openParPos int) int {
    return findMatchingClosing(s, openParPos, '{', '}')
}


func findMatchingBracket(s string, openParPos int) int {
    return findMatchingClosing(s, openParPos, '[', ']')
}


func parseBraceContent(content string) (int, int) {
    splitted := strings.Split(content, ",")
    var numbers []int

    for _, value := range splitted {
        intValue, err := strconv.Atoi(value)
        if err != nil {
            // Todo: handle errors properly
            return -1, -1
        }
        numbers = append(numbers, intValue)
    }

    if len(numbers) == 1 {
        return numbers[0], numbers[0]
    } else if len(numbers) == 2 {
        return numbers[0], numbers[1]
    } else {
        // Todo: handle errors properly
        return -1, -1
    }
}


func parseBracket(bracketContent string) []rune {
    var charSet []rune
    var lastChar rune
    pos := 0
    for pos < len(bracketContent) {
        char := rune(bracketContent[pos])
        if strings.ContainsRune(alphaChars, char) {
            
        }

        pos++
    }
    return charSet
}


