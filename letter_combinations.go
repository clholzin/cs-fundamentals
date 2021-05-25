package fundamentals

//https://leetcode.com/problems/letter-combinations-of-a-phone-number/

func letterCombinations(digits string) []string {
	lookup := make(map[byte][]string)
	lookup[2] = []string{"a", "b", "c"}
	lookup[3] = []string{"d", "e", "f"}
	lookup[4] = []string{"g", "h", "i"}
	lookup[5] = []string{"j", "k", "l"}
	lookup[6] = []string{"m", "n", "o"}
	lookup[7] = []string{"p", "q", "r", "s"}
	lookup[8] = []string{"t", "u", "v"}
	lookup[9] = []string{"w", "x", "y", "z"}

	var result []string = make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	result = permsVals("", digits, 0, lookup, result)
	return result
}

func permsVals(path string, digits string, index int, lookup map[byte][]string, result []string) []string {

	if len(path) == len(digits) {
		result = append(result, path)
		return result
	}
	dig := digits[index] - '0'
	values := lookup[dig]

	for i := 0; i < len(values); i++ {
		path += string(values[i])
		result = permsVals(path, digits, index+1, lookup, result)
		path = path[:len(path)-1]
	}
	return result
}
