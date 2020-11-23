package problems

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Printf("%v \n", CombinationSum(candidates, target))
}

func TestMySqrt(t *testing.T) {
	MySqrt(8)
}

func TestClimbStairs(t *testing.T) {
	ClimbStairs(3)
}

func TestDeleteDuplicates(t *testing.T) {
	tmp2 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	tmp := ListNode{
		Val:  1,
		Next: &tmp2,
	}
	head := ListNode{
		Val:  1,
		Next: &tmp,
	}
	cur := head
	fmt.Printf("%d", cur.Val)
	for cur.Next != nil {
		fmt.Printf("->%d", cur.Next.Val)
		cur = *cur.Next
	}
	fmt.Println()
	DeleteDuplicates(&head)

	cur = head
	fmt.Printf("%d", cur.Val)
	for cur.Next != nil {
		fmt.Printf("->%d", cur.Next.Val)
		cur = *cur.Next
	}
	fmt.Println()

}
