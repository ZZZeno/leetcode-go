package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	last := cur
	for cur.Next != nil {
		cur = cur.Next
		if last.Val == cur.Val {
			last.Next = cur.Next
		} else {
			last = cur
		}
	}
	return head
}

func DeleteDuplicates(head *ListNode)  {
	deleteDuplicates(head)
}
