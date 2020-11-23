package fundamentals

import "testing"

func TestGraph(t *testing.T) {

	graph := NewGraph(false)

	if len(graph.Edges) != (MaxVerts + 1) {
		t.Fail()
	}
	if len(graph.Degree) != (MaxVerts + 1) {
		t.Fail()
	}

	graph.Fill("./graph.txt")
}
