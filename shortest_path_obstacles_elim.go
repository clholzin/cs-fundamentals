package fundamentals

/*
type Queue struct {
	Head *Node
	Tail *Node
}
type Node struct {
	Key      string
	Row, Col int
	Next     *Node
	Count    int
}

func NewQueue(head *Node) *Queue {
	return &Queue{Head: head, Tail: head}
}

func (q *Node) IsEmpty() bool {
	return q == nil
}

func (q *Queue) Add(row, col int, key string) {
	n := &Node{
		Row:  row,
		Col:  col,
		Next: nil,
		Key:  key,
	}
	if q.Tail != nil {
		n.Count = q.Tail.Count
		q.Tail.Next = n
	}
	q.Tail = n
	q.Tail.Count += 1
	if q.Head == nil {
		q.Head = n
	}
	//fmt.Println(key,q.Tail)
}

func (q *Queue) Count() int {
	t := q.Head
	var count int
	for !t.IsEmpty() {
		fmt.Printf("%d-%d %d\n", t.Row, t.Col, t.Count)
		t = t.Next
		count++
	}
	return count
}

func shortestPath(grid [][]int, k int) int {
	var path *Queue = NewQueue(&Node{Row: 0, Col: 0, Count: 0})
	eliminate := make(map[string][]int)
	visited := make(map[string]int)

	shortestp(grid, len(grid)-1, len(grid[0])-1, k, eliminate, visited, path)

	count := path.Count()
	fmt.Println(visited, eliminate, len(eliminate))

	return count
}

func shortestp(grid [][]int, row, col, k int, elim map[string][]int, visited map[string]int, path *Queue) int {
	finishrow := len(grid) - 1
	finishcol := len(grid[0]) - 1
	if row < 0 || col < 0 || row > finishrow || col > finishcol {
		return 0
	}

	key := fmt.Sprintf("%d-%d", row, col)
	val, discovered := visited[key]
	if !discovered {
		visited[key] = grid[row][col]
		val = visited[key]
		if visited[key] == 0 {
			if row > 0 || col > 0 {
				path.Add(row, col, key)
			}
			// right
			if shortestp(grid, row, col+1, k, elim, visited, path) == 1 && shortestp(grid, row, col+2, k, elim, visited, path) == 0 {
				key = fmt.Sprintf("%d-%d", row, col+1)
				elim[key] = []int{row, col + 1}
			}
			// left
			if shortestp(grid, row, col-1, k, elim, visited, path) == 1 && shortestp(grid, row, col-2, k, elim, visited, path) == 0 {
				key = fmt.Sprintf("%d-%d", row, col-1)
				elim[key] = []int{row, col - 1}
			}
			// down
			if shortestp(grid, row+1, col, k, elim, visited, path) == 1 && shortestp(grid, row+2, col, k, elim, visited, path) == 0 {
				key = fmt.Sprintf("%d-%d", row+1, col)
				elim[key] = []int{row + 1, col}
			}

			// up
			if shortestp(grid, row-1, col, k, elim, visited, path) == 1 && shortestp(grid, row-2, col, k, elim, visited, path) == 0 {
				key = fmt.Sprintf("%d-%d", row-1, col)
				elim[key] = []int{row - 1, col}
			}

		}
	}
	return val
}*/
