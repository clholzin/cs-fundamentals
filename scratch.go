package fundamentals

/*
[1,2,3,4,5]
[[1,2,3,4,5],[1,2,3,4],[1,2,3],[1],[4],[5]]


[1]
[[2]]

*/

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

/*
func mergesort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	left := 0
<<<<<<< HEAD
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
=======
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
>>>>>>> f8bebed0d229a3039502176800305e70769a55c1
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
<<<<<<< HEAD
	return result

}
=======

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
