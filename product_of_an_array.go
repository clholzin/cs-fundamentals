package fundamentals

func productExceptSelf(nums []int) []int {
	var result []int = make([]int, len(nums))

	result[0] = 1

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	var R int = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * R
		R *= nums[i]
	}

	return result
}
