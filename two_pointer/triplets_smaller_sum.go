package fundamentals

import "sort"

func searchTriplets(arr []int, target int) int {
	sort.Ints(arr)
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		count += searchPairAndCount(arr, target-arr[i], i)
	}
	return count
}

func searchPairAndCount(arr []int, targetSum, first int) int {

	count := 0
	left := first + 1
	right := len(arr) - 1

	for left < right {
		if arr[left]+arr[right] < targetSum {
			count += right - left
			left++
		} else {
			right--
		}
	}
	return count
}
