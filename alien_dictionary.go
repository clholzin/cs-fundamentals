package fundamentals

func isAlienSorted(words []string, order string) bool {
	orderlist := []byte(order)
	orderL := map[byte]byte{}
	for i, bchar := range orderlist {
		orderL[bchar] = byte(' ' + i)
	}

	var isordered bool = true
	for i, word := range words {
		words[i] = translate(word, orderL)
	}

	last := ""
	for _, word := range words {
		if last > word {
			return false
		}
		last = word
	}

	return isordered
}

func translate(word string, order map[byte]byte) string {
	temp := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		temp[i] = order[word[i]]
	}
	return string(temp)
}
