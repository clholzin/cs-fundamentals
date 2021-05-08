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
		topotests{
			[][]int{{0, 1}, {1, 2}},
			3,
			true,
		},
		topotests{
			[][]int{{0, 1}, {1, 2}, {2, 0}},
			3,
			false,
		},
		topotests{
			[][]int{{0, 4}, {1, 4}, {3, 2}, {1, 3}},
			6,
			true,
		},
	}

	for _, test := range tabletest {

		if canFinish(test.Courses, test.PreReqs) != test.Expect {
			t.Fail()
		}

	}

}
