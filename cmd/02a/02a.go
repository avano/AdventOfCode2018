package day02a

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day02a", "Day 2 - First Part", run)
}

type statistics struct {
	twoChars   int
	threeChars int
}

var stats *statistics

func checkWord(word string) {
	charCount := make(map[string]int)
	for _, r := range word {
		c := string(r)
		charCount[c] = charCount[c] + 1
	}

	foundTwo := false
	foundThree := false

	for _, v := range charCount {
		if !foundTwo && v == 2 {
			stats.twoChars++
			foundTwo = true
		}
		if !foundThree && v == 3 {
			stats.threeChars++
			foundThree = true
		}
		if foundTwo && foundThree {
			break
		}
	}
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")
	stats = &statistics{0, 0}

	for _, line := range input {
		checkWord(line)
	}

	fmt.Printf("Checksum: %d\n", stats.twoChars*stats.threeChars)
}
