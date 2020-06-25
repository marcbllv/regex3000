package regexparser

import (
	"fmt"
	"strconv"
	"strings"
)

type RegexTree struct {
	root Node
}


func buildSubRegexTree(regex string) *Node {
	rootNode_ := NewNode(regex, "")
	rootNode := &rootNode_
	pointer := 0
	var char uint8

	for pointer < len(regex) {
		char = regex[pointer]
		if char == '(' {
			newPosition, node := buildParenthesesSubTree(regex, pointer)
			rootNode.Children = append(rootNode.Children, node)
			pointer = newPosition + 1
		} else if char == '|' {
			disjunctionNode := NewNode(regex, "|")

			// Rotating root node to new disjuctionNode
			rootNode.Value = regex[:pointer]
			rightSideContent := regex[pointer + 1:]
			rightSideNode := buildSubRegexTree(rightSideContent)
			disjunctionNode.Children = append(disjunctionNode.Children, rootNode, rightSideNode)
			rootNode = &disjunctionNode
			pointer = len(regex)
		} else if char == '{' {
			nextBrace := findMatchingBrace(regex, pointer)
			innerContent := regex[pointer + 1:nextBrace]
			repeats := strings.Split(innerContent, ",")

			lastItem := rootNode.Children[len(rootNode.Children) - 1]
			lastItem.RepeatMin, _ = strconv.Atoi(repeats[0])
			if len(repeats) == 1 {
				lastItem.RepeatMax, _ = strconv.Atoi(repeats[0])
			} else if len(repeats) == 2 {
				lastItem.RepeatMax, _ = strconv.Atoi(repeats[1])
			}
			pointer = nextBrace + 1
		} else {
			charNode := NewNode(string(char), "")
			rootNode.Children = append(rootNode.Children, &charNode)
			pointer++
		}
	}
	return rootNode
}


func BuildRegexTree(regex string) Node {
	tree := buildSubRegexTree(regex)
	return *tree
}


func DisplayRegexTree(tree Node) {
	fmt.Printf("Tree %s, type '%s', repeat %d-%d\n", tree.Value, tree.Type, tree.RepeatMin, tree.RepeatMax)
	if len(tree.Children) == 0 {
		return
	}

	for i, child := range tree.Children {
		fmt.Printf("- Child %d/%d: ", i + 1, len(tree.Children))
		DisplayRegexTree(*child)
	}
	fmt.Println("-----")
}