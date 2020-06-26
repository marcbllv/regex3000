package main

import (
    "fmt"
    "github.com/marcbllv/regexparser/internal/regexparser"
    "os"
)

func get_args_value(args []string) (string, string) {
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
    regex, stringToMatch := get_args_value(args)

    if regexparser.CheckRegexMatch(regex, stringToMatch) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
