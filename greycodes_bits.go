package fundamentals

func grayCode(n int) []int {
	result := make([]int, 0)

	seqLen := 1 << n
	for i := 0; i < seqLen; i++ {
		num := i ^ i>>1
		result = append(result, num)
	}
	return result
}
