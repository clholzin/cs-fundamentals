package list

type Lister interface {
	Push(value int) bool
	Pop() bool
	Shift() bool
	FindIndex(value int) (int, bool)
	Value(index int) (int, bool)
	Count() int
	Delete(index int) (int, bool)
	Iterate() []int
}

type List struct {
	data []int
}

func NewList() Lister {
	return &List{data: []int{}}
}

func (n *List) Push(value int) bool {
	count := n.Count()
	n.data = append(n.data, value)
	if count < n.Count() {
		return true
	}
	return false
}

func (n *List) Shift() bool {
	count := n.Count()
	if count > 0 {
		n.data = n.data[1:]
	}
	if count > n.Count() {
		return true
	}
	return false
}

func (n *List) Pop() bool {
	count := n.Count()
	if count > 0 {
		n.data = n.data[:count-1]
	}
	if count > n.Count() {
		return true
	}
	return false
}

func (n *List) Value(index int) (int, bool) {
	if index >= n.Count() {
		return 0, false
	}

	return n.data[index], true
}

func (n *List) Iterate() []int {
	return n.data[:]
}

func (n *List) FindIndex(value int) (int, bool) {
	for i := 0; i < n.Count(); i++ {
		if value == n.data[i] {
			return i, true
		}
	}
	return 0, false
}

func (n *List) Delete(index int) (int, bool) {
	if index >= n.Count() {
		return 0, false
	}
	val := n.data[index]

	if index == 0 {
		n.data = n.data[1:]
	} else if index == n.Count()-1 {
		n.data = n.data[:index]
	} else if index > 0 && index < n.Count() {
		first := n.data[:index]
		second := n.data[index+1:]
		n.data = append(first, second...)
	}

	return val, true
}

func (n *List) Count() int {
	return len(n.data)
}
