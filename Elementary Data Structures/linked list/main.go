package main

import (
	"fmt"
	LinkedList "main/linkedlist"
)

func main() {
	linkedlist := LinkedList.NewLInkedList(1)
	linkedlist.Insert(2)
	linkedlist.Insert(3)
	linkedlist.Insert(4)
	for ; linkedlist.Next != nil; linkedlist = linkedlist.Next {
		fmt.Print(linkedlist.Value, " ")
	}
	fmt.Println()
	linkedlist.Delete(4)
	for ; linkedlist.Next != nil; linkedlist = linkedlist.Next {
		fmt.Print(linkedlist.Value, " ")
	}
}
