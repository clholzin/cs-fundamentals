package fundamentals

import "fmt"

// tags: dfs,

func countComponents(n int, edges [][]int) int {
	relations := make(map[int][]int)
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		relations[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		child, parent := edges[i][0], edges[i][1]
		relations[parent] = append(relations[parent], child)
		relations[child] = append(relations[child], parent)

	}
	var components int

	for i := 0; i < len(visited); i++ {
		if visited[i] == 0 {
			components++
			dfs(relations, visited, i)
		}
	}

	return components
}

func dfs(relations map[int][]int, visited []int, index int) {
	visited[index] = 1
	for i := 0; i < len(relations[index]); i++ {
		if visited[relations[index][i]] == 0 {
			dfs(relations, visited, relations[index][i])
		}
	}
}

func countComponentsFail(n int, edges [][]int) int {

	parents := make(map[int][]int)
	children := make(map[int]int)
	for i := 0; i < n; i++ {
		parents[i] = []int{}
		children[i] = 0
	}

	for i := 0; i < len(edges); i++ {
		child, parent := edges[i][0], edges[i][1]
		parents[parent] = append(parents[parent], child)
		children[child]++
	}

	queue := make([]int, 0)
	count := 0
	visited := make(map[int]int)
	fmt.Println(queue, parents, children)
	for index := range children {
		// if _,ok := visited[index]; !ok && val == 0 {
		fmt.Println(index)
		queue := append(queue, index)
		num := bfs(queue, visited, children, parents)
		count = count + num
		fmt.Println(visited)
		// }
	}
	fmt.Println(queue, parents, children)
	return count
}

func bfs(queue []int, visited map[int]int, children map[int]int, parents map[int][]int) int {
	var result int = 1
	currSession := make(map[int]int)
	for len(queue) > 0 {
		position := queue[0]
		queue = queue[1:]
		currSession[position]++
		if children[position] > 0 {
			children[position]--
			if children[position] < 0 {
				continue
			}
		}
		if len(parents[position]) > 0 {
			next := parents[position]
			queue = append(queue, next...)
		}
	}
	for i, val := range currSession {
		if _, ok := visited[i]; ok && result == 1 {
			result = 0
		}
		visited[i] = val
	}
	return result
}
