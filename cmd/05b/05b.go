package day05b

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day05a", "Day 5 - Second Part", run)
}

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

func run(cmd *cobra.Command, _ []string) {
	input := util.ReadInput(file, example)

	var bestLetter rune
	bestLength := math.MaxInt64
	for ch := 'a'; ch <= 'z'; ch++ {
		currentLength := len(react(replace(input, string(ch))))
		if currentLength < bestLength {
			bestLength = currentLength
			bestLetter = ch
		}
	}

	fmt.Printf("Best length: %d (leaving out %c)\n", bestLength, bestLetter)
}
