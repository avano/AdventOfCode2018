package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")

	frequency := 0

	for _, line := range inputArray {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		frequency += int(num)
	}

	fmt.Println("Frequency: ", frequency)
}
