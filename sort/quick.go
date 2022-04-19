package sort

import (
	"golang.org/x/exp/constraints"
)

// Quick sorts the slice x using quicksort algorithm.
//
// https://en.wikipedia.org/wiki/Quicksort
func Quick[T constraints.Ordered](x []T) {
	// Slice having length less than 2 is already
	// ordered.
	if len(x) < 2 {
		return
	}

	// Pivot is the first element in the slice.
	// @TODO: Experiment with different pivots.
	p := 1
	for i := 1; i < len(x); i++ {
		if x[i] < x[0] {
			x[p], x[i] = x[i], x[p]
			p++
		}
	}

	// Invariant: p is 1 past the last element
	// less than pivot.
	x[0], x[p-1] = x[p-1], x[0]

	// The pivot is in the correct place and its
	// index is (p-1).
	Quick(x[:p-1])
	Quick(x[p:])
}
