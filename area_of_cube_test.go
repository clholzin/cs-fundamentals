package fundamentals

import (
	"fmt"
	"testing"
)

func TestAreaOfCube(t *testing.T) {

	cube := [][]int{
		{1, 2, 3, 2, 1, 3}, // 12 topedge,    6   12
		{2, 1, 1, 1, 2, 1}, // 12             3 8
		{2, 2, 1, 1, 2, 1}, // 12             3 1
		{3, 1, 1, 2, 3, 2}, // 12 bottomedge, 5 5  12
		// 8 leftedge  // 7 rightedge
	}

	fmt.Println(areaOfCube(cube))
}

func TestAreaOfCubeTwo(t *testing.T) {

	cube := [][]int{
		{1, 2}, // 3 1     4
		{2, 1}, // 3 1 1 1 4
		// 3 //3
	}

	fmt.Println(areaOfCube(cube))
}
