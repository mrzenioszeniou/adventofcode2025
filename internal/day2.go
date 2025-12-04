package adventofcode2025

import (
	"fmt"
	"strconv"
	"strings"
)

type Day2 struct {
	rangs []rang
}

type rang struct {
	from int64
	to   int64
}

func (d *Day2) Parse(input string) error {
	var rangs []rang
	for rangeStr := range strings.SplitSeq(input, ",") {
		numStrs := strings.Split(rangeStr, "-")
		if len(numStrs) != 2 {
			return fmt.Errorf("invalid range %q encountered", rangeStr)
		}

		from, err := strconv.ParseInt(numStrs[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid range %q encountered: %w", rangeStr, err)
		}

		to, err := strconv.ParseInt(numStrs[1], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid range %q encountered: %w", rangeStr, err)
		}

		rangs = append(rangs, rang{
			from: from,
			to:   to,
		})
	}

	d.rangs = rangs

	return nil
}

func (d *Day2) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64

	for _, rang := range d.rangs {
		for n := rang.from; n <= rang.to; n++ {
			str := strconv.FormatInt(n, 10)
			if len(str)%2 == 0 && str[:len(str)/2] == str[len(str)/2:] {
				p1 += n
			}

			for l := range len(str) / 2 {
				if strings.Count(str, str[:l+1])*(l+1) == len(str) {
					p2 += n
					break
				}
			}
		}
	}

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}
