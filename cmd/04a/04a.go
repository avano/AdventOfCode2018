package day04a

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
	"github.com/spf13/cobra"
)

var file *string
var example *bool

func init() {
	file, example = util.RegisterCommand("day04a", "Day 4 - First Part", run)
}

type guard struct {
	id    int
	sleep map[int]int
}

func getID(input string) int {
	id := strings.Fields(input)[3]
	idInt, err := strconv.Atoi(id[1:])
	if err != nil {
		panic(err)
	}
	return int(idInt)
}

func getMinute(input string) int {
	colonIndex := strings.Index(input, ":")
	minuteString := input[colonIndex+1 : colonIndex+3]
	minuteInt, err := strconv.Atoi(minuteString)
	if err != nil {
		panic(err)
	}
	return int(minuteInt)
}

func processGuard(guards []guard, timetable []string) guard {
	var g guard
	id := getID(timetable[0])
	for i := range guards {
		if guards[i].id == id {
			g = guards[i]
			break
		}
	}

	if g.id == 0 {
		g = guard{id: id, sleep: make(map[int]int)}
	}

	for i := 1; i < len(timetable); i += 2 {
		sleepMinute := getMinute(timetable[i])
		wakeMinute := getMinute(timetable[i+1])
		for j := sleepMinute; j < wakeMinute; j++ {
			g.sleep[j]++
		}
	}

	return g
}

func run(cmd *cobra.Command, _ []string) {
	input := strings.Split(util.ReadInput(file, example), "\n")
	sort.Strings(input)

	var guards []guard
	lastIndex := 0
	for i := 0; i < len(input); i++ {
		if i == len(input)-1 || strings.Contains(input[i+1], "Guard") {
			if len(input[lastIndex:i]) > 0 {
				guards = append(guards, processGuard(guards, input[lastIndex:i+1]))
			}
			lastIndex = i + 1
		}
	}

	mostSleepyGuardIndex := 0
	max := 0
	for i := range guards {
		guardMax := 0
		for _, v := range guards[i].sleep {
			guardMax += v
		}
		if guardMax > max {
			mostSleepyGuardIndex = i
			max = guardMax
		}
	}

	mostSleepyMinute := 0
	max = 0
	for k, v := range guards[mostSleepyGuardIndex].sleep {
		if v > max {
			mostSleepyMinute = k
			max = v
		}
	}

	fmt.Printf("Result: %d\n", guards[mostSleepyGuardIndex].id*mostSleepyMinute)
}
