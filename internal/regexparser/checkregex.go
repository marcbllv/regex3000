package regexparser

import (
	"strings"
)

func CheckRegexMatch(regex string, str string) bool {
	var parenthesesTree = BuildParenthesesTree(regex)
	var matchingPosition = checkSubStrRegexMatch(parenthesesTree, str)
	return matchingPosition == len(str)
}


func checkSubStrRegexMatch(parenthesesTree *TernaryNode, str string) int {
	if IsLeafTernaryNode(parenthesesTree) {
		return matchSimpleRegex(parenthesesTree.Value, str)
	}
	stringMatchPos := 0
	if parenthesesTree.LeftChild != nil {
		stringMatchPos += checkSubStrRegexMatch(parenthesesTree.LeftChild, str[stringMatchPos:])
	}
	if parenthesesTree.MiddleChild != nil {
		stringMatchPos += checkSubStrRegexMatch(parenthesesTree.MiddleChild, str[stringMatchPos:])
	}
	if parenthesesTree.RightChild != nil {
		stringMatchPos += checkSubStrRegexMatch(parenthesesTree.RightChild, str[stringMatchPos:])
	}
	return stringMatchPos
}


func matchSimpleRegex(regex string, str string) int {
	if strings.HasPrefix(str, regex) {
		return len(regex)
	} else {
		return -1
	}
}
