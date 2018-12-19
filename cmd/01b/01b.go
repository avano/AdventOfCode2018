package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

func exists(seen map[int]bool, freq int) bool {
	if seen[freq] == true {
		return true
	}
	seen[freq] = true
	return false
}

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")

	frequency := 0
	seenFrequencies := make(map[int]bool)

outer:
	for {
		for _, line := range inputArray {
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

	fmt.Println("Duplicate frequency: ", frequency)
}
