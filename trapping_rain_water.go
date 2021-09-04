package fundamentals

//https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	result := 0
	if len(height) < 3 {
		return result
	}
	size := len(height)

	leftMax, rightMax := make([]int, size), make([]int, size)
	leftMax[0] = height[0]
	rightMax[size-1] = height[size-1]

	for i := 1; i < size; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	for i := 0; i < size; i++ {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}

	return result
}

/*
func trap(height []int) int {
	if len(height) == 0 {
		return int(0)
	}
	var total int
	var countsteps int
	var lastpeakIndex int
	var foundgreaterequal bool
	var index int
	var decr int

	length := len(height) - 1

	for index < len(height) {
		val := height[index]
		lastpeak := height[lastpeakIndex] - decr

		if index > 0 && val < lastpeak {
			countsteps++
			foundgreaterequal = true
		} else if index > 0 && val >= lastpeak {
			var lowest int

			if lastpeak >= val {
				lowest = val
			} else {
				lowest = lastpeak
			}

			for i := 1; i <= countsteps; i++ {
				v := height[index-i]
				total += (lowest - v)
			}

			countsteps = 0
			lastpeak = val
			foundgreaterequal = false
		}

		if index > 0 && index == length && foundgreaterequal {
			decr++
			index = lastpeakIndex + 1
			countsteps = 0
			foundgreaterequal = false
			continue
		}

		if val >= lastpeak {
			lastpeakIndex = index
			decr = 0
		}

		index++
	}
	return total
}*/
