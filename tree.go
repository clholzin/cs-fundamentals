package fundamentals

import "fmt"

type Treer interface {
	Search(int) *Tree
	Insert(int, *Tree) *Tree
	Remove(int) bool
	Traverse()
}

// Tree Node
type Tree struct {
	Data   int
	Parent *Tree
	Left   *Tree
	Right  *Tree
}

func NewTree(data int) Treer {
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
	var node *Tree
	if t == nil {
		node = &Tree{Data: val}
		node.Parent = parent
		t = node
		if val < parent.Data {
			parent.Left = t
		} else {
			parent.Right = t
		}
		return t
	}

	if val < t.Data {
		// fmt.Println(val, "left")
		return t.Left.Insert(val, t)
	}
	// fmt.Println(val, "right")
	return t.Right.Insert(val, t)
}

func (t *Tree) Remove(val int) bool {
	found := t.Search(val)
	if found != nil {
		found = nil
		return true
	}
	return false
}

func (t *Tree) process() {
	fmt.Println("-", t.Data)
}

func (t *Tree) Traverse() {
	if t != nil {
		t.Left.Traverse()
		t.Right.Traverse()
		t.process()
	}
}
