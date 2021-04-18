package fundamentals

import (
	"testing"
)

type testTable struct {
	Values []int
	Expect int
}

func TestBSPairs(t *testing.T) {

	tabletest := []testTable{
		{Values: []int{1, 1, 2, 2, 3, 4, 4, 5, 5}, Expect: 3},
		{Values: []int{1, 1, 2, 3, 3, 4, 4, 5, 5}, Expect: 2},
		{Values: []int{1, 1, 2, 2, 3, 3, 4, 5, 5}, Expect: 4},
		{Values: []int{1, 1, 2}, Expect: 2},
		{Values: []int{1, 1, 2, 2, 3}, Expect: 3},
	}

	for _, data := range tabletest {
		if c := bs(data.Values); c != data.Expect {
			t.Errorf("failed on %d %d", c, data.Expect)
		}
		//t.Log("finished")
	}
}
