package chapter2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
	if root.Val > val {
		if root.Left == nil {
			return nil
		}
		return searchBST(root.Left, val)
	} else {
		if root.Right == nil {
			return nil
		}
		return searchBST(root.Right, val)
	}
}
