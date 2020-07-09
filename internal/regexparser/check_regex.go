package regexparser

func CheckRegexMatch(regex string, str string) bool {
	stateMachine := BuildStateMachine(regex)
	return stateMachine.Match([]rune(str), 0)
}
