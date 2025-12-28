package adventofcode2025

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Day8 struct {
	circuitsToJunctions map[int64][]Position3D
	junctionsToCircuits map[Position3D]int64
	distances           map[Position3D]map[Position3D]float64
	distancesSorted     [][]Position3D
}

func (d *Day8) Parse(input string) error {

	circuitsToJunctions := make(map[int64][]Position3D)
	junctionsToCircuits := make(map[Position3D]int64)
	for i, line := range strings.Split(input, "\n") {
		s := strings.Split(line, ",")
		if len(s) != 3 {
			return fmt.Errorf("unexpected format encountered at line %d: %q", i+1, line)
		}

		pos := Position3D{}
		var err error
		for j, p := range []*int64{&pos.I, &pos.J, &pos.K} {
			*p, err = strconv.ParseInt(s[j], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse line %d: %w", i+1, err)
			}
		}

		circuitsToJunctions[int64(i)] = append(circuitsToJunctions[int64(i)], pos)
		junctionsToCircuits[pos] = int64(i)
	}

	distances := make(map[Position3D]map[Position3D]float64)
	var sorted [][]Position3D
	for this := range junctionsToCircuits {
		for that := range junctionsToCircuits {
			if this == that {
				continue
			}

			if distances[this] == nil {
				distances[this] = make(map[Position3D]float64)
			}

			distances[this][that] = this.Euclidian(that)

			sorted = append(sorted, []Position3D{this, that})
		}
	}

	slices.SortFunc(sorted, func(this []Position3D, that []Position3D) int {
		return cmp.Compare(distances[this[0]][this[1]], distances[that[0]][that[1]])
	})

	d.junctionsToCircuits = junctionsToCircuits
	d.circuitsToJunctions = circuitsToJunctions
	d.distances = distances
	d.distancesSorted = sorted

	return nil
}

func (d *Day8) Solve() (part1 string, part2 string, err error) {
	const connections = 1000
	var p1 int64
	var p2 int64

	for i := 0; i < 2*connections; i += 2 {
		this := d.distancesSorted[i][0]
		that := d.distancesSorted[i][1]

		if d.junctionsToCircuits[this] == d.junctionsToCircuits[that] {
			continue
		}

		thisCircuit := d.junctionsToCircuits[this]
		thatCircuit := d.junctionsToCircuits[that]

		for _, j := range d.circuitsToJunctions[thisCircuit] {
			d.junctionsToCircuits[j] = thatCircuit
		}

		d.circuitsToJunctions[thatCircuit] = append(d.circuitsToJunctions[thatCircuit], d.circuitsToJunctions[thisCircuit]...)
		delete(d.circuitsToJunctions, thisCircuit)

	}

	var circuitsSizes []int
	for _, junctions := range d.circuitsToJunctions {
		circuitsSizes = append(circuitsSizes, len(junctions))
	}
	slices.Sort(circuitsSizes)

	p1 = 1
	for i := range 3 {
		p1 *= int64(circuitsSizes[len(circuitsSizes)-1-i])
	}

	for i := 2 * connections; len(d.circuitsToJunctions) > 1; i += 2 {
		this := d.distancesSorted[i][0]
		that := d.distancesSorted[i][1]

		if d.junctionsToCircuits[this] == d.junctionsToCircuits[that] {
			continue
		}

		thisCircuit := d.junctionsToCircuits[this]
		thatCircuit := d.junctionsToCircuits[that]

		for _, j := range d.circuitsToJunctions[thisCircuit] {
			d.junctionsToCircuits[j] = thatCircuit
		}

		d.circuitsToJunctions[thatCircuit] = append(d.circuitsToJunctions[thatCircuit], d.circuitsToJunctions[thisCircuit]...)
		delete(d.circuitsToJunctions, thisCircuit)

		if len(d.circuitsToJunctions) <= 1 {
			p2 = this.I * that.I
			break
		}
	}

	return strconv.FormatInt(p1, 10), strconv.FormatInt(p2, 10), nil
}
