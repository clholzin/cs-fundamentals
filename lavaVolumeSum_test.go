package fundamentals

import (
	"fmt"
	"testing"
)

func TestLavaVolumeSum(t *testing.T) {
	vals := []int{1, 0, 0, 1, 2, 1, 0, 2}
	fmt.Println(lavaVolumeSum(vals))
}
