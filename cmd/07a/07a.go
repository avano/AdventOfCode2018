package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

type step struct {
	name          string
	executed      bool
	prerequisites []*step
}

type nameSorter []*step

func (a nameSorter) Len() int           { return len(a) }
func (a nameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a nameSorter) Less(i, j int) bool { return a[i].name < a[j].name }

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

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")

	var steps []*step

	for _, line := range inputArray {
		steps = createStep(steps, line)
	}

	sort.Sort(nameSorter(steps))
	fmt.Printf("Execution order: %s\n", executeSteps(steps))
}
