package sort

import (
	"math/rand"
	"sort"
)

func compareIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func generateIntSlice(count int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	s := make([]int, count)

	for i := 0; i < count; i++ {
		s[i] = r.Int()
	}

	return s
}

func sortedCopy(x []int) []int {
	c := make([]int, len(x))
	copy(c, x)
	sort.Ints(c)
	return c
}
