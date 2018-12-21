package main

import (
	"fmt"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

type record struct {
	next   *record
	prev   *record
	number int
}

type marbleList struct {
	head    *record
	tail    *record
	current *record
}

func (list *marbleList) insert(rec *record) {
	// if there is no head, the list is empty
	if list.head == nil {
		list.head = rec
		list.tail = rec
	} else {
		// if the next element is head, this one is tail and vice versa
		if list.head == rec.next {
			list.tail = rec
		}
		if list.tail == rec.prev {
			list.head = rec
		}
		// change the pointer to include this element
		rec.prev.next = rec
		rec.next.prev = rec
	}
	list.current = rec
}

func (list *marbleList) insertMarble(num int) int {
	score := 0
	marble := record{number: num}

	if marble.number%23 == 0 {
		score += marble.number
		toRemove := list.current.prev.prev.prev.prev.prev.prev.prev
		score += toRemove.number
		// change the current to a next one
		list.current = toRemove.next
		toRemove.prev.next = toRemove.next
		toRemove.next.prev = toRemove.prev
		// change head and tail if necessary
		if toRemove == list.head {
			list.head = toRemove.next
		}
		if toRemove == list.tail {
			list.tail = toRemove.prev
		}

		return score
	}

	marble.prev = list.current.next
	marble.next = list.current.next.next
	list.insert(&marble)
	return score
}

func main() {
	input := util.GetInputString()

	var players, lastMarble int

	_, err := fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &lastMarble)

	if err != nil {
		panic(err)
	}

	init := record{number: 0}
	init.next = &init
	init.prev = &init
	list := marbleList{}
	list.insert(&init)

	playerScore := make(map[int]int)
	for i := 0; i < lastMarble; i++ {
		playerScore[i%players+1] += list.insertMarble(i + 1)
	}

	max := 0
	for _, v := range playerScore {
		if v > max {
			max = v
		}
	}
	fmt.Printf("High score: %d\n", max)
}
