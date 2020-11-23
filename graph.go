package fundamentals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Adjacency Lists:
   travers Θ(m + n)
   no small counts (m + n)
   Faster to find the degree of a vertex
   Better for most problems?


Read a graph from a file.
A typical graph format consists of an initial line featuring the number of vertices and edges in the graph, followed by a listing of the edges at one vertex pair per line.
*/
const (
	MAXVERTS int = 1000
)

type Edgenode struct {
	Y      int
	Weight int
	Next   *Edgenode
}

type Graph struct {
	Edges     []*Edgenode
	Degree    []int
	NVertices int
	NEdges    int
	Directed  bool
}

func NewGraph(directed bool) *Graph {
	g := &Graph{}
	g.NVertices = 0
	g.NEdges = 0
	g.Directed = directed
	g.Edges = make([]*Edgenode, MAXVERTS+1)
	g.Degree = make([]int, MAXVERTS+1)
	return g
}

func (g *Graph) FillGraph(filepath string) {
	file, err := os.Open(filepath)
	read := bufio.NewScanner(file)

	for read.Scan() {
		if err := read.Err(); err != nil {
			fmt.Printf("failed to read graph with error: %v", err)
			break
		}
		data := strings.Split(read.Text(), ",")
		if len(data) == 1 {
			vert, err := strconv.Atoi(data[0])
			must(err)
			g.NVertices = vert
		}

		if len(data) == 2 {
			vert, err := strconv.Atoi(data[0])
			must(err)
			x := vert
			edge, err := strconv.Atoi(data[1])
			must(err)
			y := edge
			g.Insert(x, y, g.Directed)
		}

	}
}

func (g *Graph) Insert(x, y int, directed bool) {
	next := &Edgenode{Y: y}
	next.Next = g.Edges[x]
	g.Edges[x] = next
	g.Degree[x]++
	if directed == false {
		g.Insert(y, x, true)
	} else {
		g.NEdges++
	}

}

func (g *Graph) setup_search() (processed, discovered []bool, parents []int) {

	parents = make([]int, g.NVertices)
	processed = make([]bool, g.NVertices)
	discovered = make([]bool, g.NVertices)

	for i := 0; i < g.NVertices; i++ {
		parents[i], processed[i], discovered[i] = -1, false, false
	}
	return processed, discovered, parents
}

func (g *Graph) BFS(start int) *Edgenode {

	var found *Edgenode
	processed, discovered, parents := g.setup_search()
	queue := NewQueue(NewQueueNode(start))
	var current *Edgenode
	if start < len(discovered) {
		discovered[start] = true
	}

	for queue.IsEmpty() != false {
		v, err := queue.Remove()
		if err != nil {
			fmt.Println(err)
			continue
		}
		g.process_vertex_early(v)
		processed[v] = true
		current = g.Edges[v]
		for current != nil {
			y := current.Y
			if processed[y] == false || g.Directed {
				g.process_edge(v, y)
			}
			if discovered[y] == false {
				queue.Add(y)
				discovered[y] = true
				parents[y] = v
			}
			current = current.Next
		}
		g.process_vertex_late(v)
	}

}

func (g *Graph) process_vertex_late(v int) bool {
	return false
}
func (g *Graph) process_vertex_early(v int) {
	fmt.Printf("processed vertex %d\n", v)
}

func (g *Graph) process_edge(x, y int) {
	fmt.Printf("processed edge x:%d - y:%d", x, y)
}

func (g *Graph) String() string {
	buf := make([]byte, 0)
	buffer := NewBuffer(buf)
	var current *Edgenode
	buffer.WriteString(fmt.Sprintf("Verticies: %d\n", g.NVertices))
	for i := 0; i < g.NVertecies; i++ {
		current = g.Edges[i]
		for current != nil {
			buffer.WriteString(fmt.Sprintf("Vert: %d Edge: %d\n", i, current.Y))
		}
	}
	return buffer.String()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
