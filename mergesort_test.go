package fundamentals

import "testing"

func TestMergeSort(t *testing.T) {

	data := []int{38, 27, 43, 3, 9} //, 82, 10}

	data = MergeSort(data)
	t.Log(data)
}
