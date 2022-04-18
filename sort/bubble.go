package sort

import (
	"golang.org/x/exp/constraints"
)

// Bubble sorts the slice x using bubble sort algorithm.
//
// https://en.wikipedia.org/wiki/Bubble_sort
func Bubble[T constraints.Ordered](x []T) {
	swapped := true
	for last := len(x); swapped; last-- {
		swapped = false
		for i := 1; i < last; i++ {
			if x[i] < x[i-1] {
				x[i-1], x[i] = x[i], x[i-1]
				swapped = true
			}
		}
	}
}
