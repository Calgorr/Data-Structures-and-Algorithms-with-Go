package linkedlist

import (
	"testing"
)

func TestInsertAtHead(t *testing.T) {
	head := NewDoublyNode(10)
	head = head.InsertAtHead(20)

	if head.Prev != nil {
		t.Errorf("Expected head.Prev to be nil, got %v", head.Prev)
	}
	if head.Value != 20 {
		t.Errorf("Expected head.Value to be 20, got %v", head.Value)
	}
	if head.Next.Value != 10 {
		t.Errorf("Expected head.Next.Value to be 10, got %v", head.Next.Value)
	}
}

func TestInsertAtTail(t *testing.T) {
	head := NewDoublyNode(10)
	head.InsertAtTail(20)

	if head.Next.Value != 20 {
		t.Errorf("Expected head.Next.Value to be 20, got %v", head.Next.Value)
	}
	if head.Next.Prev != head {
		t.Errorf("Expected head.Next.Prev to be head, got %v", head.Next.Prev)
	}
}

func TestInsertAt(t *testing.T) {
	head := NewDoublyNode(10)
	head.InsertAtTail(30)
	head.InsertAt(1, 20)

	if head.Next.Value != 20 {
		t.Errorf("Expected head.Next.Value to be 20, got %v", head.Next.Value)
	}
	if head.Next.Next.Value != 30 {
		t.Errorf("Expected head.Next.Next.Value to be 30, got %v", head.Next.Next.Value)
	}
	if head.Next.Prev != head {
		t.Errorf("Expected head.Next.Prev to be head, got %v", head.Next.Prev)
	}
}

func TestFind(t *testing.T) {
	head := NewDoublyNode(10)
	head.InsertAtTail(20)
	head.InsertAtTail(30)

	if !head.Find(20) {
		t.Errorf("Expected to find value 20, but did not")
	}
	if head.Find(40) {
		t.Errorf("Expected not to find value 40, but did")
	}
}

func TestDelete(t *testing.T) {
	head := NewDoublyNode(10)
	head.InsertAtTail(20)
	head.InsertAtTail(30)

	head, deleted := head.Delete(20)
	if !deleted {
		t.Errorf("Expected to delete value 20, but did not")
	}
	if head.Find(20) {
		t.Errorf("Expected not to find value 20, but did")
	}

	head, deleted = head.Delete(10)
	if !deleted {
		t.Errorf("Expected to delete value 10, but did not")
	}
	if head != nil && head.Value == 10 {
		head.Print()
		t.Errorf("Expected head value to be updated after deletion, but it was not")
	}

	head, deleted = head.Delete(30)
	if !deleted {
		t.Errorf("Expected to delete value 30, but did not")
	}
	if head != nil {
		t.Errorf("Expected head to be nil after deleting last element, but it was not")
	}
}

func TestPrint(t *testing.T) {
	head := NewDoublyNode(10)
	head.InsertAtTail(20)
	head.InsertAtTail(30)

	if head != nil {
		head.Print() // Expected output: 10 20 30
	} else {
		t.Errorf("Expected head to be not nil, but it was nil")
	}
}
