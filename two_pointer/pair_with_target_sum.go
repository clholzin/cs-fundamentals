package fundamentals

/*
Given an array of sorted numbers and a target sum, find a pair in the array whose sum is equal to the given target.

Write a function to return the indices of the two numbers (i.e. the pair) such that they add up to the given target.

Input: [1, 2, 3, 4, 6], target=6
Output: [1, 3]
Explanation: The numbers at index 1 and 3 add up to 6: 2+4=6
*/
// without added memory

func pairWithTargetSum(arr []int, targetSum int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			return []int{left, right}
		}
		if targetSum > currentSum {
			left++
		} else {
			right--
		}
	}

	return []int{-1, -1}
}

func pairWithTargetSumAddedMem(arr []int, targetSum int) []int {
	results := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if sum, ok := results[targetSum-arr[i]]; ok {
			return []int{sum, i}
		} else {
			results[targetSum-arr[i]] = i
		}
	}
	return []int{-1, -1}
}
