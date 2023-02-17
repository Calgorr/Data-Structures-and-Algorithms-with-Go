package heap

func MaxHeapify(heap []int, heap_size, index int) {
	l, r := 2*index, 2*index+1
	largest := index
	if l < heap_size && heap[l] > heap[largest] {
		largest = l
	} else if r < heap_size && heap[r] > heap[largest] {
		largest = r
	}
	heap[index], heap[largest] = heap[largest], heap[index]
	if largest != index {
		MaxHeapify(heap, heap_size, largest)
	}
}
