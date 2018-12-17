package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")

	if err != nil {
		panic(err)
	}

	frequency := 0

	stringArray := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i := 0; i < len(stringArray); i++ {
		num, err := strconv.ParseInt(stringArray[i], 0, 0)
		if err != nil {
			panic(err)
		}
		frequency += int(num)
	}

	fmt.Println("Frequency: ", frequency)
}
