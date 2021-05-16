package fundamentals

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
