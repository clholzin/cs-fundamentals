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
	if val, _ := stack.Peak(); val != 223 {
		t.Errorf("expected 223: got %d", val)
	}
	if err := stack.Pop(); err != nil {
		t.Errorf("failed to pop")
	}
	if err := stack.Pop(); err != nil {
		t.Errorf("failed to pop")
	}

	if _, err := stack.Peak(); err == nil {
		t.Errorf("expected peak error")
	}

	if err := stack.Pop(); err == nil {
		t.Errorf("expected pop error")
	}

	if !stack.IsEmpty() {
		t.Errorf("!stack.IsEmpty()")
	}
}
