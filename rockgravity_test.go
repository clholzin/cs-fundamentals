package fundamentals

import (
	"fmt"
	"strings"
	"testing"
)

func TestRocks1(t *testing.T) {
	input := `  
 .
:-
  
 .
. `

	columnRows := [][]string{}

	rows := strings.Split(input, "\n")
	for rowIndy, val := range rows {
		column := strings.Split(val, "")
		fmt.Println(column, len(column))
		columnRows = append(columnRows, make([]string, 0))
		columnRows[rowIndy] = append(columnRows[rowIndy], column...)
	}

	rowIndex := len(columnRows) - 1
	for colIndex, _ := range columnRows[0] {
		fmt.Println(rowIndex, colIndex)
		roll(columnRows, rowIndex, colIndex)
	}
	for _, col := range columnRows {
		fmt.Println(col)

	}

}

func TestRocks5(t *testing.T) {
	input := `
  .      .r
: .  :    r
-   .  -  r
. -       r
.    .    `

	columnRows := [][]string{}

	rows := strings.Split(input, "r")
	fmt.Println(rows, len(rows))
	for rowIndy, val := range rows {
		column := strings.Split(val, " ")
		columnRows[rowIndy] = []string{}
		columnRows[rowIndy] = append(columnRows[rowIndy], column...)
	}

	rowCount := len(rows)
	for rowIndex, row := range columnRows {
		for colIndex, _ := range row {
			roll(columnRows, rowCount-rowIndex, colIndex)
		}
	}

}
