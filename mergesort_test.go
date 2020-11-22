package fundamentals

import "testing"

func TestMergeSort(t *testing.T) {

	data := []int{38, 27, 43, 3, 9, 82, 10}

	data = MergeSort(data)

	for i, val := range []int{3, 9, 10, 27, 38, 43, 82} {
		if data[i] != val {
			t.Errorf("failed to equal %d = %d", data[i], val)
		}
	}
	t.Log(data)
}
