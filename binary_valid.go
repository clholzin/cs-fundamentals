package fundamentals

/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var prev interface{}

func isValidBST(root *TreeNode) bool {
	//return valid(root,nil,nil)
	prev = nil
	return inorderValid(root)
}

//Preorder traversal
func valid(root *TreeNode, low, high interface{}) bool {
	if root == nil {
		return true
	}
	if (low != nil && root.Val <= low.(int)) || (high != nil && root.Val >= high.(int)) {
		return false
	}

	return valid(root.Left, low, root.Val) && valid(root.Right, root.Val, high)
}

func inorderValid(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !inorderValid(root.Left) {
		return false
	}

	if prev != nil && root.Val <= prev.(int) {
		return false
	}
	prev = root.Val
	return inorderValid(root.Right)

}
