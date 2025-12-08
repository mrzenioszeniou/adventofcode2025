package adventofcode2025

import (
	"fmt"
	"strconv"
	"strings"
)

type Day7 struct {
	splitters map[Position2D]any
	n         int64
	m         int64
	start     Position2D
}

func (d *Day7) Parse(input string) error {
	lines := strings.Split(input, "\n")
	var start *Position2D
	var n, m int64
	splitters := make(map[Position2D]any)
	for i, line := range lines {
		for j, c := range line {
			switch c {
			case 'S':
				start = &Position2D{
					I: int64(i),
					J: int64(j),
				}
			case '^':
				splitters[Position2D{
					I: int64(i),
					J: int64(j),
				}] = nil
			case '.':
				n = int64(i + 1)
				m = max(m, int64(j+1))
			default:
				return fmt.Errorf("unexpected character %q encountered", c)
			}
		}
	}

	d.n = n
	d.m = m
	d.start = *start
	d.splitters = splitters

	return nil
}

func (d *Day7) Solve() (part1 string, part2 string, err error) {
	var p1 int64
	var p2 int64

	beams := map[Position2D]any{d.start: nil}
	splits := map[Position2D]any{}
	for len(beams) > 0 {
		newBeams := make(map[Position2D]any)
		for beam := range beams {
			beam.I++
			if beam.I >= d.n {
				continue
			}

			if _, ok := d.splitters[beam]; ok {
				splits[beam] = nil

				if beam.J > 0 {
					newBeams[Position2D{
						I: beam.I,
						J: beam.J - 1,
					}] = nil
				}

				if beam.J < d.n-1 {
					newBeams[Position2D{
						I: beam.I,
						J: beam.J + 1,
					}] = nil
				}
			} else {
				newBeams[Position2D{
					I: beam.I,
					J: beam.J,
				}] = nil
			}
		}

		beams = newBeams
	}

	p1 = int64(len(splits))
	p2 = d.realities(d.start, make(map[Position2D]int64))

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}

func (d *Day7) realities(beam Position2D, hist map[Position2D]int64) int64 {
	if c, ok := hist[beam]; ok {
		return c
	}

	var ret int64
	if beam.J >= d.m || beam.J < 0 {
		ret = 0
	} else if beam.I+1 >= d.n {
		ret = 1
	} else if _, ok := d.splitters[Position2D{
		I: beam.I + 1,
		J: beam.J,
	}]; ok {
		ret = d.realities(Position2D{
			I: beam.I + 1,
			J: beam.J - 1,
		}, hist) +
			d.realities(Position2D{
				I: beam.I + 1,
				J: beam.J + 1,
			}, hist)
	} else {
		ret = d.realities(Position2D{
			I: beam.I + 1,
			J: beam.J,
		}, hist)
	}

	hist[beam] = ret
	return ret
}
