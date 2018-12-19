package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

type guard struct {
	id    int
	sleep map[int]int
}

func getID(input string) int {
	id := strings.Fields(input)[3]
	idInt, err := strconv.Atoi(id[1:len(id)])
	if err != nil {
		panic(err)
	}
	return int(idInt)
}

func getMinute(input string) int {
	minuteString := input[strings.Index(input, ":")+1 : strings.Index(input, ":")+3]
	minuteInt, err := strconv.Atoi(minuteString)
	if err != nil {
		panic(err)
	}
	return int(minuteInt)
}

func processGuard(guards []guard, timetable []string) guard {
	var gu guard
	id := getID(timetable[0])
	for _, g := range guards {
		if g.id == id {
			gu = g
			break
		}
	}

	if gu.id == 0 {
		gu = guard{id: id, sleep: make(map[int]int)}
	}

	for i := 1; i < len(timetable); i += 2 {
		sleepMinute := getMinute(timetable[i])
		wakeMinute := getMinute(timetable[i+1])
		for j := sleepMinute; j < wakeMinute; j++ {
			gu.sleep[j]++
		}
	}

	return gu
}

func main() {
	inputArray := strings.Split(util.GetInputString(), "\n")
	sort.Strings(inputArray)

	var guards []guard
	lastIndex := 0
	for i := 0; i < len(inputArray); i++ {
		if i == len(inputArray)-1 || strings.Contains(inputArray[i+1], "Guard") {
			if len(inputArray[lastIndex:i]) > 0 {
				guards = append(guards, processGuard(guards, inputArray[lastIndex:i+1]))
			}
			lastIndex = i + 1
		}
	}

	guardID := 0
	mostSleepyMinute := 0
	max := 0
	for _, g := range guards {
		for k, v := range g.sleep {
			if v > max {
				mostSleepyMinute = k
				max = v
				guardID = g.id
			}
		}
	}

	fmt.Printf("Most sleepy minute %d (%d times) by Guard #%d\n", mostSleepyMinute, max, guardID)
	fmt.Printf("Result: %d\n", guardID*mostSleepyMinute)
}
