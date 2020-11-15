package fundamentals

import (
	"errors"
	"fmt"
)

type MinIntHeap struct {
	Cap   int
	Size  int
	items []int
}

func NewMinIntHeap() *MinIntHeap {
	return &MinIntHeap{Cap: 1}
}

func (heap *MinIntHeap) Values() []int {
	return heap.items[:]
}

func (heap *MinIntHeap) GetLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}
func (heap *MinIntHeap) GetRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}
func (heap *MinIntHeap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (heap *MinIntHeap) HasLeftChild(index int) bool {
	return heap.GetLeftChildIndex(index) < heap.Size
}
func (heap *MinIntHeap) HasRightChild(index int) bool {
	return heap.GetRightChildIndex(index) < heap.Size
}
func (heap *MinIntHeap) HasParent(index int) bool {
	return heap.GetParentIndex(index) >= 0
}

func (heap *MinIntHeap) LeftChild(index int) int {
	return heap.items[heap.GetLeftChildIndex(index)]
}
func (heap *MinIntHeap) RightChild(index int) int {
	return heap.items[heap.GetRightChildIndex(index)]
}
func (heap *MinIntHeap) Parent(index int) int {
	return heap.items[heap.GetParentIndex(index)]
}

func (heap *MinIntHeap) swap(indexOne, indexTwo int) {
	temp := heap.items[indexOne]
	heap.items[indexOne] = heap.items[indexTwo]
	heap.items[indexTwo] = temp
}

func (heap *MinIntHeap) IsEmpty() bool {
	return heap.Size == 0
}

func (heap *MinIntHeap) Peak() (int, error) {
	if heap.IsEmpty() {
		return 0, errors.New("empty")
	}
	return heap.items[heap.Cap], nil
}

func (heap *MinIntHeap) Poll() (int, error) {
	if heap.IsEmpty() {
		return 0, errors.New("empty")
	}
	item := heap.items[0]
	heap.items[0] = heap.items[heap.Size-1]
	heap.Size--
	heap.heapifyDown()
	fmt.Println("-", heap.items, heap.Size, heap.Cap)
	return item, nil
}

func (heap *MinIntHeap) ensureExtraCapacity() {
	if heap.items == nil {
		heap.items = make([]int, 1)
	}
	if heap.Size == heap.Cap {
		temp := heap.items[:]
		heap.items = make([]int, heap.Cap*2)
		heap.items = append(heap.items[:], temp...)
		heap.Cap *= 2
	}
}

func (heap *MinIntHeap) Add(val int) {
	heap.ensureExtraCapacity()
	heap.items[heap.Size] = val
	heap.Size++
	heap.heapifyUp()
	fmt.Println("+", heap.items, heap.Size, heap.Cap, val)
}

func (heap *MinIntHeap) heapifyDown() {
	index := 0
	for heap.HasLeftChild(index) {
		smallerChildIndex := heap.GetLeftChildIndex(index)
		if heap.HasRightChild(index) && heap.RightChild(index) > heap.LeftChild(index) {
			smallerChildIndex = heap.GetRightChildIndex(index)
		}
		if heap.items[index] < heap.items[smallerChildIndex] {
			return
		}
		heap.swap(index, smallerChildIndex)
		index = smallerChildIndex
	}
}

func (heap *MinIntHeap) heapifyUp() {
	index := heap.Size - 1
	predicate := heap.HasParent(index) && (heap.Parent(index) < heap.items[index])
	for predicate {
		heap.swap(heap.GetParentIndex(index), index)
		index = heap.GetParentIndex(index)
		predicate = heap.HasParent(index) && (heap.Parent(index) < heap.items[index])
	}
}
