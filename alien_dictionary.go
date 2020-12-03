package fundamentals

func isAlienSorted(words []string, order string) bool {
	orderlist := []byte(order)
	orderL := map[byte]int{}
	for i, bchar := range orderlist {
		orderL[bchar] = i
	}
	wordbreak := []byte("*")
	orderL[wordbreak[0]] = -1
	var isordered bool
	var prevword []byte = []byte(words[0])
	words = words[1:]
	var equal bool
	for _, word := range words {
		current := []byte(word)
		shorter := prevword
		bigger := current
		var swapped bool

		if len(prevword) > len(current) {
			shorter = current
			bigger = prevword
			swapped = true
		}
		for j, by := range shorter {
			preorder := orderL[by]
			postorder := orderL[bigger[j]]
			fmt.Println(preorder, postorder)
			if preorder == postorder && !equal {
				equal = true
			} else {
				equal = false
			}
			if swapped {
				if preorder != postorder && preorder > postorder {
					return true
				}
			} else {
				if preorder != postorder && preorder < postorder {
					return true
				} else if preorder != postorder && preorder > postorder {
					return false
				}
			}

		}
		if equal {
			return true
		}
		prevword = current
	}
	return isordered
}
