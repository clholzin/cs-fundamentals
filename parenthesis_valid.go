package fundamentals

// retur integer of the max depth of a set of parenthisis
func parenCount(st string) int {
	var count int
	var lastVal string
	var maxcount int
	for i := 0; i < len(st); i++ {
		val := string(st[i])
		if val == "(" {
			count++
			lastVal = val
		}
		if val == ")" && lastVal == "(" {
			maxcount = maximum(count, maxcount)
			count--
			lastVal = ""
		}
	}
	return maxcount
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/// writee a func for a string return bool if it has a valid set of parenthises
func parenCheck(st string) bool {
	var stack []string
	for i := 0; i < len(st); i++ {
		val := string(st[i])
		if val == "(" {
			stack = append(stack, val)
			continue
		}

		if val == ")" && len(stack) > 0 {
			lastindex := len(stack) - 1
			last := stack[lastindex]

			if last == "(" {
				stack = stack[:lastindex]
			}
		} else if val == ")" && len(stack) == 0 {
			return false
		}
	}
	return len(stack) == 0
}
