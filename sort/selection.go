package sort

import (
	"golang.org/x/exp/constraints"
)

func Selection[T constraints.Ordered](x []T) {
	n := len(x)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if x[j] < x[min] {
				min = j
			}
		}
		x[i], x[min] = x[min], x[i]
	}
}
