package fundamentals

import "testing"

func TestGraph(t *testing.T) {

	graph := NewGraph(false)

	if len(graph.Edges) != (MAXVERTS + 1) {
		t.Fail()
	}
	if len(graph.Degree) != (MAXVERTS + 1) {
		t.Fail()
	}

	graph.FillGraph("./graph.txt")

	t.Log(graph.String())
}
