package fundamentals

import (
	"math"
)

/*
lava pools
[1,0,0,1,2,1,0,2]

        |     |
| _ _ | | | _ |
*/

func lavaVolumeSum(arr []int) int {
	var sum int
	if len(arr) == 0 {
		return sum
	}
	for i := 1; i < len(arr)-1; i++ {
		leftIdx, rightIdx := i, i

		for leftIdx > 0 {
			if arr[i] < arr[leftIdx] {
				break
			}
			leftIdx--
		}

		for rightIdx < len(arr) {
			if arr[i] < arr[rightIdx] || rightIdx == (len(arr)-1) {
				break
			}
			rightIdx++
		}

		minSideLevel := minSide(arr[leftIdx], arr[rightIdx])
		sum += int(math.Abs(float64(minSideLevel) - float64(arr[i])))
	}
	return sum
}

func minSide(a, b int) int {
	if a < b {
		return a
	}
	return b
}
