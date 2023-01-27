package LinkedList

// A double, circular linked list implementation
type Node[V comparable] struct {
	prev, next *Node[V]
	value      V
}

func NewLInkedList[V comparable](x V) *Node[V] {
	return &Node[V]{nil, nil, x}
}

func (first *Node[V]) Insert(x V) {
	current := first
	for ; current.next != nil; current = current.next {

	}
	current.next = &Node[V]{current, first, x}
}

func (first *Node[V]) Delete(x V) {
	current := first
	for ; current.next != nil; current = current.next {
		if current.value == x {
			current.prev.next = current.next
			current.next.prev = current.prev
			return
		}
	}
}

func (first *Node[V]) Search(x V) bool {
	current := first
	for ; current.next != nil; current = current.next {
		if current.value == x {
			return true
		}
	}
	return false
}
