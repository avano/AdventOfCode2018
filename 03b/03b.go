package main

import (
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

type claim struct {
	id, posLeft, posTop, sizeX, sizeY int64
}

func processClaim(fabric [][]int, input string, check bool) bool {
	split := strings.Split(input, " ")
	pos := strings.Split(split[2], ",")
	size := strings.Split(split[3], "x")

	id, _ := strconv.ParseInt(split[0][1:len(split[0])], 0, 0)
	posLeft, _ := strconv.ParseInt(pos[0], 0, 0)
	posTop, _ := strconv.ParseInt(pos[1][0:len(pos[1])-1], 0, 0)
	sizeX, _ := strconv.ParseInt(size[0], 0, 0)
	sizeY, _ := strconv.ParseInt(size[1], 0, 0)

	c := claim{id: id, posLeft: posLeft, posTop: posTop, sizeX: sizeX, sizeY: sizeY}

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
