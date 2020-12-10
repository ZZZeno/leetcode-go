package problems

func getMinimumDifference(root *TreeNode) int {
	arr := make([]int, 0)
	midInterval(root, &arr)
	min := arr[1] - arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] - arr[i] < min {
			min = arr[i+1] - arr[i]
		}
	}
	return min
}

func midInterval(root *TreeNode, arr *[]int)  {
	if root == nil {
		return
	}
	midInterval(root.Left, arr)
	*arr = append(*arr, root.Val)
	midInterval(root.Right, arr)
}
