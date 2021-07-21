package problems

import (
	"fmt"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	root := IntArr2TreeNode([]int{5, 4, 6, 0, 0, 3, 7})
	root = IntArr2TreeNode([]int{2147483647,-2147483648})
	fmt.Println(isValidBST(root))
}
