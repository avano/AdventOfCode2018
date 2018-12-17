package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

func exists(seen map[int]bool, freq int) bool {
	if seen[freq] == true {
		return true
	}
	seen[freq] = true
	return false
}

func main() {
	stringArray := strings.Split(util.GetInputString(), "\n")

	frequency := 0
	seenFrequencies := make(map[int]bool)

outer:
	for j := 0; ; j++ {
		for i := 0; i < len(stringArray); i++ {
			num, err := strconv.ParseInt(stringArray[i], 0, 0)
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
