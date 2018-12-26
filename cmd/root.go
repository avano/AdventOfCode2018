package aoc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd root command
var RootCmd *cobra.Command

func init() {
	RootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "AOC2018",
		Long:  "Advent of Code 2018",
	}
}

// Execute executes cobra
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
