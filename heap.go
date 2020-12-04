package fundamentals

import (
	"errors"
	"fmt"
)

type MinIntHeap struct {
	cap   int
	size  int
	items []int
}

func NewMinIntHeap() *MinIntHeap {
	return &MinIntHeap{cap: 10}
}

func (h *MinIntHeap) Items() []int {
	return h.items[:]
}

func (h *MinIntHeap) Size() int {
	return h.size
}

func (h *MinIntHeap) GetLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}
func (h *MinIntHeap) GetRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}
func (h *MinIntHeap) GetParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MinIntHeap) HasLeftChild(index int) bool {
	return h.GetLeftChildIndex(index) < h.size
}
func (h *MinIntHeap) HasRightChild(index int) bool {
	return h.GetRightChildIndex(index) < h.size
}
func (h *MinIntHeap) HasParent(index int) bool {
	return h.GetParentIndex(index) >= 0
}

func (h *MinIntHeap) LeftChild(index int) int {
	return h.items[h.GetLeftChildIndex(index)]
}
func (h *MinIntHeap) RightChild(index int) int {
	return h.items[h.GetRightChildIndex(index)]
}
func (h *MinIntHeap) Parent(index int) int {
	return h.items[h.GetParentIndex(index)]
}

func (h *MinIntHeap) swap(indexOne, indexTwo int) {
	temp := h.items[indexOne]
	h.items[indexOne] = h.items[indexTwo]
	h.items[indexTwo] = temp
}

func (h *MinIntHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MinIntHeap) Peak() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("empty")
	}
	return h.items[0], nil
}

func (h *MinIntHeap) Poll() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("empty")
	}
	item := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.items[h.size-1] = 0
	h.size--
	h.heapifyDown()
	fmt.Println("-", h.items, h.size, h.cap)
	return item, nil
}

func (h *MinIntHeap) ensureExtracapacity() {
	if h.items == nil {
		h.items = make([]int, h.cap)
	}
	if h.size == h.cap {
		temp := h.items[:]
		h.items = make([]int, h.cap*2)
		h.items = append(h.items[:], temp...)
		h.cap *= 2
	}
}

func (h *MinIntHeap) Add(val int) {
	h.ensureExtracapacity()
	h.items[h.size] = val
	h.size++
	h.heapifyUp()
	fmt.Println("+", h.items, h.size, h.cap, val)
}

func (h *MinIntHeap) heapifyDown() {
	index := 0
	for h.HasLeftChild(index) {
		smallerChildIndex := h.GetLeftChildIndex(index)
		//fmt.Println("[]", h.items, h.size, h.cap, smallerChildIndex)

		if h.HasRightChild(index) && h.RightChild(index) < h.LeftChild(index) {
			smallerChildIndex = h.GetRightChildIndex(index)
		}
		if h.items[index] < h.items[smallerChildIndex] {
			break
		} else {
			h.swap(index, smallerChildIndex)
		}
		index = smallerChildIndex
	}
}

func (h *MinIntHeap) heapifyUp() {
	index := h.size - 1
	for h.HasParent(index) && (h.Parent(index) > h.items[index]) {
		h.swap(h.GetParentIndex(index), index)
		index = h.GetParentIndex(index)
		//fmt.Println("^", h.items, h.size, h.cap, index)
	}
}
