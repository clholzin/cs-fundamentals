package fundamentals

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int = 1
	var left int
	var right int

	if root.Left != nil {
		left = maxDepth(root.Left)
	}
	if root.Right != nil {
		right = maxDepth(root.Right)
	}
	if left > 0 && right < left {
		count += left
	} else if right > 0 && right >= left {
		count += right
	}

	return count
}
