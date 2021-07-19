package fundamentals

import "testing"

func TestCountTripletsSum(t *testing.T) {

	triples := []int{-1, 0, 2, 3}
	sum := 3

	if searchTriplets(triples, sum) != 2 {
		t.Fail()
	}

	triples = []int{-1, 4, 2, 1, 3}
	sum = 5

	if searchTriplets(triples, sum) != 4 {
		t.Fail()
	}
}
