package adventofcode2025

import (
	"fmt"
	"strconv"
	"strings"
)

type Day6 struct {
	numbers [][]int64
	ops     []rune
	raw     [][]rune
}

func (d *Day6) Parse(input string) error {
	var numbers [][]int64
	var ops []rune
	var raw [][]rune
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		raw = append(raw, []rune(line))
		if i < len(lines)-1 {
			var nums []int64
			for s := range strings.FieldsSeq(line) {
				num, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return fmt.Errorf("could not parse %q as int: %w", s, err)
				}

				nums = append(nums, num)
			}
			numbers = append(numbers, nums)
		} else {
			ops = []rune(strings.ReplaceAll(line, " ", ""))
		}
	}

	d.numbers = numbers
	d.ops = ops
	d.raw = raw

	return nil
}

func (d *Day6) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64

	for j := range len(d.numbers[0]) {
		var total int64
		switch d.ops[j] {
		case '*':
			total = 1
		case '+':
			total = 0
		default:
			return "", "", fmt.Errorf("unexpected operant %q found", d.ops[j])
		}
		for i := range len(d.numbers) {
			switch d.ops[j] {
			case '*':
				total *= d.numbers[i][j]
			case '+':
				total += d.numbers[i][j]
			default:
				return "", "", fmt.Errorf("unexpected operant %q found", d.ops[j])
			}
		}
		p1 += total
	}

	var total int64
	var op rune
	for j := range len(d.raw[0]) {
		numStr := ""
		for i := range len(d.raw) - 1 {
			numStr += string(d.raw[i][j])
		}
		numStr = strings.TrimSpace(numStr)

		if last := d.raw[len(d.raw)-1][j]; last != ' ' {
			op = last
			switch op {
			case '*':
				total = 1
			case '+':
				total = 0
			default:
				return "", "", fmt.Errorf("unexpected operant %q found", d.ops[j])
			}
		}

		if numStr == "" {
			p2 += total
		} else {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				return "", "", fmt.Errorf("failed to parse %q as int: %w", numStr, err)
			}

			if op == '*' {
				total *= num
			} else {
				total += num
			}
		}

	}
	p2 += total

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}
