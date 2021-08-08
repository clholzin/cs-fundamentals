package fundamentals

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			result = twoSum(nums, i, result)
		}
	}
	return result
}

func twoSum(nums []int, i int, result [][]int) [][]int {
	seen := make(map[int]int)
	for j := i + 1; j < len(nums); j++ {
		complement := -nums[i] - nums[j]
		if _, ok := seen[complement]; ok {
			result = append(result, []int{nums[i], nums[j], complement})
			for j < (len(nums)-1) && nums[j] == nums[j+1] {
				j++
			}
		}
		seen[nums[j]] = j
	}
	return result
}
