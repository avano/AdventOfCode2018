package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

const workersCount = 5
const executionTime = 60

type step struct {
	name          string
	executed      bool
	inProgress    bool
	prerequisites []*step
}

type worker struct {
	id        int
	assigment *step
	remaining int
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

func isStepReady(steps []*step) (bool, *step) {
	for i := 0; i < len(steps); i++ {
		if !steps[i].inProgress && steps[i].canExecute() {
			return true, steps[i]
		}
	}
	return false, nil
}

func assignWork(w *worker, steps []*step) {
	if w.assigment == nil {
		ready, stepPointer := isStepReady(steps)
		if ready {
			w.assigment = stepPointer
			w.assigment.inProgress = true
			w.remaining = executionTime + int((stepPointer.name[0])-64)
		}
	}
}

func decreaseRemainingTime(w *worker) {
	if w.assigment != nil {
		w.remaining--
	}
}

func finishWorkIfDone(w *worker) {
	if w.assigment != nil && w.remaining == 0 {
		w.assigment.executed = true
		w.assigment.inProgress = false
		w.assigment = nil
	}
}

func executeSteps(steps []*step) int {
	var totalTime int
	var workers [workersCount]worker
	for totalTime = 1; ; totalTime++ {
		for i := 0; i < workersCount; i++ {
			assignWork(&workers[i], steps)
			decreaseRemainingTime(&workers[i])
		}

		// Check for finished work after assigments for this second are done
		for i := 0; i < workersCount; i++ {
			finishWorkIfDone(&workers[i])
		}

		if allExecuted(steps) {
			return totalTime
		}
	}
}

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")

	var steps []*step

	for _, line := range inputArray {
		steps = createStep(steps, line)
	}

	sort.Sort(sortByName(steps))

	fmt.Printf("Execution time: %d\n", executeSteps(steps))
}
