package adventofcode2025

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day3 struct {
	banks [][]int64
}

func (d *Day3) Parse(input string) error {
	var banks [][]int64
	for line := range strings.SplitSeq(input, "\n") {
		var bank []int64
		for _, c := range line {
			bank = append(bank, int64(c-'0'))
		}
		banks = append(banks, bank)
	}

	d.banks = banks

	return nil
}

func (d *Day3) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64

	for _, bank := range d.banks {
		p1 += maxJoltage(bank, 2)
		p2 += maxJoltage(bank, 12)

	}

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}

func maxJoltage(bank []int64, batteries int) int64 {
	var digits []int
	for digit := range batteries {
		var start int
		if len(digits) == 0 {
			start = 0
		} else {
			start = digits[len(digits)-1] + 1
		}

		max := start
		for i := start; i < len(bank)-int(batteries)+digit+1; i++ {
			if bank[i] > bank[max] {
				max = i
			}
		}

		digits = append(digits, max)
	}

	slices.Reverse(digits)
	var ret int64
	for pow, i := range digits {
		ret += int64(math.Pow10(pow)) * bank[i]
	}

	return ret
}
