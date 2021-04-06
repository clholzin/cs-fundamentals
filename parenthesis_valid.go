package fundamentals

import "fmt"

func main() {

	fmt.Println(parenCount("()"))                 //1
	fmt.Println(parenCount("(())"))               //2
	fmt.Println(parenCount("((()()()))"))         //3
	fmt.Println(parenCount("((()()(())))"))       //4
	fmt.Println(parenCount("((((((((()))))))))")) //9
	fmt.Println(parenCount("((()))(())"))         //3 - fail 4
	fmt.Println(parenCount("(())((()))"))         //3 - fail 4

	fmt.Println(parenCheck("()"))                //true
	fmt.Println(parenCheck("(d(ds))"))           // true
	fmt.Println(parenCheck("(d(ds)"))            // false
	fmt.Println(parenCheck("(d(dass(()()())))")) //true
	fmt.Println(parenCheck("(d(dass(()(())))"))  //false
	fmt.Println(parenCheck("(()))"))             //false
	fmt.Println(parenCheck("(((((((())))))))"))  //true
}

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
			maxcount = max(count, maxcount)
			count--
			lastVal = ""
		}
	}
	return maxcount
}

func max(a, b int) int {
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
