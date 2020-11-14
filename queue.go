package fundamentals

import "errors"

type Queuer interface {
	IsEmpty() bool
	Peak() (int, error)
	Add(data int)
	Remove(data int) (int, error)
}

type Queue struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Next *Node
}

func NewQueueNode(data int) *Node {
	return &Node{Data: data}
}

func NewQueue(head *Node) *Queue {
	return &Queue{Head: head, Tail: head}
}

func (q *Queue) IsEmpty() bool {
	return q.Head == nil
}
func (q *Queue) Peak() (int, error) {
	if !q.IsEmpty() {
		return q.Head.Data, nil
	}
	return 0, errors.New("no head")

}
func (q *Queue) Add(data int) {
	node := NewQueueNode(data)
	if q.Tail != nil {
		q.Tail.Next = node
	}
	q.Tail = node
	if q.Head == nil {
		q.Head = node
	}
}
func (q *Queue) Remove() (int, error) {
	var data int
	if q.Head != nil {
		data = q.Head.Data
		q.Head = q.Head.Next
		if q.Head == nil {
			q.Tail = nil
		}
		return data, nil
	}
	return data, errors.New("no head")
}
