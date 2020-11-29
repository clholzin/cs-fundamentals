package fundamentals

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	if l1 == nil && l2 == nil {
		var tp *ListNode
		return tp
	}
	var l1val, l2val int
	if l2 == nil && l1 != nil {
		temp.Val = l1.Val
		GetNextVal(temp, l1.Next, l2)
	} else if l1 == nil && l2 != nil {
		temp.Val = l2.Val
		GetNextVal(temp, l1, l2.Next)
	} else {
		l1val = l1.Val
		l2val = l2.Val
		if l1val <= l2val {
			temp.Val = l1val
			GetNextVal(temp, l1.Next, l2)
		} else if l1val > l2val {
			temp.Val = l2val
			GetNextVal(temp, l1, l2.Next)
		}
	}

	return temp
}

func GetNextVal(sorted *ListNode, l1, l2 *ListNode) {
	var temp *ListNode
	if l1 == nil && l2 == nil {
		return
	}
	if sorted != nil && sorted.Next == nil {
		sorted.Next = &ListNode{}
		temp = sorted.Next
	}
	if l1 != nil && l2 == nil {
		temp.Val = l1.Val
		GetNextVal(temp, l1.Next, l2)
		return
	} else if l2 != nil && l1 == nil {
		temp.Val = l2.Val
		GetNextVal(temp, l1, l2.Next)
		return
	}
	l1val := l1.Val
	l2val := l2.Val
	if l1val <= l2val {
		temp.Val = l1val
		GetNextVal(temp, l1.Next, l2)
	} else if l1val > l2val {
		temp.Val = l2val
		GetNextVal(temp, l1, l2.Next)
	}
}
