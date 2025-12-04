package adventofcode2025

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day1 struct {
	rotations []rotation
}

type direction int

const (
	left direction = iota
	right
)

type rotation struct {
	amount    int64
	direction direction
}

func parseRotation(s string) (rotation, error) {
	rot := rotation{}
	switch s[0] {
	case 'L':
		rot.direction = left
	case 'R':
		rot.direction = right
	default:
		return rotation{}, errors.New("invalid direction")
	}

	amount, err := strconv.ParseInt(s[1:], 10, 64)
	if err != nil {
		return rotation{}, err
	}

	rot.amount = amount

	return rot, nil
}

func (d *Day1) Parse(input string) error {
	var rots []rotation

	for line := range strings.SplitSeq(input, "\n") {
		if rot, err := parseRotation(line); err == nil {
			rots = append(rots, rot)
		} else {
			return fmt.Errorf("invalid rotation %q encountered: %w", line, err)
		}
	}

	d.rotations = rots

	return nil
}

func (d *Day1) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64
	var curr int64 = 50
	for _, rot := range d.rotations {
		prev := curr
		p2 += rot.amount / 100
		amount := rot.amount % 100
		switch rot.direction {
		case left:
			curr -= amount
		case right:
			curr += amount
		}

		if curr == 0 {
			p1++
			p2++
		} else if curr < 0 {
			curr += 100
			if prev != 0 {
				p2++
			}
		} else if curr >= 100 {
			curr -= 100
			p2++
		}
	}

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}
