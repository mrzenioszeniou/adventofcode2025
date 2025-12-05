package adventofcode2025

import "iter"

type Position2D struct {
	I int64
	J int64
}

func (p Position2D) Neighbors(min_i *int64, max_i *int64, min_j *int64, max_j *int64) iter.Seq[Position2D] {
	return func(yield func(Position2D) bool) {
		for _, i := range []int64{-1, 0, 1} {
			for _, j := range []int64{-1, 0, 1} {
				if i == 0 && j == 0 {
					continue
				}

				neighbor := Position2D{
					I: p.I + i,
					J: p.J + j,
				}

				if min_i != nil && neighbor.I < *min_i {
					continue
				}

				if max_i != nil && neighbor.I > *max_i {
					continue
				}

				if min_j != nil && neighbor.J < *min_j {
					continue
				}

				if max_j != nil && neighbor.J > *max_j {
					continue
				}

				if !yield(neighbor) {
					return
				}
			}
		}
	}
}

type Rang struct {
	from int64
	to   int64
}

func (r0 Rang) Union(r1 Rang) (Rang, *Rang) {
	if r0.to < r1.from || r0.from > r1.to {
		return r0, &r1
	}

	return Rang{
		from: min(r0.from, r1.from),
		to:   max(r0.to, r1.to),
	}, nil
}
