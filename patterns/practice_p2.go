package patterns

import (
	"container/heap"
	"fmt"
)

/*
https://play.golang.org/p/WZ-Bqy46dYX

Sum of Elements (medium)
Given an array, find the sum of all numbers between the K1’th and K2’th
smallest elements of that array.

Input: [1, 3, 12, 5, 15, 11], and K1=3, K2=6
Output: 23

Problem: sum smallest numbers between k1 and k2 positions
Details:
* capture smallest values
* Min Heap structure
* Sort and parition array then sum

Cases
[1,3,5,11,12,15] k1 = 3, k2 = 6
* between 5 and 15 are 11 and 12, = 23
* sum all from k2 - k1 -1 = 2

*/

func sumOfElements2(arr []int, k1, k2 int) int {

	result := 0
	minheap := &MinInt{}
	heap.Init(minheap)

	for i := 0; i < len(arr); i++ {
		heap.Push(minheap, arr[i])
	}

	// remove all until k1
	for i := 0; i < k1; i++ {
		heap.Pop(minheap)
	}
	// sum all until k2  or k2-k1-1
	for i := 0; i < (k2 - k1 - 1); i++ {
		tmp := heap.Pop(minheap)
		result += tmp.(int)
	}
	return result

}

type MinInt []int

func (m MinInt) Len() int           { return len(m) }
func (m MinInt) Less(i, j int) bool { return m[i] < m[j] }
func (m MinInt) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MinInt) Push(val interface{}) {
	tmp := *m
	tmp = append(tmp, val.(int))
	*m = tmp
}
func (m *MinInt) Pop() interface{} {
	tmp := *m
	length := len(tmp)
	removed := tmp[length-1]
	*m = tmp[0 : length-1]
	return removed

}

/* Max Heap of between K elements
https://play.golang.org/p/nZ4Puxhox-y

func sumOfElements(arr []int, k1, k2 int) int {

	result := 0
	maxheap := &MaxInt{}
	heap.Init(maxheap)

	for i := 0; i < len(arr); i++ {
		if i < k2-1 {
			heap.Push(maxheap, arr[i])
		} else if arr[i] < maxheap.Peak().(int) {
			heap.Pop(maxheap)
			heap.Push(maxheap, arr[i])
		}
	}

	for i:=0;i<(k2-k1-1);i++ {
		tmp := heap.Pop(maxheap)
		result += tmp.(int)
	}
	return result

}

type MaxInt []int

func (m MaxInt) Len() int           { return len(m) }
func (m MaxInt) Less(i, j int) bool { return m[i] > m[j] }
func (m MaxInt) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *MaxInt) Push(val interface{}) {
	tmp := *m
	tmp = append(tmp, val.(int))
	*m = tmp
}
func (m *MaxInt) Pop() interface{} {
	tmp := *m
	length := len(tmp)
	removed := tmp[length-1]
	*m = tmp[0 : length-1]
	return removed

}
func (m *MaxInt) Peak() interface{} {
	if m.Len() > 0 {
		tmp := *m
		val := tmp[0]
		return val
	}
	return nil
}
*/

/*
Given ‘M’ sorted arrays, find the K’th smallest number among all the arrays.

Restated: among all arrays, return k'th smallest number
Details:
* Max heap to identify largest of the smallest
* M arrays is a list of lists

Cases:
Input: L1=[2, 6, 8], L2=[3, 6, 7], L3=[1, 3, 4], K=5
Output: 4

Assumptions
* all the numbers fit in 32 bit int
* k'th count < M array's length
* all M arrays are same length

var maxHeap

for i < len(M[0]); i++
  for j each M array
	   if maxHeap Length < K-1
		    maxHeap push M[j][i]
		 else M[j][i] < maxHeap Peak
		   maxHeap.Pop()
			 maxHeap.Push M[j][i]

for maxHeap Length - K
  maxHeap.Pop

	return maxHeap.Pop
*/

func KthSmallestInMSortedArrays(marr [][]int, k int) int {
	minheap := &MinHeap2{}
	heap.Init(minheap)
	for i := 0; i < len(marr); i++ {
		item := &Item2{ArrIndex: i, ElemIndex: 0, Data: marr}
		heap.Push(minheap, item)
	}

	result := 0
	count := 0
	for minheap.Len() > 0 {
		curr := minheap.Pop().(*Item2)
		result = marr[curr.ArrIndex][curr.ElemIndex]
		count++
		fmt.Println(result, curr)
		if count == k {
			break
		}

		curr.ElemIndex++
		if len(marr[curr.ArrIndex]) > curr.ElemIndex {
			heap.Push(minheap, curr)
		}
	}
	return result
}

type Item2 struct {
	ArrIndex  int
	ElemIndex int
	Data      [][]int
}

type MinHeap2 []*Item2

func (m MinHeap2) Len() int { return len(m) }
func (m MinHeap2) Less(i, j int) bool {
	data := m[i].Data
	return data[m[i].ArrIndex][m[i].ElemIndex] < data[m[j].ArrIndex][m[j].ElemIndex]
}
func (m MinHeap2) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *MinHeap2) Push(val interface{}) {
	tmp := *m
	tmp = append(tmp, val.(*Item2))
	*m = tmp
}

func (m *MinHeap2) Pop() interface{} {
	tmp := *m
	length := len(tmp)
	removed := tmp[length-1]
	*m = tmp[0 : length-1]
	return removed
}

/*
Merge sort
N Log N
*/

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merger(mergeSort(left), mergeSort(right))
}

func merger(arr1, arr2 []int) []int {
	n1n2 := len(arr1) + len(arr2)
	result := make([]int, n1n2)
	i := 0

	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			result[i] = arr1[0]
			arr1 = arr1[1:]
		} else {
			result[i] = arr2[0]
			arr2 = arr2[1:]
		}
		i++
	}

	for j := 0; j < len(arr1); j++ {
		result[i] = arr1[j]
		i++
	}

	for j := 0; j < len(arr1); j++ {
		result[i] = arr2[j]
		i++
	}

	return result
}
