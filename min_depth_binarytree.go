package fundamentals

/**
 * Definition for a binary tree node.
 */
type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

const UintSize = 32 << (^uint(0) >> 32 & 1)
const MaxInt = 1<<(UintSize-1) - 1

func minDepth(root *TreeNode3) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var count int = int(MaxInt)
	if root.Left != nil {
		if ldepth := minDepth(root.Left); count > ldepth {
			count = ldepth
		}
	}
	if root.Right != nil {
		if rdepth := minDepth(root.Right); count > rdepth {
			count = rdepth
		}
	}
	count++
	return count
}
