package main

import (
	"fmt"

	Queue "github.com/Calgorr/Data-Structures-and-Algorithms-with-Go/ElementaryDataStructures/queue/queue"
)

func main() {
	queue := Queue.NewQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
