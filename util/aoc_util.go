package util

import (
	"io/ioutil"
	"strings"
)

// GetInputString loads the input as string
func GetInputString() string {
	data, err := ioutil.ReadFile("input")

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}
