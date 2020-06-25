package regexparser

import (
	"strings"
)

func CheckRegexMatch(regex string, str string) bool {
	regexTree := BuildRegexTree(regex)
	remainingStr := checkRegexTreeMatch(regexTree, str)
	for _, str := range remainingStr {
		if str == "" {
			return true
		}
	}
	return false
}

func checkRegexTreeMatch(regexTree Node, str string) []string {
	if regexTree.Type == "|" {
		// Disjunction
		var allRemainingStr []string
		for _, childTree := range regexTree.Children {
			remainingStr := checkRegexTreeMatch(*childTree, str)
			allRemainingStr = append(allRemainingStr, remainingStr...)
		}
		return allRemainingStr
	} else if len(regexTree.Children) > 0 {
		// Concatenation of all children
		substrs := []string{str}
		for _, child := range regexTree.Children {
			var newRemainingStrings []string
			for _, substr := range substrs {
				remainingStrings := checkRegexTreeMatch(*child, substr)
				newRemainingStrings = append(newRemainingStrings, remainingStrings...)
			}
			substrs = newRemainingStrings
		}
		return substrs
	} else {
		// Simple string equality, no special chars
		var remainingStrings []string
		if strings.HasPrefix(str, regexTree.Value) {
			remainingStr := strings.TrimPrefix(str, regexTree.Value)
			remainingStrings = append(remainingStrings, remainingStr)
		}
		return remainingStrings
	}
}
