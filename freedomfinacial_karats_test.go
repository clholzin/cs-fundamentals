package fundamentals

import (
	"testing"
)

var (
	parentChildPairs1 [][]int = [][]int{
		[]int{1, 3},
		[]int{2, 3},
		[]int{3, 6},
		[]int{5, 6},
		[]int{5, 7},
		[]int{4, 5},
		[]int{4, 8},
		[]int{4, 9},
		[]int{9, 11},
		[]int{14, 4},
		[]int{13, 12},
		[]int{12, 9},
		[]int{15, 13},
	}

	parentChildPairs2 [][]int = [][]int{
		[]int{1, 3},
		[]int{11, 10},
		[]int{11, 12},
		[]int{2, 3},
		[]int{10, 2},
		[]int{10, 5},
		[]int{3, 4},
		[]int{5, 6},
		[]int{5, 7},
		[]int{7, 8},
	}
)

type data struct {
	Pair1  int
	Pair2  int
	Expect bool
}

func TestKarats(t *testing.T) {

	tableTest := []data{
		data{Pair1: 3, Pair2: 8, Expect: false},
		data{Pair1: 5, Pair2: 8, Expect: true},
		data{Pair1: 6, Pair2: 8, Expect: true},
		data{Pair1: 6, Pair2: 9, Expect: true},
		data{Pair1: 1, Pair2: 3, Expect: false},
		data{Pair1: 3, Pair2: 1, Expect: false},
		data{Pair1: 7, Pair2: 11, Expect: true},
		data{Pair1: 6, Pair2: 5, Expect: true},
		data{Pair1: 5, Pair2: 6, Expect: true},
	}
	for index, d := range tableTest {
		if hasCommonAncestor(parentChildPairs1, d.Pair1, d.Pair2) != d.Expect {
			t.Errorf("failed at %d", index)
		}
	}
}
