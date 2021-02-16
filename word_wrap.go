package fundamentals

func wraplines(words []string, maxlen int) (out []string) {

	var currow int
	out = append(out, "")
	for i := 0; i < len(words); i++ {
		amountleft := maxlen - len(out[currow])
		nextword := words[i]
		if amountleft < len(nextword) {
			currow++
			out = append(out, "")
		}
		resetrow, _ := addchars(out[currow], nextword, maxlen)
		out[currow] = resetrow
		if i+1 < len(words) {
			nextword = words[i+1]
			if (len(out[currow]) + len(nextword)) < maxlen {
				out[currow] += "-"
			}
		}

	}

	return out
}

func addchars(out string, word string, maxlen int) (string, bool) {

	if len(out) == maxlen {
		return out, len(out) == maxlen
	}
	if len(word) == 0 {
		return out, false
	}

	out += string(word[0])
	word = word[1:]

	return addchars(out, word, maxlen)
}
