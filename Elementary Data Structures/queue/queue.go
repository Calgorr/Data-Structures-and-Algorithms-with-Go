package Queue

import "fmt"

type Queue[T any] struct {
	values []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.values = append(queue.values, value)
}

func (queue *Queue[T]) Dequeue() (T, error) {
	var x T
	if queue.IsEmpty() {
		return x, fmt.Errorf("Queue is empty")
	}
	x, queue.values = queue.values[0], queue.values[1:]
	return x, nil

}

func (queue *Queue[T]) IsEmpty() bool {
	if len(queue.values) == 0 {
		return true
	}
	return false
}
