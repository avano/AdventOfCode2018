package main

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

type claim struct {
	id, posLeft, posTop, sizeX, sizeY int
}

func processClaim(fabric [][]int, input string, check bool) bool {
	c := &claim{}
	_, e := fmt.Sscanf(input, "#%d @ %d,%d: %dx%d", &c.id, &c.posLeft, &c.posTop, &c.sizeX, &c.sizeY)

	if e != nil {
		panic(e)
	}

	overlap := false

	for i := c.posTop; i < c.posTop+c.sizeY; i++ {
		for j := c.posLeft; j < c.posLeft+c.sizeX; j++ {

		}
	}
outer:
	for i := c.posTop; i < c.posTop+c.sizeY; i++ {
		for j := c.posLeft; j < c.posLeft+c.sizeX; j++ {
			if check {
				if fabric[i][j] > 1 {
					overlap = true
					break outer
				}
			} else {
				fabric[i][j] = fabric[i][j] + 1
			}
		}
	}

	if check && !overlap {
		println("Not overlapping claim: #", c.id)
	}
	return !overlap
}

func main() {
	stringArray := strings.Split(util.GetInputString(), "\n")

	fabric := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		fabric[i] = make([]int, 1000)
	}

	for i := range stringArray {
		processClaim(fabric, stringArray[i], false)
	}

	for i := range stringArray {
		if processClaim(fabric, stringArray[i], true) {
			break
		}
	}
}
