package day10

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

const displayableSize = 80

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day10", "Day 10", run)
}

type point struct {
	x, y, xVelocity, yVelocity int
}

func getArea(pointArray []*point) (point, point) {
	topLeft, bottomRight := point{x: pointArray[0].x, y: pointArray[0].y}, point{x: pointArray[0].x, y: pointArray[0].y}
	for i := 1; i < len(pointArray); i++ {
		if pointArray[i].x < topLeft.x {
			topLeft.x = pointArray[i].x
		}
		if pointArray[i].x > bottomRight.x {
			bottomRight.x = pointArray[i].x
		}
		if pointArray[i].y < topLeft.y {
			topLeft.y = pointArray[i].y
		}
		if pointArray[i].y > bottomRight.y {
			bottomRight.y = pointArray[i].y
		}
	}
	return topLeft, bottomRight
}

func isDrawable(topLeft, bottomRight point) bool {
	return bottomRight.x-topLeft.x <= displayableSize && bottomRight.y-topLeft.y <= displayableSize
}

func parsePoint(line string) *point {
	p := point{}
	fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &p.x, &p.y, &p.xVelocity, &p.yVelocity)
	return &p
}

func move(pointArray []*point) {
	for i := 0; i < len(pointArray); i++ {
		pointArray[i].x += pointArray[i].xVelocity
		pointArray[i].y += pointArray[i].yVelocity
	}
}

func printArea(pointArray []*point, topLeft, bottomRight point) {
	shiftX := topLeft.x * -1
	shiftY := topLeft.y * -1

	areaX := int(math.Abs(float64(topLeft.x-bottomRight.x))) + 1
	areaY := int(math.Abs(float64(topLeft.y-bottomRight.y))) + 1
	area := make([][]string, areaY)
	for i := 0; i < areaY; i++ {
		area[i] = make([]string, areaX)
		for j := 0; j < areaX; j++ {
			area[i][j] = " "
		}
	}

	for i := 0; i < len(pointArray); i++ {
		area[pointArray[i].y+shiftY][pointArray[i].x+shiftX] = "*"
	}

	for i := 0; i < areaY; i++ {
		fmt.Println(area[i])
	}
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")
	var pointArray []*point

	for _, line := range input {
		pointArray = append(pointArray, parsePoint(line))
	}

	lastWidth := math.MaxInt64
	for i := 0; ; i++ {
		topLeft, bottomRight := getArea(pointArray)
		width := int(math.Abs(float64(topLeft.y - bottomRight.y)))
		if lastWidth < width {
			fmt.Printf("Elapsed time: %d seconds\n", i-1)
			break
		} else {
			lastWidth = width
		}

		if isDrawable(topLeft, bottomRight) {
			printArea(pointArray, topLeft, bottomRight)
			time.Sleep(500 * time.Millisecond)
		}
		move(pointArray)
	}
}
