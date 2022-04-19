package sort

import "golang.org/x/exp/constraints"

const (
	InvalidIndex = -1
)

type MinHeap[T constraints.Ordered] []T

func parentIndex(childIndex int) int {
	if childIndex <= 0 {
		return InvalidIndex
	}
	return (childIndex - 1) / 2
}

func leftChildIndex(parentIndex, heapSize int) int {
	if childIndex := parentIndex*2 + 1; childIndex < heapSize {
		return childIndex
	}
	return InvalidIndex
}

func rightChildIndex(parentIndex, heapSize int) int {
	if childIndex := parentIndex*2 + 2; childIndex < heapSize {
		return childIndex
	}
	return InvalidIndex
}

func indexOfSmallerChild[T constraints.Ordered](parentIndex int, h MinHeap[T]) int {
	l := leftChildIndex(parentIndex, len(h))
	r := rightChildIndex(parentIndex, len(h))
	if r != InvalidIndex && h[r] < h[l] {
		return r
	}
	return l
}

func heapifyUp[T constraints.Ordered](h MinHeap[T]) MinHeap[T] {
	i := len(h) - 1

	for j := parentIndex(i); j != InvalidIndex && h[i] < h[j]; i, j = j, parentIndex(j) {
		h[i], h[j] = h[j], h[i]
	}

	return h
}

func heapifyDown[T constraints.Ordered](h MinHeap[T]) MinHeap[T] {
	p := 0
	c := indexOfSmallerChild(p, h)

	for ; c != InvalidIndex && h[c] < h[p]; p, c = c, indexOfSmallerChild(c, h) {
		h[p], h[c] = h[c], h[p]
	}

	return h
}

func pop[T constraints.Ordered](h MinHeap[T]) (T, MinHeap[T]) {
	first := h[0]
	h[0] = h[len(h)-1]
	return first, heapifyDown(h[:len(h)-1])
}

func buildMinHeap[T constraints.Ordered](arr []T) MinHeap[T] {
	h := make(MinHeap[T], 0, len(arr))
	for _, v := range arr {
		h = append(h, v)
		h = heapifyUp(h)
	}
	return h
}

// Heap sorts the slice x using heapsort algorithm.
//
// https://en.wikipedia.org/wiki/Heapsort
func Heap[T constraints.Ordered](x []T) {
	h := buildMinHeap(x)
	for i := 0; len(h) != 0; i++ {
		x[i], h = pop(h)
	}
}
