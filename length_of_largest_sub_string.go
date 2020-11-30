package fundamentals

//abcabcbb
//bbbbbb
//dvdf
//pwwkew
//ckilbkd
//nfpdmpi
//anviaj

// https://leetcode.com/explore/interview/card/facebook/5/array-and-strings/3008/
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return int(0)
	}
	checker := map[string]int{}
	seq := 0
	count := 0
	for indx, tryit := range s {
		trystring := string(tryit)
		if previndx, ok := checker[trystring]; ok {
			for val, indy := range checker {
				if indy <= previndx {
					delete(checker, val)
					count--
				}
			}
		}
		checker[trystring] = indx
		count++
		if seq < count {
			seq = count
		}
	}
	return seq
}
