package main

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
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

	fmt.Println(charCount)

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

	fmt.Println(stats)
}

func main() {
	stringArray := strings.Split(util.GetInputString(), "\n")
	stats := statistics{0, 0}

	for x := range stringArray {
		checkWord(stringArray[x], &stats)
	}

	println("Checksum: ", stats.twoChars*stats.threeChars)
}
