package day01a

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
	file, example = util.RegisterCommand("day01a", "Day 1 - First Part", run)
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	frequency := 0

	for _, line := range input {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		frequency += int(num)
	}

	fmt.Printf("Frequency: %d\n", frequency)
}
