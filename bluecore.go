package fundamentals

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

// Example 2:

// Input: head = [1], n = 1
// Output: []

// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 {
		return head
	}
	order := make([]*ListNode, 0)

	tmp := head
	for tmp != nil {
		order = append(order, tmp)
		tmp = tmp.Next
	}

	// reconnect
	kthIndx := (len(order)) - n
	// 1 2 3 4 5  - len 5 -- 5 - 1 = 4
	// 5 - 5 = 0
	var left *ListNode

	if kthIndx == 0 {
		head = head.Next
	} else if kthIndx == (len(order) - 1) {
		left = order[kthIndx-1]
		left.Next = nil
	} else {
		left := kthIndx - 1  // what if i'm zero
		right := kthIndx + 1 // watch upper bound
		leftNode := order[left]
		leftNode.Next = order[right]
	}

	return head
}

/*
func main() {

	head := &ListNode{Val: 1}
	tmp := head
	for i := 2; i < 6; i++ {
		tmp.Next = &ListNode{Val: i}
		tmp = tmp.Next
	}

  	result := removeNthFromEnd(head, 2)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}

	fmt.Println()

	head = &ListNode{Val: 1}
	result = removeNthFromEnd(head, 1)
	if result == nil {
		fmt.Println([]int{})
	}

  	fmt.Println()

  	head = &ListNode{Val: 1}
	tmp = head
	for i := 2; i < 6; i++ {
		tmp.Next = &ListNode{Val: i}
		tmp = tmp.Next
	}

  	result = removeNthFromEnd(head, 1)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}

  	fmt.Println()

  	head = &ListNode{Val: 1}
	tmp = head
	for i := 2; i < 6; i++ {
		tmp.Next = &ListNode{Val: i}
		tmp = tmp.Next
	}

  	result = removeNthFromEnd(head, 5)
	for result != nil {
		fmt.Printf("%d ", result.Val)
		result = result.Next
	}
}*/
