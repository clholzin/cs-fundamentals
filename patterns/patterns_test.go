package patterns

import (
	"fmt"
	"testing"
)

type Data struct {
	ValueStr  string
	ValuesStr []string
	Value     []int
	ValValues [][]int
	Kth       int
	Expect    int
	ExpectStr string
	ExpectArr []int
}

func TestMedianOfAStream(t *testing.T) {
	moas := InitMedianOfAStream()
	moas.insertNum(3)
	moas.insertNum(1)
	if moas.findMedian() != 2.0 {
		t.Errorf("%f median not 2.0\n", moas.findMedian())
	}
	moas.insertNum(5)
	if moas.findMedian() != 3.0 {
		t.Errorf("%f median not 3.0\n", moas.findMedian())
	}
	moas.insertNum(4)
	if moas.findMedian() != 3.5 {
		t.Errorf("%f median not 3.5\n", moas.findMedian())
	}

}

func TestFindSingleNumbers(t *testing.T) {
	vals := []int{2, 3, 2, 4, 5, 4, 6, 6}
	fmt.Println(patternFindSingleNumbers(vals))
}

/*
Binary Search

Order-agnostic Binary Search (easy)
*/
func TestBSSearch(t *testing.T) {
	table := []Data{
		Data{
			Value:  []int{4, 6, 10},
			Kth:    10,
			Expect: 2,
		},
		Data{
			Value:  []int{1, 2, 3, 4, 5, 6, 7},
			Kth:    5,
			Expect: 4,
		},
		Data{
			Value:  []int{10, 6, 4},
			Kth:    10,
			Expect: 0,
		},
	}
	for i := 0; i < len(table); i++ {
		result := patternBSSearch(table[i].Value, table[i].Kth)
		if result != table[i].Expect {
			t.Fail()
		}
	}
}

func TestPatternKnapsack(t *testing.T) {

	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	capacity := 7
	result := patternKnapsack(profits, weights, capacity)
	if result != 22 {
		t.Fail()
	}
	capacity = 6
	result = patternKnapsack(profits, weights, capacity)
	if result != 17 {
		t.Fail()
	}

}

func TestPatternKnapsackFast(t *testing.T) {

	profits := []int{1, 6, 10, 16}
	weights := []int{1, 2, 3, 5}
	capacity := 7
	result := patternKnapsackFast(profits, weights, capacity)
	if result != 22 {
		t.Fail()
	}
	capacity = 6
	result = patternKnapsackFast(profits, weights, capacity)
	if result != 17 {
		t.Fail()
	}

}

func TestPatternTopologicalSort(t *testing.T) {
	table := []Data{
		Data{
			ValValues: [][]int{
				{3, 2},
				{3, 0},
				{2, 0},
				{2, 1},
			},
			Kth:       4,
			ExpectArr: []int{3, 2, 0, 1},
		},
		Data{
			ValValues: [][]int{
				{4, 2},
				{4, 3},
				{2, 0},
				{2, 1},
				{3, 1},
			},
			Kth:       5,
			ExpectArr: []int{4, 2, 3, 0, 1},
		},
	}
	for i := 0; i < len(table); i++ {
		result := patternTopologicalSort(table[i].Kth, table[i].ValValues)
		for j := 0; j < len(result); j++ {
			if result[j] != table[i].ExpectArr[j] {
				t.Fail()
			}
		}
	}
}

func TestPatternPrintAllTopoSorts(t *testing.T) {
	table := []Data{
		Data{
			ValValues: [][]int{
				{3, 2},
				{3, 0},
				{2, 0},
				{2, 1},
			},
			Kth: 4,
		},
		Data{
			ValValues: [][]int{
				{2, 5},
				{0, 5},
				{0, 4},
				{1, 4},
				{3, 2},
				{1, 3},
			},
			Kth: 6,
		},
	}
	for i := 0; i < len(table); i++ {
		patternAllTaskSchedulingOrders(table[i].Kth, table[i].ValValues)
	}
}

func TestPatternAlienDictionaryFindOrder(t *testing.T) {
	table := []Data{
		Data{
			ValuesStr: []string{"ba", "bc", "ac", "cab"},
			ExpectStr: "bac",
		},
		Data{
			ValuesStr: []string{"cab", "aaa", "aab"},
			ExpectStr: "cab",
		},
		Data{
			ValuesStr: []string{"ywx", "wz", "xww", "xz", "zyy", "zwz"},
			ExpectStr: "ywxz",
		},
	}
	for i := 0; i < len(table); i++ {
		result := patternAlienDictionaryFindOrder(table[i].ValuesStr)
		if result != table[i].ExpectStr {
			t.Fail()
		}
	}
}

func TestFindKthSmallestNumberFast(t *testing.T) {

	table := []Data{
		Data{
			Value:  []int{1, 5, 12, 2, 11, 5},
			Kth:    3,
			Expect: 5,
		},
		Data{
			Value:  []int{1, 5, 12, 2, 11, 5},
			Kth:    4,
			Expect: 5,
		},
		Data{
			Value:  []int{5, 12, 11, -1, 12},
			Kth:    3,
			Expect: 11,
		},
	}

	for i := 0; i < len(table); i++ {
		result := findKthSmallestNumberFast(table[i].Value, table[i].Kth)
		if result != table[i].Expect {
			t.Errorf("failed: %d expect: %d", result, table[i].Expect)
		}
	}
}

func TestFindKthSmallestNumberHeap(t *testing.T) {

	table := []Data{
		Data{
			Value:  []int{1, 5, 12, 2, 11, 5},
			Kth:    3,
			Expect: 5,
		},
		Data{
			Value:  []int{1, 5, 12, 2, 11, 5},
			Kth:    4,
			Expect: 5,
		},
		Data{
			Value:  []int{5, 12, 11, -1, 12},
			Kth:    3,
			Expect: 11,
		},
	}

	for i := 0; i < len(table); i++ {
		result := findKthSmallestNumberHeap(table[i].Value, table[i].Kth)
		if result != table[i].Expect {
			t.Errorf("failed: %d expect: %d", result, table[i].Expect)
		}
	}
}
