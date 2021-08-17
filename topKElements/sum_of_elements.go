package fundamentals

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	fmt.Println(x, old)
	return x
}

func (h *IntHeap) Peak() interface{} {
	if (*h).Len() > 0 {

		return (*h)[0]
	}
	return 0
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	vals := []int{22, 1, 3, 12, 5, 15, 11}
	fmt.Println(k1k2Sum(vals, 3, 6))
}

func k1k2Sum(arr []int, k1, k2 int) int {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	for i, val := range arr {
		fmt.Println("peak ", h.Peak().(int))
		if i < (k2 - 1) {
			heap.Push(h, val)
		} else if val < h.Peak().(int) {

			heap.Pop(h)
			heap.Push(h, val)
		}
	}
	sum := 0
	for i := 0; i < (k2 - k1 - 1); i++ {

		peak := h.Peak().(int)
		val := heap.Pop(h).(int)
		sum += val
		fmt.Println("peak ", peak, val)

	}

	return sum

}
