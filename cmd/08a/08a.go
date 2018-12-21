package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/avano/AdventOfCode2018/internal/app/util"
)

type node struct {
	children []*node
	metadata []int
}

func parseNode(input []string, index int) (*node, int) {
	childrenCount, err := strconv.Atoi(string(input[index]))
	metadataCount, err2 := strconv.Atoi(string(input[index+1]))
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}

	n := node{}

	readIndex := index + 2
	for j := 0; j < childrenCount; j++ {
		child, lastIndex := parseNode(input, readIndex)
		n.children = append(n.children, child)
		readIndex = lastIndex
	}

	for k := 0; k < metadataCount; k++ {
		m, err := strconv.Atoi(string(input[readIndex]))
		if err != nil {
			panic(err)
		}
		n.metadata = append(n.metadata, m)
		readIndex++
	}

	return &n, readIndex
}

func countNodeMetadata(n *node) int {
	m := 0
	for _, num := range n.metadata {
		m += num
	}
	return m
}

func countMetadata(n *node) int {
	if len(n.children) == 0 {
		return countNodeMetadata(n)
	} else {
		sum := 0
		for i := 0; i < len(n.children); i++ {
			sum += countMetadata(n.children[i])
		}
		return countNodeMetadata(n) + sum
	}
}

func main() {
	inputArray := strings.Split(util.GetInputString(), " ")

	n, _ := parseNode(inputArray, 0)

	fmt.Println(countMetadata(n))
}
