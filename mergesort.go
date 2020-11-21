package fundamentals

import "fmt"

/*
Merge Sort is a Divide and Conquer algorithm

MergeSort(arr[], l,  r)
  ```If r > l
     1. Find the middle point to divide the array into two halves:
             middle m = (l+r)/2
     2. Call mergeSort for first half:
             Call mergeSort(arr, l, m)
     3. Call mergeSort for second half:
             Call mergeSort(arr, m+1, r)
     4. Merge the two halves sorted in step 2 and 3:
             Call merge(arr, l, m, r)
*/

func MergeSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	left := 0
	right := len(data)
	middle := (left + right) / 2

	fmt.Println(left, middle, right)

	leftdata := data[left:middle]
	rightdata := data[middle:right]

	fmt.Println("left", leftdata)
	fmt.Println("right", rightdata)

	return merge(MergeSort(leftdata), MergeSort(rightdata))
}

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]

func merge(left, right []int) (result []int) {

	n1 := len(left)
	n2 := len(right)
	result = make([]int, n1+n2)

	i := 0

	fmt.Println("--", n1, n2)

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

	fmt.Println(left, right, result)
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
