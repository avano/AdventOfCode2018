package main

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

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

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")

outer:
	for i := 0; i < len(inputArray)-1; i++ {
		for j := i + 1; j < len(inputArray); j++ {
			matches, commonString := checkWords(inputArray[i], inputArray[j])
			if matches {
				fmt.Println("Common substring: ", commonString)
				break outer
			}
		}
	}
}
