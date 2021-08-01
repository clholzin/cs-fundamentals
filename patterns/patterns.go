package fundamentals

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Sliding Window Pattern
*/

/*
Maximum Sum Subarray of Size K (easy)
time: O(n)
space: O(1)
*/
func patternFindMaxSubArray(k int, arr []int) int {
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
Two Pointers
*/
/*
Pair With Target Sum
time:O(n)
space: O(1)
*/

func patternPairWithTargetSum(arr []int, targetSum int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			return []int{left, right}
		}
		if targetSum > currentSum {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

/*
time:O(n)
space: O(n)
*/
func patternPairWithTargetSumAlt(arr []int, targetSum int) []int {
	nums := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if val, ok := nums[targetSum-arr[i]]; ok {
			return []int{val, i}
		}
		nums[arr[i]] = i
	}
	return []int{-1, -1}
}

/*
Fast & Slow Pointers
*/

// List Node is info
type ListNode struct {
	Value int
	Next  *ListNode
}

/*
LinkedList cycle
time: O(n)
space: O(1)
*/

func patternHasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true //countCycleLength(slow)
		}
	}
	return false
}

func countCycleLength(slow *ListNode) int {
	current := slow.Next
	count := 1

	for current != slow {
		current = current.Next
		count++
	}
	return count
}

/*
Merge Intervals
*/

// Intervals
type Interval struct {
	Start int
	End   int
}

func patternMergeIntervals(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	mergedIntervals := make([]Interval, 0)
	start := intervals[0].Start
	end := intervals[0].End

	for i := 1; i < len(intervals); i++ {

		interval := intervals[i]

		if interval.Start <= end {

			end = max(interval.End, end)
		} else {

			mergedIntervals = append(mergedIntervals, Interval{Start: start, End: end})
			start = interval.Start
			end = interval.End
		}
	}

	return mergedIntervals
}

/*
Cycle Sort
time: O(n)
space: O(1)
*/
func patternCycleSort(nums []int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i += 1
		}
	}
	return nums
}

/*
Find all missing Numbers
*/
func patternfindMissingNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}
	missingNums := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			missingNums = append(missingNums, i+1)
		}
	}
	return missingNums
}

/*
Pattern In-place Reversal of a LinkedList
*/
type LinkNode struct {
	Value int
	Next  *LinkNode
}

func patternReverseLinkedList(head *LinkNode) *LinkNode {
	current := head
	var previous *LinkNode
	var next *LinkNode

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
time: O(n)
space: O(n)
*/

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func patternBreadthFirstSearch(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, currentNode.Value)

			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
		result = append(result, currentLevel)
	}
	return result
}

/*
Pattern: Tree Depth First Search
time: O(n)
space: O(n)
*/

func patternTreePathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Value == sum && root.Left == nil && root.Right == nil {
		return true
	}
	return patternTreePathSum(root.Left, sum-root.Value) || patternTreePathSum(root.Right, sum-root.Value)
}

/*
Pattern Subsets
*/

/*
Balanced Parentheses (hard)
time: 2N
space: N*2N
*/
type ParenthesesString struct {
	Str        string
	Opencount  int
	Closecount int
}

func patternBallencedParentheses(num int) []string {
	result := make([]string, 0)
	queue := make([]*ParenthesesString, 0)

	queue = append(queue, &ParenthesesString{"", 0, 0})

	for len(queue) > 0 {
		ps := queue[0]
		queue = queue[1:]
		if ps.Opencount == num && ps.Closecount == num {
			result = append(result, ps.Str)
		} else {
			if ps.Opencount < num {
				queue = append(queue, &ParenthesesString{ps.Str + "(", ps.Opencount + 1, ps.Closecount})
			}
			if ps.Opencount > ps.Closecount {
				queue = append(queue, &ParenthesesString{ps.Str + ")", ps.Opencount, ps.Closecount + 1})
			}
		}
	}
	return result
}

/*
String Permutations by changing case (medium)
*/
func patternStringPermutationsByChgCase(str string) []string {

	perms := make([]string, 0)
	perms = append(perms, str)

	for i := 0; i < len(str); i++ {
		permCount := len(perms)
		for j := 0; j < permCount; j++ {
			strVal := []byte(perms[j])
			if strVal[i] >= byte('A') && strVal[i] <= byte('Z') {
				strVal[i] = []byte(strings.ToLower(string(strVal[i])))[0]
				perms[j] = string(strVal)
			} else {
				strVal[i] = []byte(strings.ToUpper(string(strVal[i])))[0]
				perms = append(perms, string(strVal))
			}
		}
	}

	return perms
}

