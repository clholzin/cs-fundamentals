package fundamentals

// tags: topological sort

func findOrder(numCourses int, prerequisites [][]int) []int {

	parents := make(map[int][]int)
	nDegree := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		parents[i] = []int{}
		nDegree[i] = 0
	}

	for _, vals := range prerequisites {
		child, parent := vals[0], vals[1]
		nDegree[child] = nDegree[child] + 1
		parents[parent] = append(parents[parent], child)
	}

	queue := []int{}

	for child, degree := range nDegree {
		if degree == 0 {
			queue = append(queue, child)
		}
	}

	sorted := []int{}

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		children := parents[parent]
		sorted = append(sorted, parent)
		for _, child := range children {
			nDegree[child]--
			if nDegree[child] == 0 {
				queue = append(queue, child)
			}
		}

	}

	if len(sorted) != numCourses {
		sorted = sorted[:0]
	}

	return sorted

}
