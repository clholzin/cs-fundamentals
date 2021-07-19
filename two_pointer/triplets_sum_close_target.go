package fundamentals

import (
	"math"
	"sort"
)

func searchTripletsCloseToTarget(arr []int, targetSum int) int {
	sort.Ints(arr)
	smallestDiff := math.MaxInt32
	for i := 0; i < len(arr)-2; i++ {
		left := i + 1
		right := len(arr) - 1

		for left < right {
			targetDiff := targetSum - arr[i] - arr[left] - arr[right]
			if targetDiff == 0 {
				return targetSum - targetDiff
			}

			if math.Abs(float64(targetDiff)) < math.Abs(float64(smallestDiff)) || targetDiff == smallestDiff && targetDiff > smallestDiff {
				smallestDiff = targetDiff
			}

			if targetDiff > 0 {
				left++
			} else {
				right--
			}
		}
	}
	return targetSum - smallestDiff
}
