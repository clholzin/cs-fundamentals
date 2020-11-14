package fundamentals

func roll(rows [][]string, rowIndex, columnIndex int) [][]string {

	above := rowIndex - 1
	if above < 0 {
		return rows
	}
	current := rows[rowIndex][columnIndex]
	aboveValue := rows[above][columnIndex]

	if somethingToMove(aboveValue) {
		if goodToMove(current, aboveValue) {
			if aboveValue == ":" && current == "." {
				rows[rowIndex][columnIndex] = ":"
				rows[above][columnIndex] = "."
			} else if aboveValue == "." && current == "." {
				rows[rowIndex][columnIndex] = ":"
				rows[above][columnIndex] = " "
			} else {
				rows[rowIndex][columnIndex] = aboveValue
				rows[above][columnIndex] = " "
			}
			if rowIndex != len(rows)-1 {
				rowIndex++
				return roll(rows, rowIndex, columnIndex)
			}
			return roll(rows, above, columnIndex)
		}
	}

	if above > 0 {
		return roll(rows, above, columnIndex)
	}

	return rows

}

func goodToMove(value, tomove string) bool {
	switch value {
	case " ":
		return true
	case ".":
		return true
	case ":":
		break
	default:
		break
	}
	return false
}

func somethingToMove(value string) bool {

	switch value {
	case ".":
		return true
	case ":":
		return true
	case " ":
	case "-":
	default:
		break
	}
	return false
}
