package fundamentals

import "testing"

func TestStack(t *testing.T) {

	stack := NewStack(4)
	if stack.IsEmpty() {
		t.Errorf("failed to create stack")
	}
	if val, _ := stack.Peak(); val != 4 {
		t.Errorf("expected 4")
	}
	stack.Push(223)

	if err := stack.Pop(); err != nil {
		t.Errorf("failed to pop")
	}
}
