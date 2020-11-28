package chapter2

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	before := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return before
}