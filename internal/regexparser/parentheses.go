package regexparser

import (
	"strconv"
	"strings"
)

func findMatchingClosing(s []rune, openParPos int, open rune, closing rune) int {
	count := 0
	substr := s[openParPos+1:]

	for pos, char := range substr {
		if char == open {
			count++
		} else if char == closing {
			if count > 0 {
				count--
			} else {
				return openParPos + 1 + pos
			}
		}
	}
	// error
	// todo: Raise error properly
	return -1
}

func findMatchingParenthesis(s []rune, openParPos int) int {
	return findMatchingClosing(s, openParPos, '(', ')')
}

func findMatchingBrace(s []rune, openParPos int) int {
	return findMatchingClosing(s, openParPos, '{', '}')
}

func findMatchingBracket(s []rune, openParPos int) int {
	return findMatchingClosing(s, openParPos, '[', ']')
}

func parseBraceContent(content []rune) (int, int) {
	splitted := strings.Split(string(content), ",")
	var numbers []int

	for _, value := range splitted {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			// Todo: handle errors properly
			return -1, -1
		}
		numbers = append(numbers, intValue)
	}

	switch len(numbers) {
	case 1:
		return numbers[0], numbers[0]
	case 2:
		return numbers[0], numbers[1]
	default:
		// Todo: handle errors properly
		return -1, -1
	}
}
