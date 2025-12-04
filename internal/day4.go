package adventofcode2025

import (
	"strconv"
	"strings"
)

type Day4 struct {
	rolls map[Position2D]any
	n     int64
	m     int64
}

func (d *Day4) Parse(input string) error {

	rolls := make(map[Position2D]any)
	var n, m int64
	for i, line := range strings.Split(input, "\n") {
		n = int64(i)
		for j, c := range line {
			if c == '@' {
				rolls[Position2D{
					I: int64(i),
					J: int64(j),
				}] = nil
			}
			m = int64(j)
		}
	}

	d.rolls = rolls
	d.n = n + 1
	d.m = m + 1

	return nil
}

func (d *Day4) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64

	for i := range d.n {
		for j := range d.m {
			pos := Position2D{I: i, J: j}
			if _, ok := d.rolls[pos]; !ok {
				continue
			}

			var cnt int64
			for n := range pos.Neighbors(nil, nil, nil, nil) {
				if _, ok := d.rolls[n]; !ok {
					continue
				}
				cnt++
			}

			if cnt < 4 {
				p1++
			}
		}
	}

	changed := true
	for changed {
		changed = false

		var removed []Position2D
		for i := range d.n {
			for j := range d.m {
				pos := Position2D{I: i, J: j}
				if _, ok := d.rolls[pos]; !ok {
					continue
				}

				var cnt int64
				for n := range pos.Neighbors(nil, nil, nil, nil) {
					if _, ok := d.rolls[n]; !ok {
						continue
					}
					cnt++
				}

				if cnt < 4 {
					removed = append(removed, pos)
				}
			}
		}

		for _, roll := range removed {
			delete(d.rolls, roll)
			p2++
		}

		if len(removed) > 0 {
			changed = true
		}
	}

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}
