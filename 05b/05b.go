package main

import (
	"fmt"
	"math"
	"strings"
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

func replace(input string, char string) string {
	if strings.Contains(strings.ToLower(input), char) {
		input = strings.Replace(input, char, "", -1)
		input = strings.Replace(input, strings.ToUpper(char), "", -1)
	}
	return input
}

func main() {
	input := util.GetInputString()

	var bestLetter rune
	bestLength := math.MaxInt64
	for ch := 'a'; ch <= 'z'; ch++ {
		currentLength := len(react(replace(input, string(ch))))
		fmt.Printf("Length: %d (leaving out %c)\n", currentLength, ch)
		if currentLength < bestLength {
			bestLength = currentLength
			bestLetter = ch
		}
	}

	fmt.Printf("Best length: %d (leaving out %c)\n", bestLength, bestLetter)
}
