package sort

import (
	"sort"
	"testing"
)

func testSort(t *testing.T, f func(x []int)) {
	t.Helper()

	for i := 0; i < 10; i++ {
		in := generateIntSlice(100, 0xcafebabe)
		want := sortedCopy(in)

		f(in)
		got := in

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

func TestQuick(t *testing.T) {
	testSort(t, Quick[int])
}

func benchSort(b *testing.B, f func(x []int)) {
	var seed int64 = 0xdeadbeef

	for n := 0; n < b.N; n++ {
		b.StopTimer()
		in := generateIntSlice(1000, seed)
		seed++
		b.StartTimer()
		f(in)
	}
}

func BenchmarkReference(b *testing.B) {
	benchSort(b, sort.Ints)
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

func BenchmarkQuick(b *testing.B) {
	benchSort(b, Quick[int])
}
