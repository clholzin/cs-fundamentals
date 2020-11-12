package fundamentals

import (
	"fmt"
	"testing"
)

func TestSearchTree(t *testing.T) {

	tree := NewTree(1)
	tree.Left = NewTree(0)
	tree.Left.Parent = tree
	tree.Right = NewTree(2)
	tree.Right.Parent = tree

	if tree.Right.Parent != tree {
		t.Error("failed to find parent")
	}

	if tree.Search(1) != tree {
		t.Error("failed to find root")
	}

	if tree.Search(2) != tree.Right {
		t.Error("failed to find 2")
	}

	if tree.Search(0) != tree.Left {
		t.Error("failed to find 0")
	}
}

func TestInsertTree(t *testing.T) {

	tree := NewTree(1)

	if tree.Left != nil {
		t.Error("tree.Left should be nil")
	}

	tree.Insert(2, tree)
	tree.Insert(0, tree)

	if tree.Search(2) != tree.Right {
		t.Error("failed to find 2")
	}

	if tree.Search(0) != tree.Left {
		t.Error("failed to find 0")
	}
}

func TestTraverseTree(t *testing.T) {

	tree := NewTree(5)

	if tree.Left != nil {
		t.Error("tree.Left should be nil")
	}

	//	          5
	//				2       7
	//				  4   6   8
	//				3

	add := []int{2, 4, 3, 7, 6, 8}

	for _, val := range add {
		child := tree.Insert(val, tree)
		if child.Data != val {
			t.Errorf("failed to insert node %d\n", val)
		}
	}

	tree.Traverse()
	//time.Sleep(800 * time.Millisecond)
	fmt.Printf("%+v\n", tree)
}
