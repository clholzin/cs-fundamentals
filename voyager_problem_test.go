package fundamentals

import "testing"

type topotests struct {
	PreReqs [][]int
	Courses int
	Expect  bool
}

func TestPrerequisites(t *testing.T) {

	tabletest := []topotests{
		topotests{
			[][]int{{0, 1}, {1, 0}},
			2,
			false,
		},
	}

	for _, test := range tabletest {

		if canFinish(test.Courses, test.PreReqs) != test.Expect {
			t.Fail()
		}

	}

}
