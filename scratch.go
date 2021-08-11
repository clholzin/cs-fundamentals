package fundamentals

/*
[1,2,3,4,5]
[[1,2,3,4,5],[1,2,3,4],[1,2,3],[1],[4],[5]]


[1]
[[2]]


*/

func pcMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := arr[:mid]
	rightArr := arr[mid:]

	return pcMerge(pcMergeSort(leftArr), pcMergeSort(rightArr))
}

func pcMerge(l1, l2 []int) []int {
	l1len := len(l1)
	l2len := len(l2)
	result := make([]int, l1len+l2len)
	i := 0

	for len(l1) > 0 && len(l2) > 0 {
		if l1[0] <= l2[0] {
			result[i] = l1[0]
			l1 = l1[1:]
		} else {
			result[i] = l2[0]
			l2 = l2[1:]
		}
		i++
	}

	for j := 0; j < len(l1); j++ {
		result[i] = l1[j]
		i++
	}
	for j := 0; j < len(l2); j++ {
		result[i] = l2[j]
		i++
	}
	return result
}

/*
func mergesort_practice(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	start := 0
	end := len(data)
	mid := (start + end) / 2

	left := data[start:mid]
	right := data[mid:end]
	return merge_practice(mergesort_practice(left), mergesort_practice(right))
}

func merge_practice(left, right []int) []int {

	lLen := len(left)
	rLen := len(right)
	res := make([]int, lLen+rLen)
	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			res[i] = left[0]
			left = left[1:]
		} else {
			res[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		res[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		res[i] = right[j]
		i++
	}

	return res

}
*/

/*
func mergesort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	left := 0

	right := len(right)
	middel := (left + right) / 2

	ldata := data[left:middle]
	rdata := data[middle:right]

	return merge(mergesort(ldata), mergesort(rdata))
}

func merge(left, right []int) []int {

	leftLen := len(left)
	rightLen := len(right)
	result := make([]int, leftLen+rightLen)
	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
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

}*/

/*

func MergeSort(data []int) []int {

	if len(data) == 0 {
		return data
	}

	left := 0
	right := len(data)
	middle := (left - right) / 2

	leftdata := data[left:middle]
	rightdata := data[middle:right]

	return merge(MergeSort(leftdata), MergeSort(rightData))

}

func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)
	result := make([]int, leftLen+rightLen)
	i := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
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



*/
