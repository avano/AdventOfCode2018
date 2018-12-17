package main

import (
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

func checkWords(w1 string, w2 string) (bool, string) {
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
	stringArray := strings.Split(util.GetInputString(), "\n")

outer:
	for i := 0; i < len(stringArray)-1; i++ {
		for j := i + 1; j < len(stringArray); j++ {
			matches, commonString := checkWords(stringArray[i], stringArray[j])
			if matches {
				println("Common substring: ", commonString)
				break outer
			}
		}
	}
}
