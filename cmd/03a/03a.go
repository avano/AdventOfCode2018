package day03a

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day03a", "Day 3 - First Part", run)
}

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

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	var claims []claim

	for _, line := range input {
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

	fmt.Printf("Overlaps: %d\n", overlaps)
}
