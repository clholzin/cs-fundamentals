package fundamentals

import "testing"

func TestMinIntHeap(t *testing.T) {

	heap := NewMinIntHeap()
	heap.Add(10)
	heap.Add(12)
	heap.Add(14)
	heap.Add(22)
	heap.Add(24)
	heap.Add(32)
	heap.Add(34)

	if heap.Size != 7 {
		t.Errorf("size is not 7: %v", heap.Values())
	}

	heap.Poll()

	if val, err := heap.Peak(); err == nil {
		if val != 14 {
			t.Errorf("peak value expected 14. was: %d", val)
		}
	} else {
		t.Errorf("peak failed")
	}

}
