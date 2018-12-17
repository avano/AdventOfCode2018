package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

func main() {
	stringArray := strings.Split(util.GetInputString(), "\n")

	frequency := 0

	for i := 0; i < len(stringArray); i++ {
		num, err := strconv.ParseInt(stringArray[i], 0, 0)
		if err != nil {
			panic(err)
		}
		frequency += int(num)
	}

	fmt.Println("Frequency: ", frequency)
}
