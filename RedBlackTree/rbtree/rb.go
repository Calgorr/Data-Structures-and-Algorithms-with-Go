package rbtree

type RB struct {
	leftChild, rightChild, parent, root *RB
	value                               int
	color                               Color
}

type Color int64

const (
	BLACK Color = 0
	RED   Color = 1
)

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
	if y_originalcolor == BLACK {
		t.RB_DeleteFixUp(x)
	}
}

func (t *RB) RB_DeleteFixUp(x *RB) {
	for x != t.root && x.color == BLACK {
		if x == x.parent.leftChild {
			w := x.parent.rightChild
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.LeftRotate(x.parent)
				w = x.parent.rightChild
			}
			if w.leftChild.color == BLACK && w.rightChild.color == BLACK {
				w.color = RED
				x = x.parent
			} else if w.rightChild.color == BLACK {
				w.leftChild.color = BLACK
				w.color = RED
				t.RightRotate(w)
				w = x.parent.rightChild
			} else {
				w.color = x.parent.color
				x.parent.color = BLACK
				w.rightChild.color = BLACK
				t.LeftRotate(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.leftChild
			if w.color == RED {
				w.color = BLACK
				x.parent.color = RED
				t.RightRotate(x.parent)
				w = x.parent.leftChild
			}
			if w.rightChild.color == BLACK && w.leftChild.color == BLACK {
				w.color = RED
				x = x.parent
			} else if w.leftChild.color == BLACK {
				w.rightChild.color = BLACK
				w.color = RED
				t.LeftRotate(w)
				w = x.parent.leftChild
			} else {
				w.color = x.parent.color
				x.parent.color = BLACK
				w.leftChild.color = BLACK
				t.RightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = BLACK
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
