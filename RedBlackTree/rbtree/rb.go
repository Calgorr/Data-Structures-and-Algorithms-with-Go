package rbtree

type RB struct {
	leftChild, rightChild, parent, root *RB
	color, value                        int //0 For Black and 1 for Red
}

func (t *RB) LeftRotate(x *RB) {
	y := x.rightChild
	x.rightChild = y.leftChild
	if y.leftChild != nil {
		y.leftChild.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.leftChild {
		x.parent.leftChild = y
	} else {
		x.parent.rightChild = y
	}
	y.leftChild = x
	x.parent = y

}

func (t *RB) RightRotate(x *RB) {
	y := x.leftChild
	x.leftChild = y.rightChild
	if y.leftChild != nil {
		y.leftChild.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x.parent.leftChild == x {
		x.parent.leftChild = y
	} else {
		x.parent.leftChild = y
	}
	y.rightChild = x
	x.parent = y

}

func (t *RB) Max() int {
	for t.rightChild != nil {
		t = t.rightChild
	}
	return t.value
}

func (t *RB) Min() int {
	for t.leftChild != nil {
		t = t.leftChild
	}
	return t.value
}
