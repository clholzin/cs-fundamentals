package fundamentals

func MergeSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	left := 0
	mid := len(data) / 2
	right := len(data)

	leftdata := data[start:mid]
	rightdata := data[mid:right]

	return merge(MergeSort(leftdata), MergeSort(rightdata))
}

func merge(left, right []int) []int {
	leftlen := len(left)
	rightlen := len(right)
	result := make([]int, leftlen+rightlen)

	i := 0

	for len(left) > 0 || len(right) > 0 {

		if left[i] < right[i] {
			result[i] = left[i]
			left = left[1:]
		} else {
			result[i] = right[i]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result

}
