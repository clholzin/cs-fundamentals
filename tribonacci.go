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
	/*  if v, ok := memo[n-1]; ok {
	        t2 = v
	    }else{
	        val := tribonacci(n-1)
	        memo[n-1] = val
	    }
	    if v, ok := memo[n-2]; ok {
	        t3 = v
	    }
	    if v, ok := memo[n-3]; ok {
	        t4 = v
	    }
	    return  tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3);*/
}
