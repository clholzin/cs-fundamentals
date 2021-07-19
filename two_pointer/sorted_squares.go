package fundamentals

func sortedSquares(nums []int) []int {
	left := 1
	right := len(nums) - 1
	nextHighest := right
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		leftsqr := nums[left-1] * nums[left-1]
		rightsqr := nums[right] * nums[right]
		if leftsqr > rightsqr {
			result[nextHighest] = leftsqr
			left++
		} else {
			result[nextHighest] = rightsqr
			right--
		}
		nextHighest--
	}

	return result
}
