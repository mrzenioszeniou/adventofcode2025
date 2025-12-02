package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

func main() {
	var dayArg uint8
	var pathArg string

	cmd := &cobra.Command{
		Use:          "adventofcode2025",
		Short:        "Advent of Code - 2025 solver",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			var problem Day
			switch dayArg {
			case 1:
				problem = &Day1{}
			case 2:
				problem = &Day2{}
			default:
				return fmt.Errorf("no solution available for day %d", dayArg)
			}

			bytes, err := os.ReadFile(pathArg)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}

			if !utf8.Valid(bytes) {
				return fmt.Errorf("file %s is not utf8", pathArg)
			}

			if err := problem.Parse(string(bytes)); err != nil {
				return err
			}

			p1, p2, err := problem.Solve()
			if err != nil {
				return fmt.Errorf("failed to solve day %d: %w", dayArg, err)
			}

			fmt.Printf("Part 1: %s\nPart 2: %s\n", p1, p2)

			return nil
		},
		SilenceErrors: true,
	}

	cmd.Flags().Uint8VarP(&dayArg, "day", "d", 0, "day to solve")
	cmd.MarkFlagRequired("day")
	cmd.Flags().StringVarP(&pathArg, "path", "p", "", "path to input file")
	cmd.MarkFlagRequired("path")

	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

type Day interface {
	Parse(string) error
	Solve() (string, string, error)
}
