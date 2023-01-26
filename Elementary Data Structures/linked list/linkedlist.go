package LinkedList

// A double, circular linked list implementation
type Node[T any] struct {
	prev, next *Node[T]
	value      T
}

func NewLInkedList[T any](x T) *Node[T] {
	return &Node[T]{nil, nil, x}
}

func (first *Node[T]) Insert(x T) {
	current := first
	for ; current.next != nil; current = current.next {

	}
	current.next = &Node[T]{current, first, x}
}

func (first *Node[T]) Delete(x T) {
	current := first
	for ; current.next != nil; current = current.next {
		if current.value == x {
			current.prev.next = current.next
			current.next.prev = current.prev
			return
		}
	}
}

func (first *Node[T]) Search(x T) bool {
	current := first
	for ; current.next != nil; current = current.next {
		if current.value == x {
			return true
		}
	}
	return false
}
