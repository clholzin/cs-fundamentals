package fundamentals

func Bubble(unsorted []int) []int {
	for i := 0; i < len(unsorted); i++ {
		var temp int
		var length int = len(unsorted) - 1
		for j := 0; j < length; j++ {
			if unsorted[i] > unsorted[j+1] {
				temp = unsorted[j]
				unsorted[j] = unsorted[j+1]
				unsorted[j+1] = temp
			}
		}
	}
	return unsorted
}
