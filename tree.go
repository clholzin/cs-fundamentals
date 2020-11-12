package fundamentals

import "fmt"

// Tree Node
type Tree struct {
	Data   int
	Parent *Tree
	Left   *Tree
	Right  *Tree
}

func NewTree(data int) *Tree {
	return &Tree{Data: data}
}

func (t *Tree) Search(val int) *Tree {

	if t == nil || val == t.Data {
		return t
	}

	if val < t.Data {
		return t.Left.Search(val)
	}
	return t.Right.Search(val)
}

func (t *Tree) Insert(val int, parent *Tree) *Tree {
	var p *Tree
	if t == nil {
		p = NewTree(val)
		p.Parent = parent
		t = p
		if val < parent.Data {
			parent.Left = t
		} else {
			parent.Right = t
		}
		return t
	}

	if val < t.Data {
		fmt.Println(val, "left")
		return t.Left.Insert(val, t)
	}
	fmt.Println(val, "right")
	return t.Right.Insert(val, t)
}

func (t *Tree) Remove(val int, parent *Tree) *Tree {
	found := t.Search(val)
	found = nil
	return found
}

func (t *Tree) Process() {
	fmt.Println("-", t.Data)
}

func (t *Tree) Traverse() {
	if t != nil {
		t.Left.Traverse()
		t.Right.Traverse()
		t.Process()
	}
}
