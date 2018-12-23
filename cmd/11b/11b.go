package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

const size = 300

type point struct {
	x, y, value int
	squares     map[int]int
}

func computeCell(x, y, serial int) *point {
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
	p := point{x: x, y: y, value: powerLevel}
	p.squares = make(map[int]int)
	return &p
}

func computeSquares(x, y int, grid [][]*point) {
	squareSize := int(math.Min(float64(x), float64(y)))
	// Always have size 1 square with the value of the point
	grid[x][y].squares[1] = grid[x][y].value

	// Compute the size of squares size 1..squareSize
	// Get the square of size i from point [x-i,y-i] until we get to the border
	// So for [3,3]: Get square size of 1 from [2,2] and compute remaining area
	//               Get square size of 2 from [1,1] and compute remaining area
	// And save the values in the respective points
	for k := 1; k < squareSize; k++ {
		value := grid[x-k][y-k].squares[k]
		for i := 1; i < k+1; i++ {
			value += grid[x-i][y].value
		}
		for i := 1; i < k+1; i++ {
			value += grid[x][y-i].value
		}
		value += grid[x][y].value

		grid[x-k][y-k].squares[k+1] = value
	}

}

func main() {
	gridSerialNo, err := strconv.Atoi(util.GetInputString())
	if err != nil {
		panic(err)
	}

	grid := make([][]*point, size+1)
	for i := 0; i < size+1; i++ {
		grid[i] = make([]*point, size+1)
	}

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			grid[j][i] = computeCell(j, i, gridSerialNo)
			computeSquares(j, i, grid)
		}
	}

	var maxPoint *point
	squareSize := 0
	max := math.MinInt64
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			for k, v := range grid[j][i].squares {
				if v > max {
					max = v
					squareSize = k
					maxPoint = grid[j][i]
				}
			}
		}
	}

	fmt.Printf("Largest total power at [%d,%d], size %d (power %d)\n", maxPoint.x, maxPoint.y, squareSize, max)
}
