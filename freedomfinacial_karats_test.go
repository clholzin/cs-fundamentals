package fundamentals

import (
	"fmt"
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

func TestKarats(t *testing.T) {
	fmt.Println(hasCommonAncestor(parentChildPairs1, 3, 8))
}
