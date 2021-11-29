package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(dutchFlag([]int{1, 0, 2, 1, 0}))
}

func dutchFlag(nums []int) []int {
	start, high := 0, len(nums)-1
	for i := 0; i < len(nums)-1; {
		cur := nums[i]
		if cur == 0 {
			swap(nums, i, start)
			i++
			start++
		} else if cur == 1 {
			i++
		} else {
			swap(nums, i, high)
			high--
		}
	}

	return nums
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}

func longestNoRepleating(str string) int {
	windowStart, maxLength := 0, 0

	mem := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		curr := str[windowEnd]
		if _, ok := mem[curr]; ok {
			windowStart = max(windowStart, mem[curr]+1)
		}
		mem[curr] = windowEnd
		maxLength = max(maxLength, windowEnd-windowStart+1)
	}
	return maxLength
}

/*
func reverse(head *ListNode) *ListNode {
	tmp := head
	for tmp != nil {
		t := head.next
	}
}*/

func CapPerms(val string) []string {

	perms := make([]string, 0)
	if len(val) == 0 {
		return perms
	}

	perms = append(perms, val)

	for i := 0; i < len(val); i++ {
		queLen := len(perms)
		for q := 0; q < queLen; q++ {
			perm := []byte(perms[q])
			if perm[i] >= byte('A') {
				if perm[i] < byte('a') {
					char := perm[i]
					charStr := strings.ToLower(string(char))
					perm[i] = []byte(charStr)[0]
					perms[q] = string(perm)
				} else {
					char := perm[i]
					charStr := strings.ToUpper(string(char))
					perm[i] = []byte(charStr)[0]
					perms = append(perms, string(perm))
				}
			}
		}
	}
	return perms
}

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	missing := 0
	for _, num := range nums {
		missing = missing ^ num
	}
	return missing
}

func knapsack(weights, profits []int, capacity int) int {
	if capacity <= 0 || len(weights) == 0 || len(weights) != len(profits) {
		return 0
	}

	dp := make([]int, capacity+1)

	for i := 0; i < capacity; i++ {
		if weights[0] <= i {
			dp[i] = profits[0]
		}
	}

	for i := 1; i < len(profits); i++ {
		for c := capacity; c >= 0; c-- {

			p1, p2 := 0, 0
			if weights[i] <= c {
				p1 = profits[i] + dp[c-weights[i]]
			}

			p2 = dp[c]
			dp[c] = max(p1, p2)

		}

	}
	return dp[capacity]
}

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
