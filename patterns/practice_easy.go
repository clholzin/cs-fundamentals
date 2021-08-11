package patterns

import (
	"container/heap"
	"sort"
)

/*
Patten: Sliding Window
Maximum Sum Subarray of Size K (easy)

Given an array of positive numbers and a positive number ‘k,’ find the maximum sum of any contiguous subarray of size ‘k’.
*/

func practiceFindMaxSumSubArray(k int, arr []int) int {
	windowSum, maxSum, windowStart := 0, 0, 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {

		windowSum += arr[windowEnd]

		if windowEnd >= k-1 {
			maxSum = max(maxSum, windowSum)
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return maxSum
}

/*
Pattern: Two Pointer

Squaring a Sorted Array (easy)

Given a sorted array, create a new array containing squares of all the numbers of the input array in the sorted order.
*/

func practiceSortedArraySquares(arr []int) []int {

	n := len(arr)
	squares := make([]int, n)
	highestSquaredIdx := n - 1
	left, right := 0, len(arr)-1

	for left <= right {

		leftsquare := arr[left] * arr[left]
		rightsquare := arr[right] * arr[right]

		if leftsquare > rightsquare {
			squares[highestSquaredIdx] = leftsquare
			left++
		} else {
			squares[highestSquaredIdx] = rightsquare
			right--
		}

		highestSquaredIdx--
	}
	return squares
}

/*
Patterns: Fast & Slow

LinkedList Cycle (easy)
*/

type Node struct {
	Value string
	Next  *Node
}

func practiceHasCycle(head *Node) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}

	return false
}

func practiceCycleLength(slow *Node) int {
	current := slow
	cycleLength := 0
	for {
		current = current.Next
		cycleLength++
		if current == slow {
			break
		}
	}
	return cycleLength
}

/*
Pattern Merge Intervals

Merge Intervals (medium)

Given a list of intervals, merge all the overlapping intervals to produce a list that has only mutually exclusive intervals.
*/

type IntervalPc struct {
	Start int
	End   int
}

func practiceMergeIntervals(intervals []IntervalPc) []IntervalPc {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a].Start < intervals[b].Start
	})

	mergedIntervals := make([]IntervalPc, 0)
	start := intervals[0].Start
	end := intervals[0].End

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval.Start <= end {
			end = max(interval.End, end)
		} else {
			mergedIntervals = append(mergedIntervals, IntervalPc{start, end})
			start = interval.Start
			end = interval.End
		}
	}
	return mergedIntervals
}

/*
Pattern: Cycle Sort

Cyclic Sort (easy)

We are given an array containing ‘n’ objects. Each object, when created, was assigned a unique number from 1 to ‘n’ based on their creation sequence.
*/

func practiceCycleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := 0
	for left < len(arr) {
		j := arr[left] - 1
		if arr[left] != arr[j] {
			arr[left], arr[j] = arr[j], arr[left]
		} else {
			left++
		}
	}
	return arr
}

/*
Pattern: In-Place Reversal Of a LinkedList

Reverse a LinkedList (easy)

Given the head of a Singly LinkedList, reverse the LinkedList. Write a function to return the new head of the reversed LinkedList.
*/
type ListNodePc struct {
	Value int
	Next  *ListNodePc
}

func practiceReverseLinkedList(head *ListNodePc) *ListNodePc {
	current := head
	var previous *ListNodePc
	var next *ListNodePc

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}

/*
Pattern: Tree Breadth First Search

Binary Tree Level Order Traversal (easy)

Given a binary tree, populate an array to represent its level-by-level traversal. You should populate the values of all nodes of each level from left to right in separate sub-arrays.
*/

type TreeNodePc struct {
	Value int
	Left  *TreeNodePc
	Right *TreeNodePc
}

func practiceLevelOrderTraversal(root *TreeNodePc) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := make([]*TreeNodePc, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		currentLevel := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			queue = queue[1:]
			currentLevel[i] = node.Value
			if node.Left != nil {
				queue = append(queue, node.Left)
			} else if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}

/*
Pattern Tree Depth First Search

Binary Tree Path Sum (easy)

Given a binary tree and a number ‘S’, find if the tree has a path from root-to-leaf such that the sum of all the node values of that path equals ‘S’.
*/

func practiceTreePathSum(root *TreeNodePc, sum int) bool {
	if root == nil {
		return false
	}

	if root.Value == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return practiceTreePathSum(root.Left, sum-root.Value) || practiceTreePathSum(root.Right, sum-root.Value)
}

/*
Patter Two Heaps - see patterns Median of a stream
*/

/*
Pattern Subsets
SubSets (easy)

Given a set with distinct elements, find all of its distinct subsets.

Time: N * 2N
Space N * 2N
*/

func practiceFindSubsets(nums []int) [][]int {
	subsets := make([][]int, 0)

	subsets = append(subsets, []int{}) // start with empty subset
	for _, currentNum := range nums {
		n := len(subsets)
		for i := 0; i < n; i++ {
			set := []int{}
			set = append(set, subsets[i]...)
			set = append(set, currentNum)
			subsets = append(subsets, set)
		}
	}
	return subsets
}

/*
Pattern Binary Search

Order-agnostic Binary Search (easy)

Given a sorted array of numbers, find if a given number ‘key’ is present in the array. Though we know that the array is sorted, we don’t know if it’s sorted in ascending or descending order. You should assume that the array can have duplicates.
*/

