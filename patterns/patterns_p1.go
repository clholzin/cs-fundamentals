package patterns

import (
	"math"
	"sort"
)

//https://play.golang.org/p/aRQmqMndUnQ
//https://play.golang.org/p/tgBOlozww6y

/*
Smallest Subarray with a given sum (easy)
problem: return least amount of numbers whose sum equeals
or greater than Sum
details:
* count min window size
cases:
implementation / tradoffs
* sliding window  - two pointer
 * O(1) space
 * O(n) time
*/
func minSizeSubArraySum(arr []int, sum int) int {
	var (
		minWindow = 0
		total     = 0
		left      = 0
		right     = 0
		size      = len(arr)
	)

	for ; right < size; right++ {
		total += arr[right]
		for total >= sum {
			minWindow = min(minWindow, (right - left))
			total -= arr[left]
			left++
		}
	}
	return minWindow
}

/*
Subarrays with Product Less than a Target (medium)
Problem: add sub arrays whose total when multiplied is less than target
Detail:
* single numbers are added as sub arrays
* each iteration push single, append previous
cases:
* 2,5,3,10 - target 30
  * [2], [5], [2, 5], [3], [5, 3], [10]
*/

func subarrayProductLessThanK(arr []int, target int) [][]int {
	var (
		result  = make([][]int, 0)
		product = 1
		left    = 0
	)

	for right := 0; right < len(arr); right++ {
		product *= arr[right]
		for product >= target && left < len(arr) {
			product /= arr[left]
			left++
		}
		curr := make([]int, 0)
		for i := right; i >= left; i-- {
			curr = append(curr, arr[i])
			result = append(result, curr)
		}
	}
	return result
}

/*
Problem: count in unsorted array 3 different numbers that are less than target
Details:
* sort numbers
* use i + j + k < target
* consider i j are 1 apart and k moves not equal them
cases:
* -1, 0, 2, 3 target 3
  * i = -1, j = 0, k = 2, k = 3
	(-1 + 0 + 2) = 1 (-1, 0, 3) = 2
	(0, 2, 3) = 5
	already seen -1 in this same pattern (0, 2, -1) = 1

*/

func tripletWithSmallerSum(arr []int, target int) int {
	counter := 0
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		counter += searchPair(arr, target-arr[i], i)
	}
	return counter
}

func searchPair(arr []int, target, start int) int {
	var (
		count = 0
		left  = start + 1
		right = len(arr) - 1
	)
	for left < right {
		if arr[left]+arr[right] < target {
			count += right - left
			left++
		} else {
			right--
		}
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func t(arr []int, target int) int {
	sort.Ints(arr)
	var (
		count = 0
	)
	for i := 0; i < len(arr); i++ {
		if arr[i] < target {
			count += tpair(arr, target-arr[i], i)
		}
	}
	return count
}

func tpair(arr []int, target, start int) int {
	var (
		left  = start + 1
		right = len(arr) - 1
		count = 0
	)

	for left < right {
		if arr[left]+arr[right] < target {
			count += right - left
			left++
		} else {
			right--
		}
	}
	return count
}

func toposort(schedule [][]int, courses int) (result []int) {

	var (
		parents = make(map[int][]int)
		degrees = make(map[int]int)
	)

	for i := 0; i < courses; i++ {
		parents[i] = make([]int, 0)
		degrees[i] = 0
	}

	for i := 0; i < len(schedule); i++ {
		parent, child := schedule[i][0], schedule[i][1]
		parents[parent] = append(parents[parent], child)
		degrees[child]++
	}

	queue := make([]int, 0)

	for node, val := range degrees {
		if val == 0 {
			queue = append(queue, node)
			break
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		children := parents[node]
		for child := range children {
			degrees[child]--
			if val, ok := degrees[child]; ok && val == 0 {
				queue = append(queue, val)
			}
		}
	}

	return result
}

/*
ShortestWindowSort

problem: find the length of the sub array which contains values out of order so that when ordered they sort the array

details
 * identify from beginning out of order value as low
 * identify from end out of order value as high
 * from low t0 high get Min and Max
 * from low decrement if larger than Min
 * from high increase if lower than Max

cases:
 *  [1, 0, -1, 3, 4] = [1, 0, -1] = 3
 *  [1, 3, 2, 0, -1, 7, 10] = [1, 3, 2, 0, -1] = 5
 *  [10,2,3,5,4,1,8] = 6
    low = 2, high = 1
		max = 10, min = 1
		low => index 0
		high => index 6
*/

func minWindowSort(arr []int) int {
	var (
		low       = 0
		high      = len(arr) - 1
		minSubVal = math.MaxInt32
		maxSubVal = math.MinInt32
	)
	for low < len(arr)-1 && arr[low] <= arr[low+1] {
		low++
	}
	for high > 1 && arr[high] >= arr[high-1] {
		high--
	}

	for k := low; k <= high; k++ {
		minSubVal = min(minSubVal, arr[k])
		maxSubVal = max(maxSubVal, arr[k])
	}

	for low > 0 && arr[low-1] > minSubVal {
		low--
	}

	for high < len(arr)-1 && arr[high+1] < maxSubVal {
		high++
	}

	return high - low + 1
}

func minWindow(arr []int) int {

	var (
		left   = 0
		right  = len(arr) - 1
		minVal = math.MaxInt
		maxVal = math.MinInt
	)

	for left < len(arr)-2 && arr[left] <= arr[left+1] {
		left++
	}

	for right > 1 && arr[right] >= arr[right-1] {
		right--
	}

	for i := left; i <= right; i++ {
		minVal = min(arr[i], minVal)
		maxVal = max(arr[i], maxVal)
	}

	for left > 0 && arr[left] > minVal {
		left--
	}

	for right < len(arr) && right < maxVal {
		right++
	}

	return right - left + 1

}
