package main

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

type claim struct {
	id, posLeft, posTop, sizeX, sizeY int64
}

func processClaim(fabric [][]int, input string) {
	c := &claim{}
	_, e := fmt.Sscanf(input, "#%d @ %d,%d: %dx%d", &c.id, &c.posLeft, &c.posTop, &c.sizeX, &c.sizeY)

	if e != nil {
		panic(e)
	}

	for i := c.posTop; i < c.posTop+c.sizeY; i++ {
		for j := c.posLeft; j < c.posLeft+c.sizeX; j++ {
			fabric[i][j] = fabric[i][j] + 1
		}
	}
}

func main() {
	stringArray := strings.Split(util.GetInputString(), "\n")

	fabric := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		fabric[i] = make([]int, 1000)
	}

	for i := range stringArray {
		processClaim(fabric, stringArray[i])
	}

	overlaps := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				overlaps++
			}
		}
	}

	println("Overlaps: ", overlaps)
}
