package adventofcode2025

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	ranges      []Rang
	ingredients []int64
}

func (d *Day5) Parse(input string) error {

	parts := strings.Split(input, "\n\n")

	if len(parts) != 2 {
		return fmt.Errorf("unexpected input parts (%d) found", len(parts))
	}

	var ranges []Rang
	for line := range strings.SplitSeq(parts[0], "\n") {
		rangeStr := strings.Split(line, "-")
		if len(rangeStr) != 2 {
			return fmt.Errorf("unexpected range found %q", line)
		}

		from, err := strconv.ParseInt(rangeStr[0], 10, 64)
		if err != nil {
			return fmt.Errorf("unexpected range found %q: %w", line, err)
		}

		to, err := strconv.ParseInt(rangeStr[1], 10, 64)
		if err != nil {
			return fmt.Errorf("unexpected range found %q: %w", line, err)
		}

		ranges = append(ranges, Rang{
			from: from,
			to:   to,
		})
	}

	var ingredients []int64
	for line := range strings.SplitSeq(parts[1], "\n") {
		ingredient, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return fmt.Errorf("unexpected ingredient found %q: %w", line, err)
		}

		ingredients = append(ingredients, ingredient)
	}

	d.ranges = ranges
	d.ingredients = ingredients

	return nil
}

func (d *Day5) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64

	for _, ingredient := range d.ingredients {
		var fresh bool
		for _, rang := range d.ranges {
			if ingredient >= rang.from && ingredient <= rang.to {
				fresh = true
				break
			}
		}

		if fresh {
			p1++
		}
	}

	var mergedRanges []Rang = slices.Clone(d.ranges)
	for len(mergedRanges) > 1 {
		var mergedSomething bool

		for i := range mergedRanges {
			for j := i + 1; j < len(mergedRanges); j++ {
				this, that := mergedRanges[i].Union(mergedRanges[j])
				if that == nil {
					mergedRanges = append(mergedRanges[:j], mergedRanges[j+1:]...)
					mergedRanges = append(mergedRanges[:i], mergedRanges[i+1:]...)
					mergedRanges = append(mergedRanges, this)
					mergedSomething = true
					goto Out
				}
			}
		}

	Out:
		if !mergedSomething {
			break
		}
	}

	for _, rang := range mergedRanges {
		p2 += rang.to - rang.from + 1
	}

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}
