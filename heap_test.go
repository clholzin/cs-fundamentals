package fundamentals

import "testing"

func TestMinIntHeap(t *testing.T) {

	heap := NewMinIntHeap()
	heap.Add(24)
	heap.Add(32)
	heap.Add(34)
	heap.Add(10)
	heap.Add(12)
	heap.Add(14)
	heap.Add(22)

	if heap.Size() != 7 {
		t.Errorf("size is not 7: %v", heap.Items())
	}
	val, err := heap.Poll()
	if err != nil {
		t.Errorf("peak failed %d - %s", val, err)
	}
	//t.Logf("first poll %d", val)
	val, err = heap.Peak()
	if err == nil {
		if val != 12 {
			t.Errorf("peak value expected 14. was: %d", val)
		}
	} else {
		t.Errorf("peak failed")
	}

	var pre int
	for !heap.IsEmpty() {
		if val, err := heap.Poll(); err == nil {
			if pre > val {
				t.Errorf("%d > %d", pre, val)
				break
			}
			t.Logf("%d", val)
			pre = val
		}

	}

}
