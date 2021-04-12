package fundamentals

// To execute Go code, please declare a func main() in a package "main"

// API, versions: 1, 2, 3, ... n
// regression - x
// good versions: 1, 2, ... x-1
// bad version: x, x+1, ...n
// goal - find first bad version
// provided - bool isBadVersion(int n) - true for >= 3
//  1 2 3 4 5
//.     x x x

import "fmt"

/*
func main() {
  fmt.Println(search(versions,len(versions)/2))
}*/

func search(versions []int, prevx, indx int) int {

	if indx >= len(versions) || indx < 0 {
		return -1
	}
	// check bounds for pervx
	val := versions[indx]
	prev := version[prevx]
	valcheck := isBadVersion(val)
	prevcheck := !isBadVersion(prev)

	// both bad -1 indx && prev = indx
	// both good +1 indx && prev = indx

	if valcheck && !prevcheck {
		tmp := indx
		indx--
		prevx = tmp
		return search(versions, prevx, indx)
	}
	if !valcheck && prevcheck {
		tmp := indx
		indx++
		prevx = tmp
		return search(versions, prevx, indx)
	}

	//if valcheck && prevcheck {
	return val
	//}

}
