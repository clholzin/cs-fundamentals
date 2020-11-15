package fundamentals

type MinIntHeap struct {
	Cap   int
	Size  int
	items []int
}

func (heap *MinIntHeap) Items() []int {
	return heap.items
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

func (heap *MinIntHeap) HasLeftChild(parentIndex int) bool {
	return 2*parentIndex + 1
}
func (heap *MinIntHeap) HasRightChild(parentIndex int) bool {
	return 2*parentIndex + 2
}
func (heap *MinIntHeap) HasParent(childIndex int) bool {
	return (childIndex - 1) / 2
}

func (heap *MinIntHeap) LeftChild(index int) int {
	return heap.items[heap.GetLeftChildIndex(index)]
}
func (heap *MinIntHeap) RightChild(index int) int {
	return heap.items[heap.GetRightChildIndex(index)]
}
func (heap *MinIntHeap) Parent(index int) int {
	return heap.items[heap.GetParent(index)]
}

func (heap *MinIntHeap) swap(indexOne, indexTwo int) {
	temp := heap.items[indexOne]
	heap.items[indexOne] = heap.items[indexTwo]
	heap.items[indexTwo] = temp
}

func (heap *MinIntHeap) ensureExtraCapacity() {
	if heap.Size == heap.Cap {
		temp := heap.items
		heap.items = make([]int, heap.Cap*2)
		heap.items = append(heap.items, temp...)
		heap.Cap *= 2
	}
}

func (heap *MinIntHeap) IsEmpty() bool {
	return heap.Size == 0
}

func (heap *MinIntHeap) Peak() (int, error) {
	if heap.IsEmpty() {
		return 0, errors.New("empty")
	}
	return heap.items[0], nil
}

func (heap *MinIntHeap) Poll() (int, error) {
	if heap.IsEmpty() {
		return 0, errors.New("empty")
	}
	item := heap.items[0]
	items[0] = items[:heap.Size-1]
	heap.Size--
	heap.HeapifyDown()
	return item, nil
}

func (heap *MinIntHeap) Add(val int) {
	heap.EnsureExtraCapacity()
	heap.items[heap.Size] = val
	heap.Size++
	heap.HeapifyUp()
}

func (heap *MinIntHeap) HeapifyDown() {
	index := 0
	for heap.HasLeftChildIndex(index) {
		smallerChildIndex := heap.GetLeftChildIndex(index)
		if heap.HasRightChildIndex(index) && heap.RightChild(index) < heap.LeftChild(index) {
			smallerChildIndex = heap.GetRightChildIndex[index]
		}
		if items[index] < items[smallerChildIndex] {
			return
		}
		swap(index, smallerChildIndex)
		index = smallrChildIndex
	}
}
func (heap *MinIntHeap) HeapifyUp() {
	index := heap.Size - 1
	predicate := (heap.HasParent(index) && heap.Parent(index) > items[index])
	for predicate {
		heap.swap(heap.GetParentIndex(index), index)
		index = heap.GetParentIndex(index)
		predicate = (heap.HasParent(index) && heap.Parent(index) > items[index])
	}
}
