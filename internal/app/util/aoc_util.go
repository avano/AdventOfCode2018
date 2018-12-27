package util

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	aoc "github.com/avano/AdventOfCode2018/cmd"
	"github.com/spf13/cobra"
)

func defaultFile(caller string) string {
	return path.Join(path.Dir(caller), "../../assets/day"+filepath.Base(caller)[:2])
}

// ReadInput loads the input as string
func ReadInput(file *string, example *bool) string {
	f := *file
	if f == "" {
		_, caller, _, _ := runtime.Caller(1)
		f = defaultFile(caller)
		if *example {
			f = f + "_example"
		}
	}
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic("Unable to read file " + f)
	}

	return strings.TrimRight(string(data), "\n")
}

// RegisterCommand registers a new command for cobra and returns the values for input and example
func RegisterCommand(use string, short string, f func(cmd *cobra.Command, _ []string)) (*string, *bool) {
	var c = &cobra.Command{
		Use:   use,
		Short: short,
		Run:   f,
	}
	aoc.RootCmd.AddCommand(c)

	file := c.Flags().StringP("file", "f", "", "Input file")
	example := c.Flags().BoolP("example", "e", false, "Use example file (dayXX_example) as default file")

	return file, example
}