/*
Pattern: Modified Binary Search
*/
/*
Order-agnostic Binary Search (easy)
time: O(log(n))
space: O(1)

*/
func patternBSSearch(arr []int, key int) int {
	ascending := arr[0] < arr[len(arr)-1]
	return bsSearchRec(arr, key, 0, len(arr), ascending)
}

func bsSearchRec(arr []int, key, start, end int, ascending bool) int {
	n := len(arr)
	if n == 0 || start > end {
		return -1
	}

	middle := start + (end-start)/2

	if arr[middle] == key {
		return middle
	}

	if ascending {
		if arr[middle] > key {
			end = middle - 1
		} else {
			start = middle + 1
		}
	} else {
		if arr[middle] < key {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return bsSearchRec(arr, key, start, end, ascending)
}

/*
BS: Ceiling of a Number (medium)
time:
space:
*/

func patternSearchCeilingOfNumber(arr []int, key int) int {
	if key > arr[len(arr)-1] {
		return -1
	}

	start, end := 0, len(arr)-1

	for start <= end {
		mid := start + (end-start)/2
		if key < arr[mid] {
			end = mid - 1
		} else if key > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}

func patternSearchFloorOfNumber(arr []int, key int) int {
	if key < arr[0] {
		return -1
	}
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if key > arr[mid] {
			start = mid + 1
		} else if key < arr[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return end
}

/*
Bitwise
*/

/*
Single Number (easy)
In a non-empty array of integers, every number appears twice except for one, find that single number.
Time: O(n)
Space: O(1)
*/

func patternFindSingleNumber(arr []int) int {
	num := 0
	for _, val := range arr {
		num = num ^ val
	}
	return num
}

/*
Two Single Numbers (medium)
Time: O(n)
Space: O(1)
*/
func patternFindSingleNumbers(arr []int) []int {
	n1n2 := 0
	for _, val := range arr {
		n1n2 ^= val
		//fmt.Printf("%b, %d\n", n1n2, n1n2)
	}

	// get the rightmost bit that is 1
	rightmostSetBit := 1
	for (rightmostSetBit & n1n2) == 0 {
		rightmostSetBit = rightmostSetBit << 1
		//fmt.Printf("%b, %d\n", rightmostSetBit, rightmostSetBit)
	}

	n1, n2 := 0, 0
	for _, val := range arr {
		if (val & rightmostSetBit) != 0 { // the bit is set
			n1 ^= val
			//fmt.Printf("$%b, %d - %d\n", n1, n1, val)
		} else { // the bit is not set
			n2 ^= val
			//fmt.Printf("*%b, %d - %d\n", n2, n2, val)
		}
	}

	return []int{n1, n2}
}

/*
Complement of Base 10 Number (medium)
*/

func patternBitwiseComplement(num int) int {
	var bitCount float64 = 0
	n := num
	for n > 0 {
		bitCount++
		n = n >> 1
	}

	allBitsSet := int(math.Pow(2, bitCount)) - 1

	return num ^ allBitsSet
}

/*
Flip and Invert an image
Time: O(N)
Space: O(1)
*/
func patternFlipAndInvertImage(arr [][]int) [][]int {

	C := len(arr[0])
	for _, row := range arr {
		for j := 0; j < (C+1)/2; j++ {
			tmp := row[j] ^ 1
			row[j] = row[C-1-j] ^ 1
			row[C-1-j] = tmp
		}
	}

	return arr
}

/*
0r1 Knapsack (Dynamic Programming)
time: time complexity is exponential O(2^n)
space; O(n)
*/

func patternKnapsack(profits, weights []int, capacity int) int {
	return knapsackRec(profits, weights, capacity, 0)
}

func knapsackRec(profits, weights []int, capacity, currentIndex int) int {
	if capacity == 0 || currentIndex >= len(profits) {
		return 0
	}

	profit1 := 0
	if weights[currentIndex] <= capacity {
		profit1 = profits[currentIndex] + knapsackRec(profits, weights, capacity-weights[currentIndex], currentIndex+1)
	}

	profit2 := knapsackRec(profits, weights, capacity, currentIndex+1)

	return int(math.Max(float64(profit1), float64(profit2)))
}

/*
0/1 Knapsack (Dynamic Programming)
time: time complexity is exponential O(C)
space; O(n)
*/

func patternKnapsackFast(profits []int, weights []int, capacity int) int {

	if capacity <= 0 || len(profits) == 0 || len(profits) != len(weights) {
		return 0
	}

	dp := make([]int, capacity+1)

	for c := 0; c < capacity; c++ {
		if weights[0] <= c {
			dp[c] = profits[0]
		}
	}

	for i := 1; i < len(profits); i++ {
		for c := capacity; c > 0; c-- {
			profit1 := 0
			profit2 := 0

			if weights[i] <= c {
				profit1 = profits[i] + dp[c-weights[i]]
			}

			profit2 = dp[c]

			dp[c] = int(math.Max(float64(profit1), float64(profit2)))
		}

	}

	return dp[capacity]
}

/*
Topological Sort

time: O(V + E)
space: O(V + E)
*/

func patternTopologicalSort(vertices int, edges [][]int) []int {
	sortedOrder := make([]int, 0)

	if vertices <= 0 {
		return sortedOrder
	}

	inDegree := make(map[int]int)
	graph := make(map[int][]int)

	for i := 0; i < vertices; i++ {
		inDegree[i] = 0
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(edges); i++ {
		parent, child := edges[i][0], edges[i][1]
		graph[parent] = append(graph[parent], child)
		inDegree[child]++
	}

	queue := make([]int, 0)

	for child, count := range inDegree {
		if count == 0 {
			queue = append(queue, child)
		}
	}

	for len(queue) > 0 {
		vertex := queue[0]
		sortedOrder = append(sortedOrder, vertex)
		queue = queue[1:]
		children := graph[vertex]
		for _, child := range children {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if len(sortedOrder) != vertices {
		return make([]int, 0) // cycle found
	}

	return sortedOrder
}

/*
Task Scheduling (medium)
time: O(V + E)
space: O(V + E)

*/

func patternIsSchedulingPossible(tasks int, prerequisites [][]int) bool {
	sortedOrder := make([]int, 0)

	inDegree := make(map[int]int)
	graph := make(map[int][]int, tasks)

	for i := 0; i < tasks; i++ {
		inDegree[i] = 0
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		parent, child := prerequisites[i][0], prerequisites[i][1]
		inDegree[child]++
		graph[parent] = append(graph[parent], child)
	}

	queue := make([]int, 0)

	for child, count := range inDegree {
		if count == 0 {
			queue = append(queue, child)
		}
	}

	for len(queue) > 0 {
		vert := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, vert)
		children := graph[vert]
		for _, child := range children {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	return len(sortedOrder) == tasks
}

/*
All Tasks Scheduling Orders (hard)
time: O(V! + E)
space: O(V! + E)
*/

func patternAllTaskSchedulingOrders(tasks int, prerequisites [][]int) {
	sortedOrder := make([]int, 0)
	if tasks <= 0 {
		return
	}

	inDegree := make(map[int]int)
	graph := make(map[int][]int)

	for i := 0; i < tasks; i++ {
		inDegree[i] = 0
		graph[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		parent, child := prerequisites[i][0], prerequisites[i][1]
		inDegree[child]++
		graph[parent] = append(graph[parent], child)
	}

	queue := make([]int, 0)
	for child, count := range inDegree {
		if count == 0 {
			queue = append(queue, child)
		}
	}
	printAllTopoSorts(inDegree, graph, queue, sortedOrder)
}

func printAllTopoSorts(inDegree map[int]int, graph map[int][]int, queue []int, sortedOrder []int) {
	if len(queue) != 0 {
		for _, vert := range queue {
			sortedOrder = append(sortedOrder, vert)
			queueForNextCall := cloneQueue(queue)
			queueForNextCall = queueForNextCall[1:]
			children := graph[vert]
			for _, child := range children {
				inDegree[child]--
				if inDegree[child] == 0 {
					queueForNextCall = append(queueForNextCall, child)
				}
			}

			printAllTopoSorts(inDegree, graph, queueForNextCall, sortedOrder)

			sortedOrder = sortedOrder[:len(sortedOrder)-1]
			for _, child := range children {
				inDegree[child]++
			}

		}
	}

	if len(sortedOrder) == len(inDegree) {
		fmt.Println(sortedOrder)
	}
}

func cloneQueue(queue []int) []int {
	clone := make([]int, len(queue))
	copy(clone, queue)
	return clone
}

/*
Alien Dictionary (hard)
time: O(V + E)
space: O(V + E)
*/

func patternAlienDictionaryFindOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	inDegree := make(map[byte]int)
	graph := make(map[byte][]byte)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			inDegree[word[i]] = 0
			graph[word[i]] = make([]byte, 0)
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		for j := float64(0); j < math.Min(float64(len(w1)), float64(len(w2))); j++ {
			parent, child := w1[int32(j)], w2[int32(j)]
			if parent != child {
				graph[parent] = append(graph[parent], child)
				inDegree[child]++
				break
			}
		}
	}

	queue := make([]byte, 0)

	for child, count := range inDegree {
		if count == 0 {
			queue = append(queue, child)
		}
	}

	sortedOrder := strings.Builder{}

	for len(queue) > 0 {
		vert := queue[0]
		queue = queue[1:]
		if err := sortedOrder.WriteByte(vert); err != nil {
			fmt.Println("failed to WriteByte ", err)
			break
		}
		children := graph[vert]
		for _, child := range children {
			inDegree[child]--
			if inDegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if sortedOrder.Len() != len(inDegree) {
		return ""
	}
	return sortedOrder.String()
}

/*
Kth Smallest Number (hard)

Using the Max Heap

time: O(K * logK + (N - K) * logK) or O(N * logK)
space: O(K)
*/

func findKthSmallestNumberHeap(nums []int, k int) int {

	maxHeap := make(patternPriorityQueue, k)
	for i := 0; i < k; i++ {
		maxHeap[i] = &Item{
			value:    "",
			priority: nums[i],
			index:    i,
		}
	}

	heap.Init(&maxHeap)

	for i := k; i < len(nums); i++ {
		if nums[i] < maxHeap[0].priority {
			maxHeap.Pop()
			maxHeap.Push(&Item{
				priority: nums[i],
			})
		}
	}
	return maxHeap[0].priority
}

/*
Kth Smallest Number (hard)

Using the Median of Medians

guaranteed O(N) worst-case time. Please see the proof of its running time here and under “Selection-based pivoting”. The worst-case space complexity is O(N).
*/

func findKthSmallestNumberFast(nums []int, k int) int {
	return findKthSmallestNumberRec(nums, k, 0, len(nums)-1)
}

func findKthSmallestNumberRec(nums []int, k, start, end int) int {

	part := partition(nums, start, end)
	if part == k-1 {
		return nums[part]
	}

	if part > k-1 {
		return findKthSmallestNumberRec(nums, k, start, part-1)
	}

	return findKthSmallestNumberRec(nums, k, part+1, end)
}

func partition(nums []int, low, high int) int {
	if low == high {
		return low
	}

	median := medianOfMedians(nums, low, high)

	for i := low; i < high; i++ {
		if nums[i] == median {
			swap(nums, i, high)
			break
		}
	}

	pivot := nums[high]

	for i := low; i < high; i++ {
		if nums[i] < pivot {
			swap(nums, low+1, i)
		}
	}

	swap(nums, low, high)
	return low
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func medianOfMedians(nums []int, low, high int) int {

	n := high - low + 1
	if n < 5 {
		return nums[low]
	}

	numOfPartition := n / 5
	medians := make([]int, numOfPartition)
	for i := 0; i < numOfPartition; i++ {
		partitionStart := low + i*5
		sortPartition(nums, partitionStart, partitionStart+5)

	}
	return partition(medians, 0, numOfPartition-1)
}

func sortPartition(original []int, low, high int) {
	nums := original[low:high]
	sort.Ints(nums)
	index := 0
	for i := low; i < high; i++ {
		original[i] = nums[index]
		index++
	}
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A patternPriorityQueue implements heap.Interface and holds Items.
type patternPriorityQueue []*Item

func (pq patternPriorityQueue) Len() int { return len(pq) }

func (pq patternPriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq patternPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *patternPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *patternPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *patternPriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
