package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var MAX_INT = 1 << 32

func isValidBST(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}
	return validateBST(root.Left, -MAX_INT, root.Val) && validateBST(root.Right, root.Val, MAX_INT)
}

func validateBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if (root.Left != nil && root.Left.Val >= root.Val) || (root.Right != nil && root.Right.Val <= root.Val) {
		return false
	}
	return validateBST(root.Left, min, root.Val) && validateBST(root.Right, root.Val, max)
}
