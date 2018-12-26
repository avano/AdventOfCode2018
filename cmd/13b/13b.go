package day13b

import (
	"fmt"
	"sort"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

const up = 0
const right = 1
const down = 2
const left = 3

const carRunes = "><^v"

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day13b", "Day 13 - Second Part", run)
}

type car struct {
	x, y      int
	direction int
	// -1 left, 0 straight, 1 right
	intersection int
	currentTile  rune
}

var slash map[int]int
var backSlash map[int]int
var direction map[int]rune

var m [][]rune
var cars []*car

func getCarAtCoordinates(x, y, index int) int {
	for i := 0; i < len(cars); i++ {
		if cars[i] != nil && index != i && cars[i].x == x && cars[i].y == y {
			return i
		}
	}
	return -1
}

func sortCars() {
	sort.Slice(cars, func(i, j int) bool {
		if cars[i].y > cars[j].y {
			return false
		} else if cars[i].y < cars[j].y {
			return true
		} else {
			return cars[i].x < cars[j].x
		}
	})
}

func moveCars() {
	for i := 0; i < len(cars); i++ {
		if cars[i] == nil {
			continue
		}
		m[cars[i].y][cars[i].x] = cars[i].currentTile
		switch cars[i].direction {
		case up:
			{
				cars[i].y--
			}
		case down:
			{
				cars[i].y++
			}
		case left:
			{
				cars[i].x--
			}
		default:
			{
				cars[i].x++
			}
		}

		if strings.ContainsAny(string(m[cars[i].y][cars[i].x]), carRunes) {
			otherCarIndex := getCarAtCoordinates(cars[i].x, cars[i].y, i)
			m[cars[otherCarIndex].y][cars[otherCarIndex].x] = cars[otherCarIndex].currentTile
			cars[otherCarIndex] = nil
			cars[i] = nil
		} else {
			if m[cars[i].y][cars[i].x] == '/' {
				cars[i].direction = slash[cars[i].direction]
			} else if m[cars[i].y][cars[i].x] == '\\' {
				cars[i].direction = backSlash[cars[i].direction]
			} else if m[cars[i].y][cars[i].x] == '+' {
				dir := (cars[i].direction + cars[i].intersection) % 4
				if dir < 0 {
					dir = dir + 4
				}
				cars[i].direction = dir
				cars[i].intersection++
				if cars[i].intersection == 2 {
					cars[i].intersection = -1
				}
			}
			cars[i].currentTile = m[cars[i].y][cars[i].x]
			m[cars[i].y][cars[i].x] = direction[cars[i].direction]
		}
	}
}

func removeCrashedCars() {
	var array []*car
	for i := 0; i < len(cars); i++ {
		if cars[i] != nil {
			array = append(array, cars[i])
		}
	}
	cars = array
}

func getDirectionInt(r rune) (bool, int) {
	for k, v := range direction {
		if v == r {
			return true, k
		}
	}
	return false, -1
}

func loadMapWithCars(input []string) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if strings.ContainsAny(string(input[y][x]), carRunes) {
				ok, d := getDirectionInt(rune(input[y][x]))
				if !ok {
					panic("Value doesn't exist in direction map")
				}
				cars = append(cars, &car{x: x, y: y, direction: d, intersection: -1, currentTile: 'X'})
			}
			m[y][x] = rune(input[y][x])
		}
	}
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	m = make([][]rune, len(input))
	for y := 0; y < len(input); y++ {
		m[y] = make([]rune, len(input[0]))
	}

	slash = map[int]int{
		up:    right,
		right: up,
		down:  left,
		left:  down,
	}
	backSlash = map[int]int{
		up:    left,
		right: down,
		down:  right,
		left:  up,
	}
	direction = map[int]rune{
		up:    '^',
		right: '>',
		down:  'v',
		left:  '<',
	}

	loadMapWithCars(input)

	for {
		sortCars()
		moveCars()
		removeCrashedCars()
		if len(cars) == 1 {
			fmt.Printf("Last car at [%d,%d]\n", cars[0].x, cars[0].y)
			break
		}
	}
}
