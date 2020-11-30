package fundamentals

const IntMax int = int(2147483647)
const MaxNext int = int(IntMax % 10)

var IntMap map[string]int = map[string]int{
	"0": int(0),
	"1": int(1),
	"2": int(2),
	"3": int(3),
	"4": int(4),
	"5": int(5),
	"6": int(6),
	"7": int(7),
	"8": int(8),
	"9": int(9),
}

func myAtoi(s string) int {
	var temp int
	tempIntsStr := ""
	prevSt := ""

	neg := "-"
	plus := "+"
	whitespace := " "

	checkpos := false
	checkneg := false
	checkfirstChar := false
	isGreaterThanMax := false

	for _, str := range s {
		st := string(str)

		_, prOk := IntMap[prevSt]
		val, ok := IntMap[st]
		if prOk && !ok {
			break
		}
		if st == neg {
			checkneg = true
		}
		if st == plus {
			checkpos = true
		}

		if checkpos && checkneg && !ok && !prOk {
			return 0
		}

		if ok {
			tempIntsStr += st
			if temp > ((IntMax - MaxNext) / 10) {
				temp = IntMax
				isGreaterThanMax = true
				break
			}
			if temp == ((IntMax-MaxNext)/10) && val > MaxNext {
				temp = IntMax
				isGreaterThanMax = true
				break
			}
			if (IntMax - temp) >= val {
				temp = temp * 10
				temp += val
			}
		} else if len(tempIntsStr) == 0 {
			if prevSt == plus || prevSt == neg {
				return 0
			} else if st != whitespace && st != plus && st != neg && !checkfirstChar {
				return 0
			}
		} else if len(tempIntsStr) > 0 {
			break
		}
		prevSt = st
	}

	if checkneg {
		if isGreaterThanMax {
			temp++
		}
		temp = temp * -1
	}
	return temp
}
