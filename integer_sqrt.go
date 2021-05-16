package fundamentals

// binary search
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	var left int = 2
	var right int = x / 2
	var pivot int
	for left <= right {
		pivot = left + (right-left)/2
		num := pivot * pivot

		if num > x {
			right = pivot - 1
		} else if num < x {
			left = pivot + 1
		} else {
			return pivot
		}

	}
	return right
}
