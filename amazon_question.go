package fundamentals

/*
A Maze is given as N*N binary matrix of blocks where source block is the upper left most block i.e., maze[0][0] and destination
block is lower rightmost block i.e., maze[N-1][N-1]. A rat starts from source and has to reach the destination. The rat can move in
4 directions: up, right, down, left. 0 means the block is a dead end and 1 means the block can be used in the path from source to destination.

Example:
{1, 0, 0, 0}
{1, 1, 0, 1}
{0, 1, 0, 0}
{1, 1, 1, 1}

Possible path:


{6, 5, 4, 3}
{5, 4, 0, 2}
{0, 3, 0, 1}
{0, 2, 1, 1}

Return the length of the shortest path, or -1 if there is no possible path.

data := [][]int{
    {1, 0, 0, 0}
    {1, 1, 0, 1}
    {0, 1, 0, 0}
    {0, 1, 1, 1}
}
startrow := 0
startcolumn := 0
endrow := len(data[0])
endcolumn := len(data)
var count int
prev := [][]int{
    {1, 0, 0, 0}
    {0, 0, 0, 0}
    {0, 0, 0, 0}
    {0, 0, 0, 0}
}
for i:=startrow;i>len(data);i++ {
    row := data[i]
    if i == startrow {
        continue
    }

    for j:=0;j<len(row);j++ {
       if j == startcolumn {
        continue
       }
       val := row[j]
       if val > 0 {//&& j < len(row) && row[j+1]  {
           path := fmt.Sprintf("%d-%d",i,j)
           prev[path] = val
           //left
           //right
           //top
           //bottom
           count++
       }else{
           break
       }
    }
}

{1, 0, 0, 0}
{1, 1, 0, 1}
{0, 1, 0, 0}
{1, 1, 1, 1}
*/
func start() {
	c := -paths(data, 0, 1)    //right
	down := -paths(data, 1, 0) //down
}

func paths(data [][]int, currow, currcol int) int {
	v := data[currrow][currcol]
	if v == 0 {
		return 0
	} else {
		return min(paths(data, currow, currcol+1), paths(data, currow+1, currcol))
	}
}

func min(a, b int) {
	if a < b {
		return a
	}
	return b
}
