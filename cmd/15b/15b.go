package day15b

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day15b", "Day 15 - Second Part", run)
}

type unit struct {
	unitType    rune
	cell        *cell
	attackPower int
	hp          int
	alive       bool
}

type cell struct {
	x, y     int
	symbol   rune
	visited  bool
	prev     *cell
	distance int
}

type nextMove struct {
	nextCell   *cell
	targetCell *cell
}

var m [][]*cell
var units []*unit
var elfAttackPower = 3

func loadMap(input []string) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			c := &cell{x: x, y: y, symbol: rune(input[y][x])}
			if input[y][x] == 'G' {
				units = append(units, &unit{cell: c, unitType: rune(input[y][x]), attackPower: 3, hp: 200, alive: true})
			} else if input[y][x] == 'E' {
				units = append(units, &unit{cell: c, unitType: rune(input[y][x]), attackPower: elfAttackPower, hp: 200, alive: true})
			}
			m[y][x] = c
		}
	}
}

func readingOrder(c1 *cell, c2 *cell) bool {
	if c1 == nil && c2 != nil {
		return false
	}
	if c2 == nil && c1 != nil {
		return true
	}
	if c1.y < c2.y {
		return true
	} else if c2.y < c1.y {
		return false
	} else {
		return c1.x < c2.x
	}
}

func manhattan(c1 *cell, c2 *cell) int {
	return (int(math.Abs(float64(c1.x-c2.x)) + math.Abs(float64(c1.y-c2.y))))
}

func contains(cells []*cell, c *cell) bool {
	for i := 0; i < len(cells); i++ {
		if cells[i] == c {
			return true
		}
	}
	return false
}

func (c *cell) isEmpty() bool {
	return c.symbol == '.'
}

func (u *unit) enemyExists() bool {
	for i := 0; i < len(units); i++ {
		if u.unitType != units[i].unitType && units[i].alive {
			return true
		}
	}
	return false
}

func bfs(c1, c2 *cell) (*nextMove, int) {
	reachableCells := c2.getEmptyNeightborCells()
	shortestPath := math.MaxInt64
	var nm *nextMove

	sort.Slice(reachableCells, func(i, j int) bool {
		return readingOrder(reachableCells[i], reachableCells[j])
	})

	for i := 0; i < len(reachableCells); i++ {
		currentLength := 0

		// reset to default values
		for y := 0; y < len(m); y++ {
			for x := 0; x < len(m[y]); x++ {
				m[y][x].visited = false
				m[y][x].distance = 0
				m[y][x].prev = nil
			}
		}

		toVisit := []*cell{c1}
		for len(toVisit) > 0 {
			current := toVisit[0]
			current.visited = true

			if current == reachableCells[i] {
				currentLength = current.distance
				break
			}

			currentCells := current.getEmptyNeightborCells()
			for i := 0; i < len(currentCells); i++ {
				if currentCells[i].visited == true {
					continue
				}

				// Add next step to the list to visit if it is not already there
				if !contains(toVisit, currentCells[i]) && currentCells[i].distance <= shortestPath {
					currentCells[i].prev = current
					currentCells[i].distance = current.distance + 1
					toVisit = append(toVisit, currentCells[i])
				}
			}

			toVisit = toVisit[1:]
		}

		if reachableCells[i].prev == nil {
			// Didn't find any path to this cell
			continue
		}

		// Find the penultimate cell - this is the next move
		prev := reachableCells[i]
		for {
			if prev.prev == c1 {
				break
			}
			prev = prev.prev
		}

		if currentLength < shortestPath {
			shortestPath = currentLength
			nm = &nextMove{nextCell: prev, targetCell: reachableCells[i]}
		} else if currentLength == shortestPath {
			if !readingOrder(nm.nextCell, prev) {
				nm = &nextMove{nextCell: prev, targetCell: reachableCells[i]}
			}
		}
	}

	return nm, shortestPath
}

func (c *cell) getEmptyNeightborCells() []*cell {
	var cells []*cell

	if m[c.y-1][c.x].isEmpty() {
		cells = append(cells, m[c.y-1][c.x])
	}
	if m[c.y][c.x-1].isEmpty() {
		cells = append(cells, m[c.y][c.x-1])
	}
	if m[c.y][c.x+1].isEmpty() {
		cells = append(cells, m[c.y][c.x+1])
	}
	if m[c.y+1][c.x].isEmpty() {
		cells = append(cells, m[c.y+1][c.x])
	}

	return cells
}

