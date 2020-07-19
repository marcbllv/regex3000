package regexparser

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

func findMatchingBracket(s []rune, openParPos int) int {
	return findMatchingClosing(s, openParPos, '[', ']')
}
