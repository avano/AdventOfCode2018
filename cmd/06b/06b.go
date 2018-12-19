package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

const MaxDistance = 10000

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

func isWithinDistance(plotPoint point, points []point, maxDistance int) bool {
	totalDistance := 0
	for _, p := range points {
		totalDistance += distance(plotPoint, p)
	}
	return totalDistance < maxDistance
}

func calculateDistances(plot [][]bool, topLeft point, points []point) {
	for i := topLeft.x; i < len(plot); i++ {
		for j := topLeft.y; j < len(plot[i]); j++ {
			plot[i][j] = isWithinDistance(point{x: i, y: j}, points, MaxDistance)
		}
	}
}

func getDistanceAreaSize(plot [][]bool, topLeft point) int {
	size := 0
	for i := topLeft.x; i < len(plot); i++ {
		for j := topLeft.y; j < len(plot[i]); j++ {
			if plot[i][j] {
				size++
			}
		}
	}

	return size
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

	plot := make([][]bool, bottomRight.x+1)
	for i := 0; i < bottomRight.x+1; i++ {
		plot[i] = make([]bool, bottomRight.y+1)
	}

	calculateDistances(plot, topLeft, points)
	fmt.Printf("There are %d points within given area\n", getDistanceAreaSize(plot, topLeft))
}
