package bst

import (
	"fmt"
)

type node struct {
	value                         int
	leftChild, rightChild, parent *node
}

func (x *node) InOrder_Tree_Walk() {
	if x != nil {
		x.leftChild.InOrder_Tree_Walk()
		fmt.Println(x.value)
		x.rightChild.InOrder_Tree_Walk()
	}
}

func (x *node) PreOrder_Tree_Walk() {
	fmt.Println(x.value)
	x.leftChild.PreOrder_Tree_Walk()
	x.rightChild.PreOrder_Tree_Walk()
}

func (x *node) PostOrder_Tree_Walk() {
	x.leftChild.PostOrder_Tree_Walk()
	x.rightChild.PostOrder_Tree_Walk()
	fmt.Println(x.value)
}

func (x *node) RecursiveSearch(k int) *node {
	if x == nil || x.value == k {
		return x
	} else if x.value < k {
		x.rightChild.RecursiveSearch(k)
	}
	x.leftChild.RecursiveSearch(k)
	return nil
}

func (x *node) IterativeSearch(k int) *node {
	for x != nil || x.value != k {
		if x.value > k {
			x = x.leftChild
		} else {
			x = x.rightChild
		}
	}
	return x
}

func (x *node) Minumum() *node {
	for x != nil {
		x = x.leftChild
	}
	return x
}

func (x *node) Maximum() *node {
	for x != nil {
		x = x.rightChild
	}
	return x
}

func (x *node) Tree_Successor() *node {
	if x.rightChild != nil {
		return x.Minumum()
	}
	y := x.parent
	for x == y.rightChild || y != nil {
		x = y
		y = x.parent
	}
	return y
}
func Tree_Insert(root, z *node) {
	if root == nil {
		root = z
		return
	}
	x, y := root, &node{}
	for x != nil {
		y = x
		if z.value > x.value {
			x = x.rightChild
		} else {
			x = x.leftChild
		}
	}
	z.parent = y
	if z.value > y.value {
		y.rightChild = z
	} else {
		y.leftChild = z
	}

}

func Transplant(root, u, v *node) {
	if u.parent == nil {
		root = v
		return
	}
	if u == u.parent.leftChild {
		u.parent.leftChild = v
	} else {
		u.parent.rightChild = v
	}
	if v != nil {
		v.parent = u.parent
	}

}

func Deletion(root, v *node) {
	if v.leftChild == nil {
		Transplant(root, v, v.rightChild)
	} else if v.leftChild == nil {
		Transplant(root, v, v.rightChild)
	} else {

	}
}
