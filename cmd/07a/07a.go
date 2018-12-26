package day07a

import (
	"fmt"
	"sort"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day07a", "Day 7 - First Part", run)
}

type step struct {
	name          string
	executed      bool
	prerequisites []*step
}

type sortByName []*step

func (a sortByName) Len() int           { return len(a) }
func (a sortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByName) Less(i, j int) bool { return a[i].name < a[j].name }

func getStep(steps []*step, name string) (bool, *step) {
	for i := 0; i < len(steps); i++ {
		if name == steps[i].name {
			return true, steps[i]
		}
	}
	st := step{name: name}
	return false, &st
}

func createStep(steps []*step, input string) []*step {
	var name string
	var prerequisite string
	_, err := fmt.Sscanf(input, "Step %s must be finished before step %s can begin.", &name, &prerequisite)
	if err != nil {
		panic(err)
	}

	foundDepend, dependStep := getStep(steps, prerequisite)
	foundCurrent, currentStep := getStep(steps, name)

	dependStep.prerequisites = append(dependStep.prerequisites, currentStep)

	if !foundDepend {
		steps = append(steps, dependStep)
	}

	if !foundCurrent {
		steps = append(steps, currentStep)
	}

	return append(steps)
}

func allExecuted(steps []*step) bool {
	for _, step := range steps {
		if !step.executed {
			return false
		}
	}
	return true
}

func (step step) canExecute() bool {
	return !step.executed && allExecuted(step.prerequisites)
}

func executeSteps(steps []*step) string {
	if allExecuted(steps) {
		return ""
	}

	var executed string
	for i := 0; i < len(steps); i++ {
		if steps[i].canExecute() {
			executed = steps[i].name
			steps[i].executed = true
			break
		}
	}

	return executed + executeSteps(steps)
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")

	var steps []*step

	for _, line := range input {
		steps = createStep(steps, line)
	}

	sort.Sort(sortByName(steps))
	fmt.Printf("Execution order: %s\n", executeSteps(steps))
}
