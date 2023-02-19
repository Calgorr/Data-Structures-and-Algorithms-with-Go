package prorityqueue

import (
	Heap "github.com/Calgorr/Data-Structures-and-Algorithms-with-Go/Heap"
)

type Number interface {
	int | interface{}
}

func HeapMax(heap []int) Number {
	return heap[0]
}

func HeapExtractMax(heap []int) Number {
	if len(heap) < 1 {
		return nil
	}
	max := HeapMax(heap)
	heap[0] = heap[len(heap)-1]
	heap = heap[1:]
	Heap.MaxHeapify(heap, len(heap), 0)
	return max
}
