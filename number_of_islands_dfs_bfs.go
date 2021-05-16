package fundamentals

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	lenRow := len(grid)
	lenCol := len(grid[0])
	countIslands := 0

	for i := 0; i < lenRow; i++ {
		for j := 0; j < lenCol; j++ {
			if grid[i][j] == '1' {
				countIslands++
				//clearIslandsDFS(grid,i,j)
				clearIslandsBFS(grid, i, j)
			}
		}
	}
	return countIslands
}

/*BFS
Complexity Analysis

Time complexity : O(M x N)O(M×N) where MM is the number of rows and NN is the number of columns.

Space complexity : O(min(M, N))O(min(M,N)) because in worst case where the grid is filled with lands, the size of queue can grow up to min(M,NM,N).
*/
func clearIslandsBFS(grid [][]byte, row, col int) {
	lenRow := len(grid)
	lenCol := len(grid[0])

	directions := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} //up,right,down,left

	que := [][]int{{row, col}}
	for len(que) > 0 {
		curr := que[0]
		currow, curcol := curr[0], curr[1]
		grid[currow][curcol] = '0'
		que = que[1:]
		for _, dir := range directions {
			r := currow + dir[0]
			c := curcol + dir[1]
			if r >= 0 && c >= 0 && r < lenRow && c < lenCol && grid[r][c] == '1' {
				// fmt.Println(currow,curcol)
				grid[r][c] = '0'
				que = append(que, []int{r, c})
			}
		}
	}
}

/*DFS
Complexity Analysis

Time complexity : O(M x N)O(M×N) where MM is the number of rows and NN is the number of columns.

Space complexity : worst case O(M \times N)O(M×N) in case that the grid map is filled with lands where DFS goes by M \times NM×N deep.
*/

func clearIslandsDFS(grid [][]byte, row, col int) {
	lenRow := len(grid)
	lenCol := len(grid[0])
	if row < 0 || col < 0 || row >= lenRow || col >= lenCol || grid[row][col] == '0' {
		return
	}

	grid[row][col] = '0'
	clearIslandsDFS(grid, row-1, col)
	clearIslandsDFS(grid, row+1, col)
	clearIslandsDFS(grid, row, col-1)
	clearIslandsDFS(grid, row, col+1)
}
