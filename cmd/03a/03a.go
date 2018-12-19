package main

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

type claim struct {
	id, posLeft, posTop, sizeX, sizeY int
}

func parseClaim(input string) claim {
	c := claim{}
	_, e := fmt.Sscanf(input, "#%d @ %d,%d: %dx%d", &c.id, &c.posLeft, &c.posTop, &c.sizeX, &c.sizeY)

	if e != nil {
		panic(e)
	}

	return c
}

func getDimensions(claims []claim) (int, int) {
	maxX, maxY := 0, 0

	for _, c := range claims {
		if c.posLeft+c.sizeX > maxX {
			maxX = c.posLeft + c.sizeX
		}
		if c.posTop+c.sizeY > maxY {
			maxY = c.posTop + c.sizeY
		}
	}

	return maxX, maxY
}

func processClaim(fabric [][]int, c claim) {
	for i := c.posTop; i < c.posTop+c.sizeY; i++ {
		for j := c.posLeft; j < c.posLeft+c.sizeX; j++ {
			fabric[i][j] = fabric[i][j] + 1
		}
	}
}

func main() {
	stringArray := strings.Split(util.GetInputString(), "\n")

	var claims []claim

	for _, line := range stringArray {
		claims = append(claims, parseClaim(line))
	}

	x, y := getDimensions(claims)

	fabric := make([][]int, x+1)
	for i := 0; i < x+1; i++ {
		fabric[i] = make([]int, y+1)
	}

	for _, c := range claims {
		processClaim(fabric, c)
	}

	overlaps := 0
	for i := 0; i < x+1; i++ {
		for j := 0; j < y+1; j++ {
			if fabric[i][j] > 1 {
				overlaps++
			}
		}
	}

	println("Overlaps: ", overlaps)
}
