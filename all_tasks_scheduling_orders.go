package fundamentals

import (
	"fmt"
)

func printOrders(tasks int, prerequisites [][]int) {
	sortedOrder := make([]int, 0)

	if tasks <= 0 {
		return
	}

	degree := make(map[int]int)
	parents := make(map[int][]int)

	for i := 0; i < tasks; i++ {
		degree[i] = 0
		parents[i] = make([]int, 0)
	}

	for i := 0; i < len(prerequisites); i++ {
		parent, child := prerequisites[i][0], prerequisites[i][1]
		parents[parent] = append(parents[parent], child)
		degree[child]++
	}

	queue := make([]int, 0)
	for child, count := range degree {
		if count == 0 {
			queue = append(queue, child)
		}
	}

	printAllTopologicalSorts(queue, sortedOrder, parents, degree)

}

func printAllTopologicalSorts(queue, sortedOrder []int, parents map[int][]int, degree map[int]int) {

	if len(queue) > 0 {
		for len(queue) > 0 {
			child := queue[0]
			queue = queue[1:]
			sortedOrder = append(sortedOrder, child)
			sortedChildIndx := len(sortedOrder) - 1
			nextQueue := make([]int, len(queue))
			copy(queue, nextQueue)
			children := parents[child]
			for _, nextChild := range children {
				degree[nextChild]--
				if degree[nextChild] == 0 {
					nextQueue = append(nextQueue, nextChild)
				}
			}

			printAllTopologicalSorts(nextQueue, sortedOrder, parents, degree)

			if sortedChildIndx == 0 {
				sortedOrder = sortedOrder[1:]
			} else if sortedChildIndx == len(sortedOrder)-1 {
				sortedOrder = sortedOrder[:len(sortedOrder)-1]
			} else {
				sortedHalf := sortedOrder[:sortedChildIndx-1]
				sortedOrder = sortedOrder[sortedChildIndx:]
				sortedOrder = append(sortedHalf, sortedOrder...)
			}

			for _, nextChild := range children {
				degree[nextChild]++
			}
		}
	}

	if len(sortedOrder) == len(degree) {
		fmt.Println(sortedOrder)
	}
}
