package fundamentals

import "math"

func shortestDistance3(maze [][]int, start, dest []int) int {

	distance := make([][]int, len(maze))

	queue := make([][]int, 0)
	queue = append(queue, start)

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[0]); j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	distance[start[0]][start[1]] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, dir := range dirs {
			x := curr[0] + dir[0]
			y := curr[1] + dir[1]
			count := 0

			for x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0]) && maze[x][y] == 0 {
				x += dir[0]
				y += dir[1]
				count++
			}

			nextCount := distance[curr[0]][curr[1]] + count
			currentCount := distance[x-dir[0]][y-dir[1]]
			if nextCount < currentCount {
				distance[x-dir[0]][y-dir[1]] = nextCount
				next := []int{x - dir[0], y - dir[y]}
				queue = append(queue, next)
			}
		}
	}

	if distance[dest[0]][dest[1]] == math.MaxInt32 {
		return -1
	}

	return distance[dest[0]][dest[1]]
}

/*
func shortestDistance3(maze [][]int, start, dest []int) int {

	distance := make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		distance[i] = make([]int, len(maze[0]))
		for j := 0; j < len(maze[i]); j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	distance[start[0]][start[1]] = 0

	dirs := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	que := [][]int{start}

	for len(que) > 0 {
		curr := que[0]
		que = que[1:]

		for _, dir := range dirs {
			x := curr[0] + dir[0]
			y := curr[1] + dir[1]
			count := 0

			for x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0]) && maze[x][y] == 0 {
				x += dir[0]
				y += dir[1]
				count++
			}

			nextCount := distance[curr[0]][curr[1]] + count
			existingCount := distance[x-dir[0]][y-dir[1]]

			if nextCount < existingCount {
				distance[x-dir[0]][y-dir[1]] = nextCount
				next := []int{x - dir[0], y - dir[1]}
				que = append(que, next)
			}
		}
	}

	if distance[dest[0]][dest[1]] == math.MaxInt32 {
		return -1
	}

	return distance[dest[0]][dest[1]]
}*/
