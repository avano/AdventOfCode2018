package main

import (
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

type statistics struct {
	twoChars   int
	threeChars int
}

func checkWord(word string, stats *statistics) {
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

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")
	stats := statistics{0, 0}

	for _, line := range inputArray {
		checkWord(line, &stats)
	}

	println("Checksum: ", stats.twoChars*stats.threeChars)
}
