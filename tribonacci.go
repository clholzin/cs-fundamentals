package fundamentals

func tribonacci(n int) int {
	memo := make(map[int]int)
	return tribonaccimemo(n, memo)
}

func tribonaccimemo(n int, memo map[int]int) int {
	var t int
	if n == 0 {
		return t
	}
	if n <= 2 {
		return (t + 1)
	}

	for _, v := range []int{1, 2, 3} {
		check := n - v
		if vv, ok := memo[check]; ok {
			t += vv
		} else {
			val := tribonaccimemo(check, memo)
			memo[check] = val
			t += val
		}
	}
	return t
}
