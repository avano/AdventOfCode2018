package main

import (
	"fmt"
	"strconv"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

const size = 300

type point struct {
	x, y int
}

func computeBackwards(grid [][]int, x, y int) (point, int) {
	p := point{x: x - 2, y: y - 2}
	value := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			value += grid[p.x+i][p.y+j]
		}
	}
	return p, value
}

func computeCell(x, y, serial int) int {
	rackID := x + 10
	powerLevel := y * rackID
	powerLevel += serial
	powerLevel *= rackID

	if powerLevel > 99 {
		powerLevelString := strconv.Itoa(powerLevel)
		length := len(powerLevelString)
		hundreds, err := strconv.Atoi(string(powerLevelString[length-3]))
		if err != nil {
			panic(err)
		}
		powerLevel = hundreds
	} else {
		powerLevel = 0
	}
	powerLevel -= 5
	return powerLevel
}

func main() {
	gridSerialNo, err := strconv.Atoi(util.GetInputString())
	if err != nil {
		panic(err)
	}

	grid := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		grid[i] = make([]int, size+1)
	}

	powerLevels := make(map[point]int)

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			grid[i][j] = computeCell(i, j, gridSerialNo)
			if i >= 3 && j >= 3 {
				p, value := computeBackwards(grid, i, j)
				powerLevels[p] = value
			}
		}
	}

	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			p := point{x: i, y: j}
			value := 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					value += grid[p.x+k][p.y+l]
				}
			}
			powerLevels[p] = value
		}
	}

	max := 0
	var maxPoint point
	for k, v := range powerLevels {
		if v > max {
			max = v
			maxPoint = k
		}
	}
	fmt.Printf("Largest total power at [%d,%d]\n", maxPoint.x, maxPoint.y)
}
