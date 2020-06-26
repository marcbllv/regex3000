package regexparser

import (
	"fmt"
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
		for i := regexTree.RepeatMin; i <= regexTree.RepeatMax; i++ {
			for _, child := range regexTree.Children {
				var newRemainingStrings []string
				for _, substr := range substrs {
					remainingStrings := checkRegexTreeMatch(*child, substr)
					newRemainingStrings = append(newRemainingStrings, remainingStrings...)
				}
				substrs = newRemainingStrings
			}
		}
		return substrs
	} else {
		// Simple string equality, no special chars
		remainingString := str
		fmt.Println("Handling:", str, regexTree.Value)
		for i := 0; i < regexTree.RepeatMin; i++ {
			if !strings.HasPrefix(remainingString, regexTree.Value) {
				fmt.Println("No solution:", i, remainingString, regexTree.Value)
				return []string{} // No solution with this repeat count
			}
			remainingString = strings.TrimPrefix(remainingString, regexTree.Value)
		}

		fmt.Println("Remains:", remainingString)
		remainingStrings := []string{remainingString}
		for i := regexTree.RepeatMin; i < regexTree.RepeatMax; i++ {
			fmt.Println("Remaining strings", remainingStrings)
			if strings.HasPrefix(remainingString, regexTree.Value) {
				remainingStrings = append(remainingStrings, remainingString)
			}
			remainingString = strings.TrimPrefix(remainingString, regexTree.Value)
		}
		fmt.Println("Returning", remainingStrings)
		fmt.Println("----")
		return remainingStrings
	}
}
