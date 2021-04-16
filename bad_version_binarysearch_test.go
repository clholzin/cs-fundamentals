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

func TestBS(t *testing.T) {

	versions := []int{0, 1, 2, 3, 4, 5}

	for i := 0; i < len(versions)+1; i++ {

		globalIsBadVal = i

		if val := binarySearch(len(versions)/2, (len(versions)/2)-1, versions); val != globalIsBadVal {
			t.Errorf("failed %d", val)
		}

	}

	globalIsBadVal = 6

	if val := binarySearch(len(versions)/2, (len(versions)/2)-1, versions); val == -1 {
		t.Errorf("failed %d", val)
	}

}
