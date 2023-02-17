package heap

func MaxHeapify(heap []int, heap_size, index int) {
	l, r := 2*(index+1), 2*(index+1)+1
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

func Build_Max_Heap(heap []int) {
	len := len(heap)
	for i := len/2 - 1; i >= 0; len-- {
		MaxHeapify(heap, len, i)
	}
}

func HeapSort(heap []int) {
	Build_Max_Heap(heap)
	for i := len(heap) - 1; i >= 0; i-- {
		heap[1], heap[i] = heap[i], heap[1]
		MaxHeapify(heap, len(heap)-1, 1)
	}
}
