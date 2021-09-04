package patterns

import (
	"container/heap"
	"fmt"
	"strings"
)

// Pattern: Subsets//
//Unique Generalized Abbreviations (hard)
type AbbrWord struct {
	strings.Builder
	Count int
	Start int
}

func generateGeneralizedAbbreviation(word string) []string {
	wordlen := len(word)
	result := make([]string, 0)
	queue := make([]*AbbrWord, 1)
	queue[0] = &AbbrWord{Count: 0, Start: 0}
	for len(queue) > 0 {
		abword := queue[0]
		queue = queue[1:]
		if abword.Start == wordlen {
			if abword.Count != 0 {
				abword.WriteString(fmt.Sprintf("%d", abword.Count))
			}
			result = append(result, abword.String())
		} else {
			next := &AbbrWord{Start: abword.Start + 1, Count: abword.Count + 1}
			next.WriteString(abword.String())
			queue = append(queue, next)

			if abword.Count != 0 {
				abword.WriteString(fmt.Sprintf("%d", abword.Count))
			}

			fromStart := &AbbrWord{Start: abword.Start + 1, Count: 0}
			fromStart.WriteString(abword.String())
			fromStart.WriteByte(word[abword.Start])
			queue = append(queue, fromStart)

		}
	}
	return result
}

// Pattern Bitwise
// Two Single Numbers (medium)
func twoSingleNumbers(arr []int) []int {
	oneandtwo := 0
	for _, val := range arr {
		oneandtwo = oneandtwo ^ val
	}

	rightmostbit := 1
	for (rightmostbit & oneandtwo) == 0 {
		rightmostbit = rightmostbit << 1
	}

	one := 0
	two := 0
	for _, val := range arr {
		if rightmostbit&val != 0 {
			one ^= val
		} else {
			two ^= val
		}
	}
	return []int{one, two}
}

// Pattern: Top 'K' Elements
//Sum of Elements (medium)

func sumOfElements(arr []int, k1, k2 int) int {
	sum := 0
	maxheap := &MaxInts2{}
	heap.Init(maxheap)
	for i := 0; i < len(arr); i++ {
		if k2-1 > i {
			heap.Push(maxheap, arr[i])
		} else if maxheap.Peak().(int) > arr[i] {
			heap.Pop(maxheap)
			heap.Push(maxheap, arr[i])
		}
	}

	for i := 0; i < (k2 - k1 - 1); i++ {
		sum += heap.Pop(maxheap).(int)
	}

	return sum
}

// MaxInts
type MaxInts2 []int

func (max MaxInts2) Len() int           { return len(max) }
func (max MaxInts2) Swap(i, j int)      { max[i], max[j] = max[j], max[i] }
func (max MaxInts2) Less(i, j int) bool { return max[i] > max[j] }
func (max MaxInts2) Peak() interface{} {
	if max.Len() > 0 {
		return max[0]
	}
	return 0
}
func (max *MaxInts2) Pop() interface{} {
	tmp := *max
	length := tmp.Len()
	result := tmp[length-1]
	*max = tmp[:length-1]
	return result
}
func (max *MaxInts2) Push(val interface{}) {
	tmp := *max
	tmp = append(tmp, val.(int))
	*max = tmp
}
