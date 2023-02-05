package bst

import "fmt"

type node[V comparable] struct {
	value                         V
	leftChild, rightChild, parent *node[V]
}

func main() {

}

func (x *node[V]) InOrder_Tree_Walk() {
	if x != nil {
		x.leftChild.InOrder_Tree_Walk()
		fmt.Println(x.value)
		x.rightChild.InOrder_Tree_Walk()
	}
}

func (x *node[V]) PreOrder_Tree_Walk() {
	fmt.Println(x.value)
	x.leftChild.PreOrder_Tree_Walk()
	x.rightChild.PreOrder_Tree_Walk()
}

func (x *node[V]) PostOrder_Tree_Walk() {
	x.leftChild.PostOrder_Tree_Walk()
	x.rightChild.PostOrder_Tree_Walk()
	fmt.Println(x.value)
}
