package fundamentals

// tags: topological sort

func canFinish(numCourses int, prerequisites [][]int) bool {
	parents := map[int][]int{}
	ndegree := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		parents[i] = make([]int, 0)
		ndegree[i] = 0
	}

	for i := 0; i < len(prerequisites); i++ {
		child, parent := prerequisites[i][0], prerequisites[i][1]
		ndegree[child] = ndegree[child] + 1
		parents[parent] = append(parents[parent], child)
	}

	queue := []int{}
	for parent, v := range ndegree {
		if v == 0 {
			queue = append(queue, parent)
		}
	}
	sortedOrder := []int{}
	for len(queue) > 0 {
		parent := queue[0]
		sortedOrder = append(sortedOrder, parent)
		queue = queue[1:]
		children := parents[parent]
		for i := 0; i < len(children); i++ {
			child := children[i]
			ndegree[child]--
			if ndegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}
	return len(sortedOrder) == numCourses
}
