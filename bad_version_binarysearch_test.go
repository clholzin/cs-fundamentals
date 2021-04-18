package fundamentals

import (
	"testing"
)

// To execute Go code, please declare a func main() in a package "main"

// API, versions: 1, 2, 3, ... n
// regression - x
// good versions: 1, 2, ... x-1
// bad version: x, x+1, ...n
// goal - find first bad version
// provided - bool isBadVersion(int n) - true for >= 3
//  1 2 3 4 5
//.     x x x

func TestBSBadVersions(t *testing.T) {

	versions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(versions); i++ {

		globalIsBadVal = i

		if val := binarySearchVersions(len(versions)/2, (len(versions)/2)-1, versions); val != globalIsBadVal {
			t.Errorf("failed %d for globalIsBadVal %d", val, globalIsBadVal)
		}
		//t.Log("finish")

	}

	globalIsBadVal = 10

	if val := binarySearchVersions(len(versions)/2, (len(versions)/2)-1, versions); val != -1 {
		t.Errorf("failed %d for globalIsBadVal %d", val, globalIsBadVal)
	}
	//t.Log("finish")

}
