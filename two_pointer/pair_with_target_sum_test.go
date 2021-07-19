package fundamentals

import "testing"

func TestPairTargetSum(t *testing.T) {

	result := pairWithTargetSum([]int{1, 2, 3, 4, 6}, 6)
	if len(result) == 2 {
		if result[0] != 1 && result[1] != 3 {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}
