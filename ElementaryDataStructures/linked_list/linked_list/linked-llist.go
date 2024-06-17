package linkedlist

type DoublyNode[V comparable] struct {
	Value V
	Next  *DoublyNode[V]
	Prev  *DoublyNode[V]
}

func NewDoublyNode[V comparable](value V) *DoublyNode[V] {
	return &DoublyNode[V]{Value: value}
}

func (head *DoublyNode[V]) InsertAtHead(value V) *DoublyNode[V] {
	newNode := NewDoublyNode(value)
	newNode.Next = head
	head.Prev = newNode
	return newNode
}

func (head *DoublyNode[V]) InsertAtTail(value V) {
	newNode := NewDoublyNode(value)
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	newNode.Prev = current
}

func (head *DoublyNode[V]) InsertAt(index int, value V) {
	if index == 0 {
		head.InsertAtHead(value)
		return
	}
	newNode := NewDoublyNode(value)
	var prev, next *DoublyNode[V]
	prev = head
	for i := 0; i < index-1; i++ {
		if prev == nil {
			return
		}
		prev = prev.Next
	}
	if prev.Next == nil {
		head.InsertAtTail(value)
		return
	}
	next = prev.Next
	prev.Next = newNode
	next.Prev = newNode
	newNode.Prev = prev
	newNode.Next = next
}

func (head *DoublyNode[V]) Find(value V) bool {
	current := head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (head *DoublyNode[V]) Delete(value V) (*DoublyNode[V], bool) {
	current := head
	for current != nil {
		if current.Value == value {
			if current.Next == nil && current.Prev == nil {
				return nil, true
			} else if current.Prev == nil {
				head = current.Next
				head.Prev = nil
				return head, true
			} else if current.Next == nil {
				current.Prev.Next = nil
				return head, true
			}
			prev := current.Prev
			next := current.Next
			prev.Next = next
			next.Prev = prev
			return head, true
		}
		current = current.Next
	}
	return head, false
}

func (head *DoublyNode[V]) Print() {
	current := head
	for current != nil {
		print(current.Value, " ")
		current = current.Next
	}
	println()
}
