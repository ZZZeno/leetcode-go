package problems

func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		res = [][]int{}
		path = []int{}
	)
	pathSumBT(&res, path, root, targetSum)
	return res
}

func pathSumBT(res *[][]int, path []int, root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	if isLeaf(*root) {
		if targetSum == root.Val {
			path = append(path, root.Val)
			var temp = make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
			return
		}
		return
	}
	path = append(path, root.Val)
	pathSumBT(res, path, root.Left, targetSum-root.Val)
	pathSumBT(res, path, root.Right, targetSum-root.Val)
}

func isLeaf(node TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
