package day12b

import (
	"fmt"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

const generations = 50000000000
const edgePotsSize = 5
const diffCheckCount = 1000

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day12b", "Day 12 - Second Part", run)
}

var rules []*rule

type rule struct {
	rule   string
	result byte
}

type pot struct {
	id        int
	status    byte
	newStatus byte
}

func parseRule(input string) *rule {
	r := rule{rule: input[0:5], result: input[9]}
	return &r
}

func parsePots(input string) []*pot {
	var arr []*pot
	for i := 0; i < len(input); i++ {
		arr = append(arr, &pot{id: i, status: input[i]})
	}
	return arr
}

func applyRules(pots []*pot, index int) {
	var subStr string
	subArray := pots[index-2 : index+3]
	for _, p := range subArray {
		subStr = subStr + string(p.status)
	}

	for _, r := range rules {
		if string(subStr) == r.rule {
			pots[index].newStatus = r.result
			return
		}
	}
	pots[index].newStatus = '.'
}

func changeStatus(pots []*pot) {
	for i := 0; i < len(pots); i++ {
		pots[i].status = pots[i].newStatus
	}
}

func appendEdgePots(pots []*pot, start bool) []*pot {
	firstIndex := pots[0].id
	lastIndex := pots[len(pots)-1].id
	for i := 1; i <= edgePotsSize; i++ {
		if start {
			lp := pot{id: firstIndex - i, status: '.', newStatus: '.'}
			pots = append([]*pot{&lp}, pots...)
		} else {
			rp := pot{id: lastIndex + i, status: '.', newStatus: '.'}
			pots = append(pots, &rp)
		}
	}
	return pots
}

func doesEdgeChange(pots []*pot, index int) bool {
	var subStr string
	subArray := pots[index-2 : index+3]
	for _, p := range subArray {
		subStr = subStr + string(p.status)
	}
	for _, r := range rules {
		if string(subStr) == r.rule {
			return pots[index].status != r.result
		}
	}
	return false
}

func countPlants(pots []*pot) int {
	result := 0
	for _, p := range pots {
		if p.status == '#' {
			result += p.id
		}
	}
	return result
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	pots := parsePots(strings.Split(input[0], "initial state: ")[1])

	for i := 2; i < len(input); i++ {
		rules = append(rules, parseRule(input[i]))
	}

	// Add empty pots to both sides
	pots = appendEdgePots(pots, true)
	pots = appendEdgePots(pots, false)
	sameDiffCount := 0
	lastDiff := 0
	lastResult := 0
	for gen := 1; gen <= generations; gen++ {
		for i := 2; i < len(pots)-2; i++ {
			applyRules(pots, i)
		}
		changeStatus(pots)
		// Check if there is a change in first/last 5 pots according to the rules, if yes, add new empty pots on respective side
		if doesEdgeChange(pots, 2) {
			pots = appendEdgePots(pots, true)
		}
		if doesEdgeChange(pots, len(pots)-3) {
			pots = appendEdgePots(pots, false)
		}
		currentResult := countPlants(pots)
		if currentResult-lastResult == lastDiff {
			sameDiffCount++
		}
		if sameDiffCount == diffCheckCount {
			// No reason to continue as the diff is constant now, so compute the final result and end
			fmt.Printf("Result: %d\n", currentResult+(generations-gen)*lastDiff)
			break
		}

		lastDiff = currentResult - lastResult
		lastResult = currentResult
	}
}
