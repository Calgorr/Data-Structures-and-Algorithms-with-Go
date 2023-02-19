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

func HeapIncreaseKey(heap []int, i, key int) {
	//hope that the user has some sho,oor so that the key is greater than heap[i]
	heap[i] = key
	for i > 0 && heap[i] > heap[i/2] {
		heap[i], heap[(i-1)/2] = heap[(i-1)/2], heap[i]
		i = (i - 1) / 2
	}
}

func MaxHeapInsert(heap []int, key int) {
	heap = append(heap, -1)
	HeapIncreaseKey(heap, len(heap)-1, key)
}
