package sort

import (
	"sort"
	"testing"
)

func TestBubble(t *testing.T) {
	in := generateIntSlice(1000, 0xcafebabe)
	want := sortedCopy(in)

	got := Bubble(in)
	if !compareIntSlices(got, want) {
		t.Errorf("didn't get what wanted :(")
	}
}

func BenchSort(b *testing.B, f func(x []int) []int) {
	var seed int64 = 0xdeadbeef

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		in := generateIntSlice(1000, seed)
		seed++
		b.StartTimer()
		_ = f(in)
	}
}

var reference = func(x []int) []int {
	sort.Ints(x)
	return x
}

func BenchmarkReference(b *testing.B) {
	BenchSort(b, reference)
}

func BenchmarkBubble(b *testing.B) {
	BenchSort(b, Bubble[int])
}
