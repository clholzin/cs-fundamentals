package fundamentals

/*
1,1,2,2,3,4,4,5,5

len(9)


BS pairs  - return missing pair

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
func bs(values []int) int {

	split := len(values) / 2
	p1 := values[:split]
	p2 := values[split:]

	if p1[len(p1)-1] == p2[len(p2)-1] {
		sp := len(p2) - 1
		tmp := p1[sp:]
		p1 = p1[sp:]
		p2 = append(tmp, p2)
		return bs(p1)
	}

	if len(p1)%2 == 1 {
		//odd
	}

	return bs(values)

}
