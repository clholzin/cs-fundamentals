package fundamentals

/**
source: https://www.interviewbit.com/problems/subtract/
 *
 * Definition for singly-linked list.*/

type LinkedList struct {
	Head *listNode
	Tail *listNode
}
type listNode struct {
	value int
	next  *listNode
}

func listNode_new(val int) *listNode {
	var node *listNode = new(listNode)
	node.value = val
	node.next = nil
	return node
}

func (h *LinkedList) Add(val int) {
	if h.Head == nil {
		h.Head = listNode_new(val)
		h.Tail = h.Head
		return
	}
	h.Tail.next = listNode_new(val)
	h.Tail = h.Tail.next
}

func (h *LinkedList) ToArray() (output []int) {
	var next *listNode = h.Head
	for next != nil {
		output = append(output, next.value)
		next = next.next
	}
	return output
}

/**
 * @input A : Head pointer of linked list
 *
 * @Output head pointer of list.
 */
var head *listNode
var total int
var memo []int
var iter bool

func subtract(A *listNode) *listNode {
	current := A
	if head == nil {
		head = current
	}

	if iter && len(memo) != 0 {
		current.value = memo[len(memo)-1] - current.value
		memo = memo[:len(memo)-1]
		return subtract(current.next)
	} else if iter && len(memo) == 0 {
		return head
	}

	memo = append(memo, current.value)

	if current.next == nil && !iter {
		total := len(memo)
		if total%2 == 1 {
			total++
		}
		memo = memo[total/2:]
		iter = true
		return subtract(head)
	}

	return subtract(current.next)
}
