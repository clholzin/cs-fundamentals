package fundamentals

import (
	"testing"
)

func TestSubtract(t *testing.T) {

	expected := []int{-88, -73, -24, -18, 21, -4, -27, -44, -16, 76, 70, -78, 7, 61, -46, -30, -12, -17, 42, -18, -63, 71, 50, 35, 80, 99, 98, 10, 20, 46, 29, 85, 43, 13, 99, 12, 20, 71, 79, 1, 1, 57, 31, 78, 59, 16, 11, 7}

	input := []int{95, 84, 40, 77, 57, 35, 84, 45, 17, 3, 1, 98, 5, 38, 59, 73, 97, 46, 4, 38, 73, 27, 49, 45, 80, 99, 98, 10, 20, 46, 29, 85, 43, 13, 99, 12, 20, 71, 79, 1, 1, 57, 31, 78, 59, 16, 11, 7}

	header := &LinkedList{}
	for _, val := range input {
		header.Add(val)
	}

	subtract(header.Head)
	output := header.ToArray()

	if len(output) != len(expected) {
		t.Errorf("not equal - %v\n", output)
		t.FailNow()
	}
	for i := 0; i < len(expected); i++ {
		if expected[i] != output[i] {
			t.Errorf("failed to equal- %d - %d", expected[i], output[i])
			t.FailNow()
		}
	}
}
