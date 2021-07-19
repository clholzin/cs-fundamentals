package fundamentals

//https://leetcode.com/problems/the-maze-ii/submissions/

/*
test case:
maze  [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]
start [0,4]
end   [4,4]

log
[
  [6 7 U 1 0]
	[5 U 9 2 3]
	[6 9 8 U 2]
	[U U U U U]
	[12 U 10 U 12]
	]

*/

import (
	"errors"
	"fmt"
)

const IntMaxBFS int = int(2147483647)

func shortestDistance(maze [][]int, start []int, destination []int) int {
	distance := make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		distance[i] = make([]int, len(maze[0]))
		for j := 0; j < len(maze[0]); j++ {
			distance[i][j] = IntMaxBFS
		}
	}

	distance[start[0]][start[1]] = 0
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	queue := NewBfsQueue(QueueBfsNode(start))
	for !queue.IsEmpty() {
		place, _ := queue.Remove()
		for _, dir := range dirs {
			x := place[0] + dir[0]
			y := place[1] + dir[1]
			count := 0
			for x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0]) && maze[x][y] == 0 {
				x += dir[0]
				y += dir[1]
				count++
			}
			check := (distance[place[0]][place[1]] + count) < distance[x-dir[0]][y-dir[1]]
			if check {
				distance[x-dir[0]][y-dir[1]] = distance[place[0]][place[1]] + count
				next := []int{x - dir[0], y - dir[1]}
				queue.Add(next)
			}
		}
	}

	if distance[destination[0]][destination[1]] == IntMaxBFS {
		return -1
	}
	fmt.Println(distance)
	return distance[destination[0]][destination[1]]

}

type BfsQueuer interface {
	IsEmpty() bool
	Peak() ([]int, error)
	Add(data []int)
	Remove() ([]int, error)
}

type BfsQueue struct {
	Head *BfsNode
	Tail *BfsNode
}

type BfsNode struct {
	Data []int
	Next *BfsNode
}

func QueueBfsNode(data []int) *BfsNode {
	return &BfsNode{Data: data}
}

func NewBfsQueue(head *BfsNode) BfsQueuer {
	return &BfsQueue{Head: head, Tail: head}
}

func (q *BfsQueue) IsEmpty() bool {
	return q.Head == nil
}

func (q *BfsQueue) Peak() ([]int, error) {
	if !q.IsEmpty() {
		return q.Head.Data, nil
	}
	return nil, errors.New("no head")

}

func (q *BfsQueue) Enqueue(data []int) {
	q.Add(data)
}

func (q *BfsQueue) Dequeue() ([]int, error) {
	return q.Remove()
}

func (q *BfsQueue) Add(data []int) {
	node := QueueBfsNode(data)
	if q.Tail != nil {
		q.Tail.Next = node
	}
	q.Tail = node
	if q.Head == nil {
		q.Head = node
	}
}

func (q *BfsQueue) Remove() ([]int, error) {
	var data []int
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
