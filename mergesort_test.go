package fundamentals

import "testing"

func TestMergeSort(t *testing.T) {

	data := []int{38, 27, 43, 3, 9} //, 82, 10}

	data = MergeSort(data)
	t.Log(data)
}

func TestStuff(t *testing.T) {
	data := []int{38, 27, 43, 3, 9} //, 82, 10}

	t.Log(data[4:5])
	t.Log(data[2:5])

	t.Log(1 + 0)
	t.Log(0 + 3/2)
	t.Log(1 - 1 - 0)
}
