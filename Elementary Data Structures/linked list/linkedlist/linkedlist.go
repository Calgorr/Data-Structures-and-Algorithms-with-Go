package LinkedList

// A double linked list implementation
type Node[V comparable] struct {
	Prev, Next *Node[V]
	Value      V
}

func NewLInkedList[V comparable](x V) *Node[V] {
	return &Node[V]{nil, nil, x}
}

func (first *Node[V]) Insert(x V) {
	current := first
	for ; current.Next != nil; current = current.Next {
	}
	current.Next = &Node[V]{current, nil, x}
}

func (first *Node[V]) Delete(x V) {
	current := first
	for ; current.Next != nil; current = current.Next {
		if current.Value == x {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
			return
		}
	}
}

func (first *Node[V]) Search(x V) bool {
	current := first
	for ; current.Next != nil; current = current.Next {
		if current.Value == x {
			return true
		}
	}
	return false
}
