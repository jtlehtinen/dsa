package sort

import (
	"sort"
	"testing"
)

func testSort(t *testing.T, f func(x []int) []int) {
	t.Helper()

	for i := 0; i < 10; i++ {
		in := generateIntSlice(1000, 0xcafebabe)
		want := sortedCopy(in)

		got := f(in)
		if !compareIntSlices(got, want) {
			t.Errorf("didn't get what wanted :(")
		}
	}
}

func TestBubble(t *testing.T) {
	testSort(t, Bubble[int])
}

func TestInsertion(t *testing.T) {
	testSort(t, Insertion[int])
}

func TestSelection(t *testing.T) {
	testSort(t, Selection[int])
}

func benchSort(b *testing.B, f func(x []int) []int) {
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
	benchSort(b, reference)
}

func BenchmarkBubble(b *testing.B) {
	benchSort(b, Bubble[int])
}

func BenchmarkInsertion(b *testing.B) {
	benchSort(b, Insertion[int])
}

func BenchmarkSelection(b *testing.B) {
	benchSort(b, Selection[int])
}
