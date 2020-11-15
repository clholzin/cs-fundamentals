package fundamentals

import "errors"

type Stacker interface {
	IsEmpty() bool
	Peak() (int, error)
	Push(data int)
	Pop() error
}

type Stack struct {
	Top *StackNode
}

type StackNode struct {
	Data int
	Next *StackNode
}

func NewStack(data int) Stacker {
	return &Stack{Top: &StackNode{Data: data}}
}

func (stack *Stack) IsEmpty() bool {
	return stack.Top == nil
}

func (stack *Stack) Peak() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return stack.Top.Data, nil
}

func (stack *Stack) Push(data int) {
	if stack.IsEmpty() {
		stack.Top = &StackNode{Data: data}
		return
	}

	stack.Top = &StackNode{Data: data, Next: stack.Top}
}

func (stack *Stack) Pop() error {
	if stack.IsEmpty() {
		return errors.New("stack is empty")
	}

	stack.Top = stack.Top.Next
	return nil

}
