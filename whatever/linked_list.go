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

// 获取链表的中间节点，如果是奇数个节点返回中间节点，如果是偶数个返回中间两个节点的前面一个节点
func getMiddleNodeOfLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var fast, slow *ListNode
	fast = head
	slow = head
	for {
		fast = fast.Next.Next
		if fast == nil {
			return slow
		}
		if fast.Next == nil {
			return slow.Next
		}
		slow = slow.Next
	}
}

// 合并2个有序的单链表
func merge2SortedLinkedList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var newHead *ListNode
	var ptr1, ptr2, tail *ListNode
	if head1.Val < head2.Val {
		newHead = head1
		ptr1 = head1.Next
		ptr2 = head2
	} else {
		newHead = head2
		ptr1 = head1
		ptr2 = head2.Next
	}
	tail = newHead
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			tail.Next = ptr1
			tail = tail.Next
			ptr1 = ptr1.Next
		} else {
			tail.Next = ptr2
			tail = tail.Next
			ptr2 = ptr2.Next
		}
	}
	if ptr1 == nil {
		tail.Next = ptr2
	} else {
		tail.Next = ptr1
	}
	return newHead
}

