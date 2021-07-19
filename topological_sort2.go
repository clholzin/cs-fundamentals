package fundamentals

// tags: topological sort
// time V + E
// space V + E

func findOrder2(numCourses int, prerequisites [][]int) []int {

	parents := make(map[int][]int)
	edges := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		parents[i] = []int{}
		edges[i] = 0
	}

	for _, course := range prerequisites {
		child, parent := course[0], course[1]
		edges[child] = edges[child] + 1
		parents[parent] = append(parents[parent], child)
	}

	queue := []int{}

	for child, count := range edges {
		if count == 0 {
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
			edges[child]--
			if edges[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if len(sorted) != numCourses {
		sorted = sorted[:0]
	}

	return sorted

}
