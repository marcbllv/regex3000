package regexparser

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func findMatchingClosing(s string, openParPos int, open rune, closing rune) int {
	count := 0
	substr := s[openParPos:]
	_, runeSize := utf8.DecodeRuneInString(s)
	substr = substr[runeSize:]

	for pos, char := range substr {
		if char == open {
			count++
		} else if char == closing {
			if count > 0 {
				count--
			} else {
				return openParPos + runeSize + pos
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
