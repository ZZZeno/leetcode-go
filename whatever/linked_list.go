package whatever

import "fmt"

func PrintLinkedList(head *ListNode)  {
	for {
		if head == nil {
			break
		}
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
		if head == nil {
			break
		}
	}
	fmt.Println("nil")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// remove last N index Node
// 删除单链表里倒数第N个节点
func removeLastNIndexNode(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head.Next
	}
	var fast, slow *ListNode
	fast = head
	for i := 0; i < n; i ++ {
		if fast == nil || fast.Next == nil{
			return head
		}
		fast = fast.Next
	}
	slow = head
	var pre *ListNode
	pre = slow
	for fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = slow.Next
	slow = nil
	return head
}
