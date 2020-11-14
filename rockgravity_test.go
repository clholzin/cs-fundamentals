package fundamentals

import (
	"fmt"
	"strings"
	"testing"
)

func TestRocks1(t *testing.T) {
	input := `  .
 . 
:--
   
 ..
.  `

	columnRows := [][]string{}

	rows := strings.Split(input, "\n")
	for rowIndy, val := range rows {
		column := strings.Split(val, "")
		fmt.Println(column)
		columnRows = append(columnRows, make([]string, 0))
		columnRows[rowIndy] = append(columnRows[rowIndy], column...)
	}

	rowIndex := len(columnRows) - 1
	for colIndex, _ := range columnRows[0] {
		columnRows = roll(columnRows, rowIndex, colIndex)
	}
	fmt.Println("---")
	for _, col := range columnRows {
		fmt.Println(col)
	}
}

func TestRocks5(t *testing.T) {
	input := `:.   :
 .   .
.. :. 
- . - 
.-  : 
.  .  `

	columnRows := [][]string{}

	rows := strings.Split(input, "\n")
	for rowIndy, val := range rows {
		column := strings.Split(val, "")
		fmt.Println(column)
		columnRows = append(columnRows, make([]string, 0))
		columnRows[rowIndy] = append(columnRows[rowIndy], column...)
	}

	rowIndex := len(columnRows) - 1
	for colIndex, _ := range columnRows[0] {
		columnRows = roll(columnRows, rowIndex, colIndex)
	}
	fmt.Println("---")

	for _, col := range columnRows {
		fmt.Println(col)
	}
}
