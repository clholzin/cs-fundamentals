package fundamentals

import "strconv"

const boxsize int = 3
const rowsize int = 9

var (
	rows         [][]int
	columns      [][]int
	boxes        [][]int
	gboard       [][]byte
	sudokuSolved bool
)

func init() {
	rows = make([][]int, rowsize)
	columns = make([][]int, rowsize)
	boxes = make([][]int, rowsize)
	gboard = make([][]byte, rowsize)

	for index := range rows {
		rows[index] = make([]int, rowsize+1)
		columns[index] = make([]int, rowsize+1)
		boxes[index] = make([]int, rowsize+1)
	}
}

func solveSudoku(board [][]byte) {
	gboard = board
	for i := 0; i < rowsize; i++ {
		for j := 0; j < rowsize; j++ {
			var val byte = board[i][j] - '0'
			if val != '.' {
				data, _ := strconv.Atoi(string(val))
				placeNumber(data, i, j)
			}
		}
	}
	backtrack(0, 0)
}

func couldPlace(data, row, col int) bool {
	index := (row/boxsize)*boxsize + col/boxsize
	return rows[row][data]+columns[col][data]+boxes[index][data] == 0
}

func placeNumber(data, row, col int) {
	index := (row/boxsize)*boxsize + col/boxsize
	rows[row][data]++
	columns[row][data]++
	boxes[row][index]++
	gboard[row][col] = byte('0' + data)
}

func removeNumber(data, row, col int) {
	index := (row/boxsize)*boxsize + col/boxsize
	rows[row][data]--
	columns[row][data]--
	boxes[row][index]--
	gboard[row][col] = '.'
}

func placeNextNumbers(row, col int) {
	if col == rowsize-1 && row == rowsize-1 {
		sudokuSolved = true
	} else {
		if col == rowsize-1 {
			backtrack(row+1, 0)
		} else {
			backtrack(row, col+1)
		}
	}
}

func backtrack(row, col int) {
	if gboard[row][col] == '.' {
		for d := 1; d < 10; d++ {
			if couldPlace(d, row, col) {
				placeNumber(d, row, col)
				placeNextNumbers(row, col)
				if !sudokuSolved {
					removeNumber(d, row, col)
				}
			}
		}
	} else {
		placeNextNumbers(row, col)
	}
}
