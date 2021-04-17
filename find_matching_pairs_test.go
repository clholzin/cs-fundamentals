package fundamentals

import (
	"testing"
)

func TestBSPairs(t *testing.T) {

	vals := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}

	if c := bs(vals); c != 3 {
		t.Errorf("failed on %d %d", c, val)
	}

}
