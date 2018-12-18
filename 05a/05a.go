package main

import (
	"fmt"
	"unicode"

	"github.com/avano/AdventOfCode2018/util"
)

func react(input string) string {
	reacted := false
	var i int
	for i = 0; i < len(input)-1; i++ {
		if unicode.ToLower(rune(input[i])) == unicode.ToLower(rune(input[i+1])) &&
			((unicode.IsLower(rune(input[i])) && unicode.IsUpper(rune(input[i+1]))) || (unicode.IsUpper(rune(input[i])) && unicode.IsLower(rune(input[i+1])))) {
			reacted = true
			break
		}
	}

	if reacted {
		input = react(input[0:i] + input[i+2:len(input)])
	}
	return input
}

func main() {
	input := util.GetInputString()
	fmt.Println("Final length: ", len(react(input)))
}
