package problems

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (q == nil && p != nil) || (p == nil && q != nil) {
		return false
	}
	tmpP := make([]int, 0)
	tmpQ := make([]int, 0)
	tranverseTreeLeftFirstOrder(p, &tmpP)
	tranverseTreeLeftFirstOrder(q, &tmpQ)
	if !arrayEq(tmpP, tmpQ) {
		return false
	}
	index := 0
	index2 := 0
	tranverseTreeParentFirstOrder(p, tmpP, &index)
	tranverseTreeParentFirstOrder(q, tmpQ, &index2)
	if !arrayEq(tmpP, tmpQ) {
		return false
	}
	return true
}

func arrayEq(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	l := len(arr1)
	for i := 0; i < l; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func tranverseTreeLeftFirstOrder(p *TreeNode, res *[]int)  {
	if p.Left == nil && p.Right != nil {
		*res = append(*res, -1)
	}
	if p.Left != nil {
		tranverseTreeLeftFirstOrder(p.Left, res)
	}
	if p.Right != nil {
		tranverseTreeLeftFirstOrder(p.Right, res)
	}
	*res = append(*res, p.Val)
}
func tranverseTreeParentFirstOrder(p *TreeNode, res []int, index *int) {
	//*res = append(*res, p.Val)
	res[*index] = p.Val
	*index += 1
	if p.Left != nil {
		tranverseTreeParentFirstOrder(p.Left, res, index)
	}
	if p.Right != nil {
		tranverseTreeParentFirstOrder(p.Right, res, index)
	}
	if p.Left == nil && p.Right != nil {
		res[*index] = -1
	}
}

func TestOutputTree() {
	a := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil},
	}
	res := make([]int, 3)
	index := 0
	tranverseTreeParentFirstOrder(&a, res, &index)
	fmt.Println(res)
}

func IsSameTree(p , q *TreeNode)  {
	fmt.Println(isSameTree(p, q))
}
