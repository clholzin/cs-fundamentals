package fundamentals

// accepted answer
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1map [26]int
	var s2map [26]int

	for i := 0; i < len(s1); i++ {
		s1map[s1[i]-'a']++
		s2map[s2[i]-'a']++
	}
	//fmt.Println(s1map,s2map)
	for i := 0; i < len(s2)-len(s1); i++ {
		if matches(s1map, s2map) {
			return true
		}
		s2map[s2[(i+len(s1))]-'a']++
		s2map[s2[i]-'a']--
	}
	//fmt.Println(s1map, s2map)

	return matches(s1map, s2map)
}

func matches(s1map, s2map [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1map[i] != s2map[i] {
			return false
		}
	}
	return true
}

/*
fails on
"abcdxabcde"
"abcdeabcdx"
*/

func checkInclusionSlidingWindow(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}
	subsets := findSubsets(s1)
	// fmt.Println("subsets",subsets)
	for _, sub := range subsets {
		if len(sub) == len(s2) {
			//fmt.Println(sub, s2)
			if sub == s2 {
				return true
			} else {
				continue
			}
		}
		start := 0
		end := len(sub)
		for i := 0; i < len(s2); i++ {
			val := s2[start:end]
			if sub == val {
				fmt.Println(sub, val)
				return true
			}
			start++
			end++
			if end > len(s2) {
				break
			}
		}
	}
	return false
}

// build all permutaions of a string
func findSubsets(s1 string) (result []string) {
	perms := []string{""}
	for i := 0; i < len(s1); i++ {
		sval := string(s1[i])
		permSize := len(perms)
		for j := 0; j < permSize; j++ {
			oldperm := perms[0]
			perms = perms[1:]
			if len(oldperm) == 0 {
				perms = append(perms, sval)
			}
			for p := 0; p < len(oldperm); p++ {
				oldp := oldperm
				o := oldp[:p]
				n := oldp[p:]
				newval := o + sval + n
				if len(newval) == len(s1) {
					result = append(result, newval)
				} else {
					perms = append(perms, newval)
				}
			}
		}
	}
	result = append(result, s1)
	return
}
