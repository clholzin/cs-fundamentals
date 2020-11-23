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

func TestMergeSort2(t *testing.T) {

	data := []int{22, 89, 4, 8, 50, 20, 28, 58, 6, 9, 67}

	data = MergeSort(data)

	for i, val := range []int{4, 6, 8, 9, 20, 22, 28, 50, 58, 67, 89} {
		if data[i] != val {
			t.Errorf("failed to equal %d = %d", data[i], val)
		}
	}
	t.Log(data)
}
