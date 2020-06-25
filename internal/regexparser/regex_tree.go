package regexparser

import (
	"fmt"
	"strings"
)

type RegexTree struct {
	root Node
}


func isSpecialChar(char uint8) bool {
	return strings.ContainsRune("|()", rune(char))
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
	fmt.Printf("Tree %s, type '%s'\n", tree.Value, tree.Type)
	if len(tree.Children) == 0 {
		return
	}

	for i, child := range tree.Children {
		fmt.Printf("- Child %d/%d: ", i + 1, len(tree.Children))
		DisplayRegexTree(*child)
	}
	fmt.Println("-----")
}