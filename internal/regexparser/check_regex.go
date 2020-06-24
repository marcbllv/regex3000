package regexparser

import (
	"strings"
)

func CheckRegexMatch(regex string, str string) bool {
	var parenthesesTree = BuildParenthesesTree(regex)
	var matchingPositions = checkParenthesesSubStrRegexMatch(parenthesesTree, str)

	for _, matchingPosition := range matchingPositions {
		if len(str) == matchingPosition {
			return true
		}
	}
	return false
}


func checkParenthesesSubStrRegexMatch(parenthesesTree *TernaryNode, str string) []int {
	if parenthesesTree == nil {
		return []int{0}
	}
	if isLeafTernaryNode(*parenthesesTree) {
		return matchSimpleRegex(parenthesesTree.Value, str)
	}

	matchingPositions := []int{0}
	var childTreeMatchinPositions []int
	// Loop over left, middle and right subtrees
	for _, childTree := range getChildren(*parenthesesTree) {
		childTreeMatchinPositions = []int{}
		for _, matchingPosition := range matchingPositions {
			subString := str[matchingPosition:]
			newMatchingPositions := checkParenthesesSubStrRegexMatch(childTree, subString)
			for _, newMatchingPosition := range newMatchingPositions {
				childTreeMatchinPositions = append(childTreeMatchinPositions, matchingPosition + newMatchingPosition)
			}
		}
		matchingPositions = childTreeMatchinPositions
	}
	return matchingPositions
}


func matchSimpleRegex(regex string, str string) []int {
	disjuctionTree := buildDisjunctionTree(regex)
	return matchDisjunctionFromTree(disjuctionTree, str)
}


func matchDisjunctionFromTree(disjunctionTree BinaryNode, str string) []int {
	var matchingPositions []int
	if isLeafBinaryNode(disjunctionTree) {
		matchingPosition := matchPrefixString(disjunctionTree.Value, str)
		if matchingPosition >= 0 {
			return []int{matchingPosition}
		} else {
			return []int{}
		}
	}

	if disjunctionTree.LeftChild != nil {
		leftMatchingPositions := matchDisjunctionFromTree(*disjunctionTree.LeftChild, str)
		matchingPositions = append(matchingPositions, leftMatchingPositions...)
	}
	if disjunctionTree.RightChild != nil {
		rightMatchingPositions := matchDisjunctionFromTree(*disjunctionTree.RightChild, str)
		matchingPositions = append(matchingPositions, rightMatchingPositions...)
	}
	return matchingPositions
}


func matchPrefixString(regex string, str string) int {
	if strings.HasPrefix(str, regex) {
		return len(regex)
	} else {
		return -1
	}
}
