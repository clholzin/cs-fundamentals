package fundamentals

/*
1,1,2,2,3,4,4,5,5

len(9)


BS pairs  - return missing pair

*/
func bs(values []int) int {

	split := len(values) / 2
	p1 := values[:split]
	p2 := values[split:]
	endp1, startp2 := p1[len(p1)-1], p2[0]

	//fmt.Printf("endp1: %d startp2: %d middle: %d\n", endp1, startp2, values[split])
	middle := values[split]
	if middle != endp1 && middle != startp2 {
		return middle
	}
	if endp1 == startp2 {
		p1 = append(p1, startp2)
		p2 = p2[1:]
	}
	if len(p2) == 1 {
		return p2[0]
	}
	if (len(p1))%2 == 1 {
		if len(p1) == 1 {
			return p1[0]
		}
		return bs(p1)
	}
	return bs(p2)

}

/* Notes
partion
p1 1,1,2,2
p2 3,4,4,5,5

p1 1,1 - 2,2

p1 1,1 - len(2) true and aretheypairs() - true
p1 2,2 - len(2) true and aretheypairs() -  true


p2 3,4  - 4,5,5

p2 3,4 - len(2) true and aretheypairs() - false - 3


1,1,2,3,3,4,4,5,5

p1 1,1,2,3

p1 1,1 - 2,3

p1 1,1 - len(2) true and aretheypairs() - true
p1 2,3 - return 2

p2 3,4,4,5,5



p1 1,1,2,   3,3,4,4,5,5
partion in middle but keep pair intact
p1 1,1,2 - odd

p1 1,1
p2 2  - return 2
*/

//1,1,2,3,3,4,4,5,5
//1,1,2,2,3,4,4,5,5
//1,1,2,2,3,3,4,5,5
