package sort

import (
	"golang.org/x/exp/constraints"
)

// Insertion sorts the slice x using insertion sort algorithm.
//
// https://en.wikipedia.org/wiki/Insertion_sort
func Insertion[T constraints.Ordered](x []T) {
	n := len(x)

	for i := 1; i < n; i++ {
		me := x[i]

		j := i - 1
		for ; j >= 0 && me < x[j]; j-- {
			x[j+1] = x[j]
		}

		if j != i {
			x[j+1] = me
		}
	}
}
