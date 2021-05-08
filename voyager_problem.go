package fundamentals

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.

*/

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
	//fmt.Println(sortedOrder)
	return len(sortedOrder) == numCourses

}
