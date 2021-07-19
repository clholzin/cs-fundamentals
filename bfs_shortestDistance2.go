package fundamentals

//https://leetcode.com/problems/the-maze-ii/submissions/

/*
test case:
maze  [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]
start [0,4]
end   [4,4]

//path
[[0 4] [0 3] [2 4] [1 3] [1 4] [1 0] [0 0] [2 0] [0 1] [2 2] [2 1] [1 2] [4 2] [4 4] [4 0]]
//

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
	"math"
)

func shortestDistance2(maze [][]int, start []int, dest []int) int {
	// setup default distances
	distance := make([][]int, len(maze))
	for i := 0; i < len(maze); i++ {
		distance[i] = make([]int, len(maze[0]))
		for j := 0; j < len(maze[0]); j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	// add start location
	distance[start[0]][start[1]] = 0
	// follow each dir for each edge
	dirs := [][]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	// establish queue
	queue := [][]int{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// iterate over closest neighbors
		for _, dir := range dirs {
			// to the next direction
			x := curr[0] + dir[0]
			y := curr[1] + dir[1]
			count := 0
			// iterate over the next direction moving in that
			// direction while maze path == 0 and count steps
			for x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0]) && maze[x][y] == 0 {
				// continue to move and count
				x += dir[0]
				y += dir[1]
				count++
			}

			// check the distance + count is less than Max
			// it should be because we set the start to 0
			nextCount := distance[curr[0]][curr[1]] + count
			existingCount := distance[x-dir[0]][y-dir[1]]
			if nextCount < existingCount {
				// keep min distance and set it with new count of steps taken
				// update the last stepped location
				distance[x-dir[0]][y-dir[1]] = nextCount
				// add this next location to the queue to move from there
				next := []int{x - dir[0], y - dir[1]}
				queue = append(queue, next)
			}
		}
	}

	// check we can get to dest
	if distance[dest[0]][dest[1]] == math.MaxInt32 {
		return -1
	}

	// return last captured count for destination
	return distance[dest[0]][dest[1]]

}
