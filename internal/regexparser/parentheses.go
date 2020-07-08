package regexparser

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	lettersChars        = "abcdefghijklmnopqrstuvwxyz"
	capitalLettersChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitsChars         = "0123456789"
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

func isDigit(r rune) bool {
	return strings.ContainsRune(digitsChars, r)
}

func isLetter(r rune) bool {
	return strings.ContainsRune(lettersChars, r)
}

func isCapitalLetter(r rune) bool {
	return strings.ContainsRune(capitalLettersChars, r)
}

func parseBracket(bracketContent string) []rune {
	var charSet []rune
	pos := 0
	bracketRunes := []rune(bracketContent)
	for pos < len(bracketRunes) {
		if pos < len(bracketRunes)-2 && bracketRunes[pos+1] == '-' {
			threeCharsPattern := bracketRunes[pos : pos+3]
			pos += 3

			letter1 := threeCharsPattern[0]
			letter2 := threeCharsPattern[2]
			var baseString string

			switch {
			case isDigit(letter1) && isDigit(letter2):
				baseString = digitsChars
			case isLetter(letter1) && isLetter(letter2):
				baseString = lettersChars
			case isCapitalLetter(letter1) && isCapitalLetter(letter2):
				baseString = capitalLettersChars
			default:
				// TODO: raise error here
				continue
			}
			letter1Idx := strings.IndexRune(baseString, letter1)
			letter2Idx := strings.IndexRune(baseString, letter2)
			for _, digit := range baseString[letter1Idx : letter2Idx+1] {
				charSet = append(charSet, digit)
			}
		} else {
			charSet = append(charSet, bracketRunes[pos])
			pos++
		}
	}
	return charSet
}
