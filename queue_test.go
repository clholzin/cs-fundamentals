package fundamentals

import "testing"

func TestQueue(t *testing.T) {

	queue := NewQueue(NewQueueNode(4))
	if queue.IsEmpty() {
		t.Errorf("failed set data")
	}
	queue.Add(10)
	val, err := queue.Peak()
	if err != nil {
		t.Errorf("Failed see head")
	}
	if val != 4 {
		t.Errorf("val != 4")
	}
	val, err = queue.Remove()
	if err != nil {
		t.Errorf("Remove failed: %s", err)
	}
	if val != 4 {
		t.Errorf("val != 4")
	}
	val, err = queue.Remove()
	if err != nil {
		t.Errorf("Remove failed: %s", err)
	}
	if val != 10 {
		t.Errorf("val != 10")
	}

	if !queue.IsEmpty() {
		t.Errorf("should be empty")
	}
}
