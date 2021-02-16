package fundamentals

/*
Suppose we have some input data describing a graph of relationships between parents and children over multiple generations. The data is formatted as a list of (parent, child) pairs, where each individual is assigned a unique positive integer identifier.

For example, in this diagram, 6 and 8 have common ancestors of 4 and 14.

             15
             |
         14  13
         |   |
1   2    4   12
 \ /   / | \ /
  3   5  8  9
   \ / \     \
    6   7     11

parentChildPairs1 = [
    (1, 3), (2, 3), (3, 6), (5, 6), (5, 7), (4, 5),
    (4, 8), (4, 9), (9, 11), (14, 4), (13, 12),
    (12, 9),(15, 13)
]

Write a function that takes the graph, as well as two of the individuals in our dataset, as its inputs and returns true if and only if they share at least one ancestor.

Sample input and output:

hasCommonAncestor(parentChildPairs1, 3, 8) => false
hasCommonAncestor(parentChildPairs1, 5, 8) => true
hasCommonAncestor(parentChildPairs1, 6, 8) => true
hasCommonAncestor(parentChildPairs1, 6, 9) => true
hasCommonAncestor(parentChildPairs1, 1, 3) => false
hasCommonAncestor(parentChildPairs1, 3, 1) => false
hasCommonAncestor(parentChildPairs1, 7, 11) => true
hasCommonAncestor(parentChildPairs1, 6, 5) => true
hasCommonAncestor(parentChildPairs1, 5, 6) => true

Additional example: In this diagram, 4 and 12 have a common ancestor of 11.

        11
       /  \
      10   12
     /  \
1   2    5
 \ /    / \
  3    6   7
   \        \
    4        8

parentChildPairs2 = [
    (1, 3), (11, 10), (11, 12), (2, 3), (10, 2),
    (10, 5), (3, 4), (5, 6), (5, 7), (7, 8),
]

hasCommonAncestor(parentChildPairs2, 4, 12) => true
hasCommonAncestor(parentChildPairs2, 1, 6) => false
hasCommonAncestor(parentChildPairs2, 1, 12) => false

n: number of pairs in the input


*/

import "fmt"

func main() {
	parentChildPairs1 := [][]int{
		[]int{1, 3},
		[]int{2, 3},
		[]int{3, 6},
		[]int{5, 6},
		[]int{5, 7},
		[]int{4, 5},
		[]int{4, 8},
		[]int{4, 9},
		[]int{9, 11},
		[]int{14, 4},
		[]int{13, 12},
		[]int{12, 9},
		[]int{15, 13},
	}

	parentChildPairs2 := [][]int{
		[]int{1, 3},
		[]int{11, 10},
		[]int{11, 12},
		[]int{2, 3},
		[]int{10, 2},
		[]int{10, 5},
		[]int{3, 4},
		[]int{5, 6},
		[]int{5, 7},
		[]int{7, 8},
	}

}

/*

func hasCommonAncestor(pairs [][]int,pair1,pair2 int) bool {
  parents1 := make([]int,0)
  for
  recur(parentChildPairs1,pair1,parents1,make(map[int]int))

}
func recur(c map[int]int, value int,savedpr []int,memo map[int]int)  {
  if _,ok := memo[value]; ok {
    return
  }
  if v, ok := c[value]; ok {
    parents1 = append(parents1,v)
    recur(parentChildPairs1,pair1,parents1,memo)
  }
}
*/
func findNodesWithZeroAndOneParents(parentChildPairs [][]int) ([]int, []int) {
	parents := make(map[int]int)
	justoneparent := make(map[int]int)
	for _, v := range parentChildPairs {
		parents[v[0]] = 0
		if _, ok := justoneparent[v[1]]; ok {
			justoneparent[v[1]] = justoneparent[v[1]] + 1
		} else {
			justoneparent[v[1]] = 0
		}

	}
	var childrenwithoneparent []int
	for _, v := range parentChildPairs {
		if _, ok := parents[v[1]]; ok {
			parents[v[1]] += 1
		}
		if val, ok := justoneparent[v[1]]; ok && val == 0 {
			childrenwithoneparent = append(childrenwithoneparent, v[1])
		}

	}
	var parentschildrenofnoone []int
	preventdups := make(map[int]int)
	for _, v := range parentChildPairs {
		if _, ok := preventdups[v[0]]; !ok {
			if count, ok := parents[v[0]]; ok && count == 0 {
				parentschildrenofnoone = append(parentschildrenofnoone, v[0])
			}
		}
		preventdups[v[0]] = 0
	}
	return parentschildrenofnoone, childrenwithoneparent
}
