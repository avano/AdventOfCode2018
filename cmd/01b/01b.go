package day01b

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day01b", "Day 1 - Second Part", run)
}

func exists(seen map[int]bool, freq int) bool {
	if _, ex := seen[freq]; ex {
		return true
	}
	seen[freq] = true
	return false
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	frequency := 0
	seenFrequencies := make(map[int]bool)

outer:
	for {
		for _, line := range input {
			num, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			frequency += int(num)

			if exists(seenFrequencies, frequency) {
				break outer
			}
		}
	}

	fmt.Printf("Duplicate frequency: %d\n", frequency)
}
