package fundamentals

/*
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
*/
type MinStack struct {
	normalStack []int
	minStack    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{}
	return stack
}

func (this *MinStack) Push(val int) {
	this.normalStack = append(this.normalStack, val)
	if len(this.minStack) == 0 || this.minStack[len(this.minStack)-1] >= val {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	n1 := len(this.normalStack)
	val := this.normalStack[n1-1]
	this.normalStack = this.normalStack[:n1-1]
	n2 := len(this.minStack)
	if val == this.minStack[n2-1] {
		this.minStack = this.minStack[:n2-1]
	}
}

func (this *MinStack) Top() int {
	n1 := len(this.normalStack)
	return this.normalStack[n1-1]
}

func (this *MinStack) GetMin() int {
	n2 := len(this.minStack)
	return this.minStack[n2-1]
}
