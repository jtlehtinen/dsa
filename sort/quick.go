package sort

import (
	"golang.org/x/exp/constraints"
)

func Quick[T constraints.Ordered](x []T) []T {
	return x
}
