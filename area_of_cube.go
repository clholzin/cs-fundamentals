package fundamentals

import "math"

func areaOfCube(arr [][]int) int {

	sum := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			// top edge, left edge, right edge, bottom edge
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				sum += arr[i][j]
			}
			sum += 2 // bottom and current

			if j > 0 {
				// current left edge diff
				sum += int(math.Abs(float64(arr[i][j-1]) - float64(arr[i][j])))
			}

			if i > 0 {
				// current top edge diff
				sum += int(math.Abs(float64(arr[i-1][j]) - float64(arr[i][j])))
			}
		}
	}

	return sum
}
