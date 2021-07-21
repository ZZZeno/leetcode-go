package utils

type Queue []*TreeNode

func (q *Queue) Pop() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:len(*q)]
	return ret
}

func (q *Queue) Push(node *TreeNode) {
	*q = append(*q, node)
}

