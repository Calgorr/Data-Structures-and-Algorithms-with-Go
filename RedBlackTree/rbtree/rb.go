package rbtree

type RB struct {
	leftChild, rightChild, parent, root *RB
	color, value                        int //0 For Black and 1 for Red
}

func (t *RB) RB_Delete(z *RB) {
	y := z
	y_originalcolor := y.color
	var x *RB
	if z.leftChild == nil {
		x = z.rightChild
		t.Transplant(z, z.rightChild)
	} else if z.rightChild == nil {
		x = z.leftChild
		t.Transplant(z, z.leftChild)
	} else {
		y.value = z.rightChild.Min()
		y_originalcolor = y.color
		x = y.rightChild
		if y.parent == z {
			x.parent = z
		} else {
			t.Transplant(y, y.rightChild)
			y.rightChild = z.rightChild
			y.rightChild.parent = y
		}
		t.Transplant(z, y)
		y.leftChild = z.leftChild
		y.leftChild.parent = y
		y.color = z.color
	}
	if y_originalcolor == 0 {
		t.RB_DeleteFixUp(x)
	}
}

func (t *RB) RB_DeleteFixUp(x *RB) {

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

func (t *RB) Transplant(u, v *RB) {
	if u.parent == nil {
		t.root = v
	} else if u.parent.leftChild == u {
		u.parent.leftChild = v
	} else {
		u.parent.rightChild = v
	}
	v.parent = u.parent
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
