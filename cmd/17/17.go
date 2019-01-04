package day17

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day17", "Day 17", run)
}

var m [][]rune
var leftX, leftY, rightX, rightY int

func loadCoordinates(input string) (int, int, int, int) {
	var xMin, xMax, yMin, yMax int
	_, err := fmt.Sscanf(input, "x=%d, y=%d..%d", &xMin, &yMin, &yMax)
	xMax = xMin
	if err != nil {
		_, err2 := fmt.Sscanf(input, "y=%d, x=%d..%d", &yMin, &xMin, &xMax)
		yMax = yMin
		if err2 != nil {
			panic(err)
		}
	}

	return xMin, xMax, yMin, yMax
}

func getArea(input []string) {
	leftX, leftY, rightX, rightY = 500, 0, 500, 0

	for i := 0; i < len(input); i++ {
		xMin, xMax, yMin, yMax := loadCoordinates(string(input[i]))
		if xMin < leftX {
			leftX = xMin
		}
		if xMax > rightX {
			rightX = xMax
		}
		if yMin < leftY {
			leftY = yMin
		}
		if yMax > rightY {
			rightY = yMax
		}
	}
}

func getCountYCoordinate() (int, int) {
	minY, maxY := -1, -1
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == '#' {
				if minY == -1 {
					minY = y
				}
				maxY = y
			}
		}
	}
	return minY, maxY
}

func loadMap(input []string) {
	for i := 0; i < len(input); i++ {
		xMin, xMax, yMin, yMax := loadCoordinates(string(input[i]))
		for x := xMin - leftX; x <= xMax-leftX; x++ {
			for y := yMin; y <= yMax; y++ {
				m[y][x] = '#'
			}
		}
	}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == 0 {
				m[y][x] = '.'
			}
		}
	}
}

func canFlow(x, y int) bool {
	return m[y][x] == '.' || m[y][x] == '|'
}

func flow(x, y int) {
	if (y + 1) == len(m) {
		return
	}
	if m[y][x] == '#' {
		return
	}

	if !canFlow(x, y+1) {
		left := x
		for canFlow(left, y) && !canFlow(left, y+1) {
			m[y][left] = '|'
			left--
		}
		right := x + 1
		for canFlow(right, y) && !canFlow(right, y+1) {
			m[y][right] = '|'
			right++
		}
		if canFlow(left, y+1) || canFlow(right, y+1) {
			flow(left, y)
			flow(right, y)
		} else if m[y][left] == '#' && m[y][right] == '#' {
			for currX := left + 1; currX < right; currX++ {
				m[y][currX] = '~'
			}
		}
	} else if m[y][x] == '.' {
		m[y][x] = '|'
		flow(x, y+1)
		if m[y+1][x] == '~' {
			flow(x, y)
		}
	}
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	getArea(input)
	leftX = leftX - 1

	m = make([][]rune, rightY-leftY+2)
	for y := 0; y < len(m); y++ {
		m[y] = make([]rune, rightX-leftX+2)
	}

	loadMap(input)

	flow(500-leftX, 0)

	countRest := 0
	countFlowing := 0
	minY, maxY := getCountYCoordinate()
	for y := minY; y <= maxY; y++ {
		for x := 0; x < len(m[y]); x++ {
			if m[y][x] == '~' {
				countRest++
			} else if m[y][x] == '|' {
				countFlowing++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", countRest+countFlowing)
	fmt.Printf("Part 2: %d\n", countRest)
}
