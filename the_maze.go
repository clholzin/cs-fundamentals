package fundamentals

import "fmt"

// # amazon question answer
func hasPath(maze [][]int, start []int, destination []int) bool {
	dest := destination
	var startcopy []int = make([]int, len(start))
	copy(startcopy, start)
	var halt []bool
	tries := []string{"up", "down", "left", "right"}
	for _, dir := range tries {
		memo := make(map[string]int)
		direction(memo, &halt, maze, cont(startcopy, dir), start, dest, dir)
		if len(halt) > 0 {
			return halt[0]
		}
	}
	return false
}

func direction(memo map[string]int, halt *[]bool, maze [][]int, curr, start, dest []int, dir string) bool {
	if len(*halt) > 0 {
		return true
	}

	if checkend(curr, dest) {
		canstop := checkwall(maze, cont(curr, dir))
		if canstop {
			*halt = append(*halt, canstop)
			return true
		}
	}

	if checkstart(curr, start) {
		return false
	}

	if checkwall(maze, curr) {
		return false
	}

	hash := fmt.Sprintf("%s-%d-%d", dir, curr[0], curr[1])
	if _, ok := memo[hash]; ok {
		return false
	} else {
		memo[hash] = 0
	}

	var horizontal, vertical bool = true, true

	switch dir {
	case "left":
		horizontal = checkwall(maze, cont(curr, dir))
		if horizontal {
			return direction(memo, halt, maze, cont(curr, "up"), curr, dest, "up") ||
				direction(memo, halt, maze, cont(curr, "down"), curr, dest, "down") ||
				direction(memo, halt, maze, cont(curr, "right"), curr, dest, "right")
		}
		break
	case "right":
		horizontal = checkwall(maze, cont(curr, dir))
		if horizontal {
			return direction(memo, halt, maze, cont(curr, "down"), curr, dest, "down") ||
				direction(memo, halt, maze, cont(curr, "up"), curr, dest, "up") ||
				direction(memo, halt, maze, cont(curr, "left"), curr, dest, "left")
		}
		break
	case "down":
		vertical = checkwall(maze, cont(curr, "down"))
		if vertical {
			return direction(memo, halt, maze, cont(curr, "left"), curr, dest, "left") ||
				direction(memo, halt, maze, cont(curr, "right"), curr, dest, "right") ||
				direction(memo, halt, maze, cont(curr, "up"), curr, dest, "up")
		}
		break
	case "up":
		vertical = checkwall(maze, cont(curr, dir))
		if vertical {
			return direction(memo, halt, maze, cont(curr, "right"), curr, dest, "right") ||
				direction(memo, halt, maze, cont(curr, "left"), curr, dest, "left") ||
				direction(memo, halt, maze, cont(curr, "down"), curr, dest, "down")
		}
		break
	}
	return direction(memo, halt, maze, cont(curr, dir), curr, dest, dir)
}

func cont(pos []int, dir string) []int {
	cp := make([]int, len(pos))
	copy(cp, pos)
	switch dir {
	case "left":
		cp[1] -= 1
		break
	case "right":
		cp[1] += 1
		break
	case "down":
		cp[0] += 1
		break
	case "up":
		cp[0] -= 1
		break
	}

	return cp
}

func checkwall(maze [][]int, pos []int) bool {
	if (pos[0] > len(maze)-1 || pos[0] < 0) || (pos[1] >
		len(maze[0])-1 || pos[1] < 0) {
		return true
	}
	return maze[pos[0]][pos[1]] == 1
}

func checkstart(pos, start []int) bool {
	return compare(pos, start)
}

func checkend(pos, dest []int) bool {
	return compare(pos, dest)
}

func compare(first, second []int) bool {
	return first[0] == second[0] && first[1] == second[1]
}
