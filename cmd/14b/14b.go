package day14b

import (
	"fmt"
	"strconv"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day14b", "Day 14 - Second Part", run)
}

var recipes []int

func run(cmd *cobra.Command, _ []string) {
	input := util.ReadInput(file, example)

	elf1 := 0
	elf2 := 1

	recipes = []int{3, 7}

	found := 0
outer:
	for {
		score := recipes[elf1] + recipes[elf2]
		scoreString := strconv.Itoa(score)
		for i := 0; i < len(scoreString); i++ {
			if scoreString[i] == input[found] {
				found++
			} else {
				found = 0
				// Re-check if this can be a starting num again
				if scoreString[i] == input[found] {
					found++
				}
			}
			num, _ := strconv.Atoi(string(scoreString[i]))
			recipes = append(recipes, num)
			if found == len(input) {
				break outer
			}
		}

		elf1 = (elf1 + 1 + recipes[elf1]) % len(recipes)
		elf2 = (elf2 + 1 + recipes[elf2]) % len(recipes)
	}

	fmt.Printf("Score %s found after crafting %d recipes\n", input, len(recipes)-len(input))
}
