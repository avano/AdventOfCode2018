package day02b

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day02b", "Day 2 - Second Part", run)
}

func checkWords(w1, w2 string) (bool, string) {
	mismatches := 0
	substring := ""
	for i := 0; i < len(w1); i++ {
		if w1[i] == w2[i] {
			substring += string(w1[i])
		} else {
			mismatches++
		}
		if mismatches > 1 {
			return false, ""
		}
	}
	return true, substring
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

outer:
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			matches, commonString := checkWords(input[i], input[j])
			if matches {
				fmt.Printf("Common substring: %s\n", commonString)
				break outer
			}
		}
	}
}
