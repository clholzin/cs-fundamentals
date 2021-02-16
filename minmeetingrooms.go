package fundamentals

// time complexity: nlog(n) + nm
// space: linear / mergesortdata + meetingtimes
func minMeetingRooms(intervals [][]int) int {
	var meetingtimes []int
	if len(intervals) == 0 {
		return 0
	}
	intervals = MergeSorter(intervals)
	meetingtimes = append(meetingtimes, intervals[0][1])
	intervals = intervals[1:]
	for _, v := range intervals {
		var allocated bool
		for i, ends := range meetingtimes {
			if ends <= v[0] {
				meetingtimes[i] = v[1]
				allocated = true
				break
			}
		}
		if !allocated {
			meetingtimes = append(meetingtimes, v[1])
		}
	}
	return len(meetingtimes)
}

func MergeSorter(data [][]int) [][]int {

	if len(data) <= 1 {
		return data
	}

	left := 0
	right := len(data)
	middle := (left + right) / 2

	leftdata := data[left:middle]
	rightdata := data[middle:right]

	return mergesortby(MergeSort(leftdata), MergeSort(rightdata))
}

func mergesortby(left, right [][]int) (result [][]int) {

	n1 := len(left)
	n2 := len(right)
	result = make([][]int, n1+n2)

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0][0] <= right[0][1] && left[0][0] <= right[0][0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
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
