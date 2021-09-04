package fundamentals

func p2IsAlienSorted(words []stirng, order string) bool {

	sorted := true
	if len(words) == 0 {
		return true
	}

	orderList := []byte(order)
	orderMap := make(map[byte]byte)

	for idx, bchar := range orderList {
		orderMap[bchar] = byte(' ' + i)
	}

	prev := ""
	for i, word := range words {
		words[i] = p2Translate(word, orderMap)
	}

	for _, word := range words {
		if prev > word {
			return false
		}
		prev = word
	}

	return sorted

}

func p2Translate(word string, orderMap map[byte]byte) string {
	tmp := make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		tmp[i] = orderMap[word[i]]
	}
	return string(tmp)
}
