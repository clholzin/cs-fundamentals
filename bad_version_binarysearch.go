package fundamentals

import "fmt"

var globalIsBadVal int

// notes
/* This style isn't partitioning the
array into chuncks but moving the index logarithmically
*/
func binarySearchVersions(indx, prevx int, versions []int) int {
	if indx >= len(versions) || indx < 0 {
		return -1
	}
	if indx == prevx {
		return -1
	}
	val := versions[indx]
	valcheck := isBadVersion(val)
	if indx == 0 && valcheck {
		return val
	}

	if indx == 0 && !valcheck || indx == len(versions)-1 && !valcheck {
		return -1
	}

	prev := versions[prevx]
	prevcheck := isBadVersion(prev)

	fmt.Printf("globalIsBadVal %d indx %d prev %d val %d\n", globalIsBadVal, indx, prevx, val)

	if valcheck && prevcheck {
		tmp := indx - 1
		prevx = tmp - 1
		fmt.Println(tmp, tmp-1)
		return binarySearchVersions(tmp, prevx, versions)
	}
	if !valcheck && !prevcheck {
		tmp := indx/2 + indx
		if indx == 1 {
			tmp = indx + 1
		} else if tmp == indx {
			tmp++
		} else if tmp >= len(versions) {
			tmp = len(versions) - 1
		}
		fmt.Println(tmp, tmp-1)
		prevx = tmp - 1
		return binarySearchVersions(tmp, prevx, versions)
	}

	return val
}

func isBadVersion(val int) bool {
	return val >= globalIsBadVal
}

/*
func search(versions []int, prevx, indx int) int {

	if indx >= len(versions) || indx < 0 {
		return -1
	}
	// check bounds for pervx
	val := versions[indx]
	prev := versions[prevx]
	valcheck := isBadVersion(val)
	prevcheck := !isBadVersion(prev)

	// both bad -1 indx && prev = indx
	// both good +1 indx && prev = indx

	if valcheck && !prevcheck {
		tmp := indx / 2
		//indx--
		prevx = indx
		return search(versions, prevx, indx)
	}
	if !valcheck && prevcheck {
		tmp := indx/2 + indx
		//indx++
		prevx = indx
		return search(versions, prevx, indx)
	}

	//if valcheck && prevcheck {
	return val
	//}

}

*/
