package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/avano/AdventOfCode2018/util"
)

type point struct {
	id, x, y int
}

func distance(p1, p2 point) int {
	return int((math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y))))
}

func getSearchCoordinates(points []point) (point, point) {
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

func getClosestDistancePointID(plotPoint point, points []point) int {
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

func calculateDistances(plot [][]int, topLeft point, points []point) {
	for i := topLeft.x; i < len(plot); i++ {
		for j := topLeft.y; j < len(plot[i]); j++ {
			plot[i][j] = getClosestDistancePointID(point{x: i, y: j}, points)
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

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")

	var points []point

	for i := 0; i < len(inputArray); i++ {
		p := point{id: i + 1}
		_, e := fmt.Sscanf(inputArray[i], "%d, %d", &p.x, &p.y)
		if e != nil {
			panic(e)
		}
		points = append(points, p)
	}

	topLeft, bottomRight := getSearchCoordinates(points)
	fmt.Printf("Search coordinates: [%d,%d] x [%d,%d]\n", topLeft.x, topLeft.y, bottomRight.x, bottomRight.y)

	plot := make([][]int, bottomRight.x+1)
	for i := 0; i < bottomRight.x+1; i++ {
		plot[i] = make([]int, bottomRight.y+1)
	}

	calculateDistances(plot, topLeft, points)

	fmt.Printf("Largest area is %d\n", getLargestArea(plot, topLeft))
}