func practiceBinarySearch(arr []int, key int) int {
	start, end := 0, len(arr)-1
	isAscending := arr[start] < arr[end]

	for start <= end {
		mid := start + (end-start)/2
		if key == arr[mid] {
			return mid
		}

		if isAscending {
			if key < arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if key > arr[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return -1
}

/*
Pattern: Bitwise XOR

Single Number (easy)

In a non-empty array of integers, every number appears twice except for one, find that single number.
*/

func practiceFindSingleNumber(arr []int) int {
	val := 0
	for _, num := range arr {
		val = val ^ num
	}
	return val
}

/*
Pattern: Top 'K' Elements

Top 'K' Numbers (easy)

Given an unsorted array of numbers, find the ‘K’ largest numbers in it.
*/

type MinHeap []int

func (min MinHeap) Len() int           { return len(min) }
func (min MinHeap) Less(i, j int) bool { return i < j }
func (min MinHeap) Swap(i, j int)      { min[i], min[j] = min[j], min[i] }
func (min *MinHeap) Pop() interface{} {
	old := *min
	n := len(old)
	x := old[n-1]
	*min = old[:n-1]
	return x
}
func (min *MinHeap) Push(val interface{}) {
	*min = append(*min, val.(int))
}
func (min *MinHeap) Peak() interface{} {
	temp := *min
	top := temp[len(temp)-1]
	return top
}

func practiceKLargestNumbers(nums []int, k int) []int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for i := 0; i < k; i++ {
		heap.Push(minHeap, nums[i])
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > minHeap.Peak().(int) {
			heap.Pop(minHeap)
			heap.Push(minHeap, nums[i])
		}
	}
	return *minHeap
}

/*
Pattern: K-way merge

Merge K Sorted Lists (medium)

Given an array of ‘K’ sorted LinkedLists, merge them into one sorted list.
*/
type MinHeapNode []*ListNodePc

func (min MinHeapNode) Len() int           { return len(min) }
func (min MinHeapNode) Less(i, j int) bool { return min[i].Value < min[j].Value }
func (min MinHeapNode) Swap(i, j int)      { min[i], min[j] = min[j], min[i] }
func (min *MinHeapNode) Pop() interface{} {
	old := *min
	n := len(old)
	x := old[n-1]
	*min = old[:n-1]
	return x
}
func (min *MinHeapNode) Push(val interface{}) {
	*min = append(*min, val.(*ListNodePc))
}
func (min *MinHeapNode) Peak() interface{} {
	temp := *min
	top := temp[len(temp)-1]
	return top
}
func practiceMergeKSortedLists(lists []*ListNodePc) *ListNodePc {
	minHeap := &MinHeapNode{}
	heap.Init(minHeap)

	for _, root := range lists {
		if root != nil {
			heap.Push(minHeap, root)
		}
	}

	var resultHead *ListNodePc
	var resultTail *ListNodePc
	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*ListNodePc)
		if resultHead == nil {
			resultHead = node
			resultTail = node
		} else {
			resultTail.Next = node
			resultTail = node.Next
		}
		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
	}
	return resultHead
}

/*
Pattern : 0/1 Knapsack (Dynamic Programming)

0/1 Knapsack (medium)

Given the weights and profits of ‘N’ items, we are asked to put these items in a knapsack with a capacity ‘C.’ The goal is to get the maximum profit out of the knapsack items. Each item can only be selected once, as we don’t have multiple quantities of any item.
*/

func practiceKnapSack(profits, weights []int, capacity int) int {
	if capacity <= 0 || len(profits) == 0 || len(profits) != len(weights) {
		return -1
	}

	dp := make([]int, capacity+1)

	for c := 0; c < capacity; c++ {
		if weights[0] <= c {
			dp[c] = profits[0]
		}
	}

	for i := 1; i < len(profits); i++ {
		for c := capacity; c > 0; c-- {
			profits1, profits2 := 0, 0
			if weights[i] <= c {
				profits1 = profits[i] + dp[c-weights[i]]
			}
			profits2 = dp[c]
			dp[c] = max(profits1, profits2)
		}
	}
	return dp[capacity]
}

/*
Pattern: Topological Sort (Graph)

Topological Sort (medium)

Given a directed graph, find the topological ordering of its vertices
Time: V+E
Space: V+E
*/

func practiceTopologicalSort(verticies int, edges [][]int) []int {
	sorted := make([]int, 0)
	if verticies <= 0 {
		return sorted
	}

	degree := make(map[int]int)
	parents := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		parents[i] = make([]int, len(edges[i]))
		degree[i] = 0
	}

	for i := 0; i < len(edges); i++ {
		parent, child := edges[i][0], edges[i][1]
		parents[parent] = append(parents[parent], child)
		degree[child]++
	}

	queue := make([]int, 0)

	for child, count := range degree {
		if count == 0 {
			queue = append(queue, child)
		}
	}

	for len(queue) > 0 {
		vert := queue[0]
		queue = queue[1:]
		sorted = append(sorted, vert)
		children := parents[vert]
		for i := 0; i < len(children); i++ {
			child := children[i]
			degree[child]--
			if degree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	return sorted
}
