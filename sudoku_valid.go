package fundamentals

func isValidSudoku(board [][]byte) bool {

	rowchecker := make([]map[byte]int, len(board))
	colchecker := make([]map[byte]int, len(board))
	for i := 0; i < len(board); i++ {
		rowchecker[i] = make(map[byte]int)
	}
	for i := 0; i < len(board); i++ {
		colchecker[i] = make(map[byte]int)
	}
	empty := byte(46)

	for i, row := range board {
		for j, col := range row {
			if col != empty {
				rowchecker[i][atoitwo(col)]++
				colchecker[j][atoitwo(col)]++
			}
		}
	}

	allrowscolumns := append(rowchecker, colchecker...)
	//fmt.Println("check row column")
	for _, row := range allrowscolumns {
		for _, val := range row {
			if val > 1 {
				return false
			}
		}
	}
	//fmt.Println("check boxes")

	for len(board) > 0 {
		b1 := make([][]byte, 3)
		b2 := make([][]byte, 3)
		b3 := make([][]byte, 3)
		for i := 0; i < 3; i++ {
			b1[i] = make([]byte, 3)
			b2[i] = make([]byte, 3)
			b3[i] = make([]byte, 3)
		}
		row := make([]byte, 9)
		for i := 0; i < 3; i++ {
			copy(row, board[i])
			b1[i] = row[:3]
			row = row[3:]
			b2[i] = row[:3]
			row = row[3:]
			b3[i] = row[:]
			row = make([]byte, 9)

		}
		board = board[3:]
		//fmt.Println(board)
		for _, bd := range [][][]byte{b1, b2, b3} {
			rowchecker = make([]map[byte]int, 3)
			for i := 0; i < 3; i++ {
				rowchecker[i] = make(map[byte]int)
			}
			for _, row := range bd {
				for _, col := range row {
					if col != empty {
						rowchecker[0][atoitwo(col)]++
					}
				}
			}
			// fmt.Println(rowchecker)

			for _, row := range rowchecker {
				for _, val := range row {
					if val > 1 {
						return false
					}
				}
			}

		}

	}

	return true
}

func atoitwo(a byte) byte {
	return a - '0'
}