func (u *unit) hasEmptySpace() bool {
	return len(u.cell.getEmptyNeightborCells()) > 0
}

func (u *unit) attack(targets []*unit) bool {
	var closestTarget *unit

	for i := 0; i < len(targets); i++ {
		if manhattan(u.cell, targets[i].cell) == 1 {
			if closestTarget == nil {
				closestTarget = targets[i]
			} else {
				if targets[i].hp == closestTarget.hp {
					if !readingOrder(closestTarget.cell, targets[i].cell) {
						closestTarget = targets[i]
					}
				} else if targets[i].hp < closestTarget.hp {
					closestTarget = targets[i]
				}
			}
		}
	}

	if closestTarget == nil {
		return true
	}

	closestTarget.hp = closestTarget.hp - u.attackPower

	if closestTarget.hp <= 0 {
		if closestTarget.unitType == 'E' {
			return false
		}
		closestTarget.alive = false
		m[closestTarget.cell.y][closestTarget.cell.x].symbol = '.'
	}
	return true
}

func (u *unit) moveToTarget(targets []*unit) {
	for i := 0; i < len(targets); i++ {
		if targets[i].alive && manhattan(u.cell, targets[i].cell) == 1 {
			// No need to move, already standing next to someone
			return
		}
	}

	var nm *nextMove
	nextMoves := make(map[int][]*nextMove)

	for i := 0; i < len(targets); i++ {
		// Get the "best" path to any empty cell around the target, save it as nextMove(nextCell, targetCell)
		nextMoveForTarget, dist := bfs(u.cell, targets[i].cell)
		if nextMoveForTarget != nil {
			// Some path was found, save it
			nextMoves[dist] = append(nextMoves[dist], nextMoveForTarget)
		}
	}

	if len(nextMoves) == 0 {
		// No paths to any target
		return
	}

	// Sort the next moves by the length of the path
	var keys []int
	for k, v := range nextMoves {
		if len(v) > 0 {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)

	sort.Slice(nextMoves[keys[0]], func(i, j int) bool {
		// Reading order on target cell, not the next cell
		return readingOrder(nextMoves[keys[0]][i].targetCell, nextMoves[keys[0]][j].targetCell)
	})
	nm = nextMoves[keys[0]][0]

	u.cell.symbol = '.'
	u.cell = nm.nextCell
	u.cell.symbol = u.unitType
}

func (u *unit) getTargets() []*unit {
	var targets []*unit
	for i := 0; i < len(units); i++ {
		if units[i] == u || units[i].unitType == u.unitType || !units[i].alive {
			continue
		}
		if (manhattan(u.cell, units[i].cell)) == 1 {
			targets = append(targets, units[i])
		} else {
			if units[i].unitType != u.unitType && units[i].hasEmptySpace() {
				targets = append(targets, units[i])
			}
		}
	}
	return targets
}

func removeDeadUnits() {
	var aliveUnits []*unit
	for i := 0; i < len(units); i++ {
		if units[i].alive {
			aliveUnits = append(aliveUnits, units[i])
		}
	}
	units = aliveUnits
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")
	turns := 0
simulation:
	for {
		elfAttackPower++
		units = []*unit{}
		m = make([][]*cell, len(input))
		for y := 0; y < len(input); y++ {
			m[y] = make([]*cell, len(input[0]))
		}

		loadMap(input)
		turns = 0

		for {
			sort.Slice(units, func(i, j int) bool {
				return readingOrder(units[i].cell, units[j].cell)
			})
			for i := 0; i < len(units); i++ {
				if !units[i].alive {
					continue
				}

				if !units[i].enemyExists() {
					break simulation
				}

				targets := units[i].getTargets()

				sort.Slice(targets, func(i, j int) bool {
					return readingOrder(targets[i].cell, targets[j].cell)
				})

				units[i].moveToTarget(targets)
				if !units[i].attack(targets) {
					// elf died
					continue simulation
				}

			}
			removeDeadUnits()
			turns++
		}
	}
	hp := 0
	for _, u := range units {
		if u.alive {
			hp += u.hp
		}
	}

	fmt.Printf("Elf attack power: %d\n", elfAttackPower)
	fmt.Printf("Turns: %d, hp: %d\n", turns, hp)
	fmt.Printf("Final score: %d\n", turns*hp)
}
