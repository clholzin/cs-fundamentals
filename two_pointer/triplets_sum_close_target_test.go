package fundamentals

import "testing"

func TestTripletsCloseToTarget(t *testing.T) {

	triples := []int{-2, 0, 1, 2}
	sum := 3

	if searchTripletsCloseToTarget(triples, sum) != 1 {
		t.Fail()
	}

	triples = []int{-3, -1, 1, 2}
	sum = 5

	if searchTripletsCloseToTarget(triples, sum) != 0 {
		t.Fail()
	}
}
