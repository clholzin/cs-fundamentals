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

func MergeSort(data []int, left, right int) []int {

	if left >= right {
		return data
	}

	middle := (left + right) / 2

	fmt.Println(left, middle, right)

	fmt.Println("left", data[left:middle])
	fmt.Println("right", data[middle:right])

	data = MergeSort(data, left, middle)

	data = MergeSort(data, middle+1, right)

	return merge(data, left, middle, right)
}

// Merges two subarrays of arr[].
// First subarray is arr[l..m]
// Second subarray is arr[m+1..r]

func merge(data []int, l, m, r int) []int {

	n1 := m - l
	n2 := r - m

	left := data[l:m]
	right := data[m:r]

	i := 0
	j := 0
	k := l

	fmt.Println("--", n1, n2, data)
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
		k++
	}
	fmt.Println("-*", i, n1, j, n2, data)

	for i < n1 {
		data[k] = left[i]
		i++
		k++
	}

	for j < n2 {
		data[k] = right[j]
		j++
		k++
	}

	return data
}
