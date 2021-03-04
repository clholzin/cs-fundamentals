package fundamentals

type node struct {
	d int
	i int
}

type heap struct {
	size   int
	length int
	list   []*node
}

func (h *heap) sink(i int) {
	for 2*i+1 < h.length {
		j := 2*i + 1
		if j+1 < h.length && h.list[j].d < h.list[j+1].d {
			j++
		}

		if h.list[i].d > h.list[j].d {
			break
		}

		h.list[i], h.list[j] = h.list[j], h.list[i]
		i = j
	}
}

func (h *heap) swim(i int) {
	for i > 0 {
		pi := (i - 1) / 2
		if h.list[pi].d > h.list[i].d {
			break
		}

		h.list[pi], h.list[i] = h.list[i], h.list[pi]
		i = pi
	}
}

func (h *heap) max() int {
	if h.length > 0 {
		return h.list[0].d
	} else {
		return 0
	}
}

func kClosestMaxHeap(points [][]int, K int) [][]int {
	h := heap{K, 0, []*node{}}

	for i, p := range points {
		d := p[0]*p[0] + p[1]*p[1]

		if h.length < h.size {
			h.list = append(h.list, &node{d, i})
			h.length++
			h.swim(h.length - 1)
		} else if d < h.max() {
			h.list[0].d = d
			h.list[0].i = i
			h.sink(0)
		}
	}

	res := make([][]int, h.length)

	for i, v := range h.list {
		res[i] = points[v.i]
	}

	return res
}
