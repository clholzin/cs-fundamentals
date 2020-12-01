package fundamentals

import "testing"

func TestBubble(t *testing.T) {
	data := []int{1, 6, 2, 5, 7, 4, 8}
	expected := []int{1, 2, 4, 5, 6, 7, 8}
	result := Bubble(data)
	for ind, val := range result {
		if val != expected[ind] {
			t.Logf("failed to match %d = %d\n", val, expected[ind])
			t.Fail()
		}
	}
}
