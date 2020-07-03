package main

import (
	"fmt"
	"os"

	"github.com/marcbllv/regex3000/internal/regexparser"
)

func getArgsValue(args []string) (string, string) {
	if len(args) != 2 {
		fmt.Println("Exactly 2 arguments are needed: regex and string to test.")
		fmt.Printf("Recieved %d args: ", len(args))
		fmt.Println(args)
		os.Exit(1)
	}
	return args[0], args[1]
}

func main() {
	args := os.Args[1:]
	regex, stringToMatch := getArgsValue(args)

	if regexparser.CheckRegexMatch(regex, stringToMatch) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
