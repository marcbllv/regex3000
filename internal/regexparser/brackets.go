package regexparser

import (
	"strings"
	"unicode/utf8"
)

const (
	lettersChars        = "abcdefghijklmnopqrstuvwxyz"
	capitalLettersChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitsChars         = "0123456789"
)

func isDigit(r rune) bool {
	return strings.ContainsRune(digitsChars, r)
}

func isLetter(r rune) bool {
	return strings.ContainsRune(lettersChars, r)
}

func isCapitalLetter(r rune) bool {
	return strings.ContainsRune(capitalLettersChars, r)
}

func parseBracket(bracketContent string) ([]rune, bool) {
	var charSet []rune
	pos := 0
	bracketRunes := []rune(bracketContent)
	firstRune, _ := utf8.DecodeRuneInString(bracketContent)
	isOppositeSet := firstRune == '^'
	if isOppositeSet {
		pos++
	}

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
	return charSet, isOppositeSet
}
