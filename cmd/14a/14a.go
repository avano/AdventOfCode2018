package main

import (
	"fmt"
	"strconv"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

var recipes []int

func main() {
	input, err := strconv.Atoi(util.GetInputString())

	if err != nil {
		panic(err)
	}

	elf1 := 0
	elf2 := 1

	recipes = []int{3, 7}

	for len(recipes) < input+10 {
		score := recipes[elf1] + recipes[elf2]
		scoreString := strconv.Itoa(score)
		for i := 0; i < len(scoreString); i++ {
			num, _ := strconv.Atoi(string(scoreString[i]))
			recipes = append(recipes, num)
		}

		elf1 = (elf1 + 1 + recipes[elf1]) % len(recipes)
		elf2 = (elf2 + 1 + recipes[elf2]) % len(recipes)
	}

	score := ""
	for i := input; i < input+10; i++ {
		score = score + strconv.Itoa(recipes[i])
	}

	fmt.Printf("Score after %d recipes: %s\n", input, score)
}
