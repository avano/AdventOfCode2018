package day06a

import (
	"fmt"
	"math"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day06a", "Day 6 - First Part", run)
}

type point struct {
	id, x, y int
}

var points []point

func distance(p1, p2 point) int {
	return int((math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y))))
}

func getSearchCoordinates() (point, point) {
	topLeft := point{x: math.MaxInt64, y: math.MaxInt64}
	bottomRight := point{}
	for _, p := range points {
		if p.x > bottomRight.x {
			bottomRight.x = p.x
		}
		if p.x < topLeft.x {
			topLeft.x = p.x
		}
		if p.y > bottomRight.y {
			bottomRight.y = p.y
		}
		if p.y < topLeft.y {
			topLeft.y = p.y
		}
	}
	return topLeft, bottomRight
}

func getClosestDistancePointID(plotPoint point) int {
	closestDistance := math.MaxInt64
	closestDistanceID := -1
	for _, p := range points {
		currentDistance := distance(plotPoint, p)
		if closestDistance == currentDistance {
			closestDistanceID = -1
		} else if closestDistance > currentDistance {
			closestDistance = currentDistance
			closestDistanceID = p.id
		}
	}

	return closestDistanceID
}

func calculateDistances(plot [][]int, topLeft point) {
	for i := topLeft.x; i < len(plot); i++ {
		for j := topLeft.y; j < len(plot[i]); j++ {
			plot[i][j] = getClosestDistancePointID(point{x: i, y: j})
		}
	}
}

func getLargestArea(plot [][]int, topLeft point) int {
	results := make(map[int]int)
	for i := topLeft.x; i < len(plot); i++ {
		for j := topLeft.y; j < len(plot[i]); j++ {
			results[plot[i][j]]++
		}
	}

	max := 0
	for _, v := range results {
		if v > max {
			max = v
		}
	}

	return max
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	for i := 0; i < len(input); i++ {
		p := point{id: i + 1}
		_, e := fmt.Sscanf(input[i], "%d, %d", &p.x, &p.y)
		if e != nil {
			panic(e)
		}
		points = append(points, p)
	}

	topLeft, bottomRight := getSearchCoordinates()
	plot := make([][]int, bottomRight.x+1)
	for i := 0; i < bottomRight.x+1; i++ {
		plot[i] = make([]int, bottomRight.y+1)
	}

	calculateDistances(plot, topLeft)

	fmt.Printf("Largest area is %d\n", getLargestArea(plot, topLeft))
}
