package fundamentals

import "testing"

func TestList(t *testing.T) {
	node := NewList()
	if node.Count() != 0 {
		t.Error("node.Count() != 0")
	}

	if !node.Push(99) {
		t.Error("!node.Push(99)")

	}

	if !node.Push(1) {
		t.Error("!node.Push(1)")

	}

	if node.Count() != 2 {
		t.Error("node.Count() != 2")
	}

	if !node.Pop() {
		t.Error("!node.Pop()")
	}

	if node.Count() != 1 {
		t.Error("node.Count() != 1")
	}

	for _, val := range []int{1, 3, 22, 14, 55, 223} {
		if !node.Push(val) {
			t.Errorf("!node.Push(%d)\n", val)
		}
	}

	if node.Count() != 7 {
		t.Error("node.Count() != 7")
	}

	if !node.Shift() {
		t.Error("!node.Shift()")
	}

	if node.Count() != 6 {
		t.Error("node.Count() != 6")
	}

	if index, ok := node.FindIndex(22); ok {
		if val, ok := node.DeleteAt(index.(int)); !ok {
			t.Errorf("!node.Delete(%d) : %d \n", index, val)
		}
	} else {
		t.Error("index,ok := node.FindIndex(22); !ok")
	}

	for node.Count() > 0 {
		count := node.Count()
		getvalueat := count - 1
		if getvalueat < 0 {
			getvalueat = 0
		}
		if value, ok := node.Value(count - 1); ok {
			if index, ok := node.FindIndex(value); ok {
				if val, ok := node.DeleteAt(index.(int)); !ok {
					t.Errorf("!node.Delete(%d) : %d \n", index, val)
				}

				if node.Count() == count {
					t.Errorf("%d == %d\n", node.Count(), count)
				}
			} else {
				t.Errorf("node.FindIndex(%d); !ok \n", value)
			}
		} else {
			t.Errorf("node.Value(%d); !ok \n", node.Count())
		}
	}

	if node.Count() != 0 {
		t.Errorf("node.Count() != 0 \n %v", node.Iterate())
	}
}

func TestDeleteValue(t *testing.T) {
	node := NewList()
	if node.Count() != 0 {
		t.Error("node.Count() != 0")
	}
	for _, val := range []int{1, 3, 22, 14, 55, 223} {
		if !node.Push(val) {
			t.Errorf("!node.Push(%d)\n", val)
		}
	}

	if index, ok := node.FindIndex(55); ok {
		if val, ok := node.DeleteAt(index.(int)); !ok {
			t.Errorf("!node.Delete(%d) : %d \n", index, val)
		}
	} else {
		t.Error("index,ok := node.FindIndex(55); !ok")
	}

	if index, ok := node.FindIndex(14); ok {
		if val, ok := node.DeleteAt(index.(int)); !ok {
			t.Errorf("!node.Delete(%d) : %d \n", index, val)
		}
	} else {
		t.Error("index,ok := node.FindIndex(14); !ok")
	}

	if index, ok := node.FindIndex(1); ok {
		if val, ok := node.DeleteAt(index.(int)); !ok {
			t.Errorf("!node.Delete(%d) : %d \n", index, val)
		}
	} else {
		t.Error("index,ok := node.FindIndex(1); !ok")
	}

	if index, ok := node.FindIndex(223); ok {
		if val, ok := node.DeleteAt(index.(int)); !ok {
			t.Errorf("!node.Delete(%d) : %d \n", index, val)
		}
	} else {
		t.Error("index,ok := node.FindIndex(223); !ok")
	}
}
