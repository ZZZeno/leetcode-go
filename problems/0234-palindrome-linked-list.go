package problems

import "fmt"

func PrintLinkedList(head *ListNode)  {
	for {
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

// recursion
func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseLinkedList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// none recursion
func reverseLinkedList2(head *ListNode) *ListNode {
	var temp, newHead *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	for head != nil {
		temp = head
		head = head.Next
		temp.Next = newHead
		newHead = temp
	}
	return newHead
}

func linkedListLength(head *ListNode) int {
	length := 0
	ptr := head
	for {
		length += 1
		ptr = ptr.Next
		if ptr == nil {
			break
		}
	}
	return length
}

func isPalindrome(head *ListNode) bool {
	length := linkedListLength(head)
	if length == 1 {
		return true
	}
	if length == 2 {
		if head.Val == head.Next.Val {
			return true
		}
		return false
	}
	var (
		part2 *ListNode
		part1 *ListNode
		ptr = head
	)
	for i := 0; i < length/2-1; i ++ {
		ptr = ptr.Next
	}
	part1 = head
	if length % 2 == 0 {
		part2 = reverseLinkedList(ptr.Next)
		ptr.Next = nil
	} else {
		temp := ptr.Next.Next
		ptr.Next.Next = nil
		ptr.Next = nil
		part2 = reverseLinkedList(temp)
	}

	for i := 0; i < length/2; i ++ {
		if part1 == nil || part2 == nil {
			return true
		}
		if part1.Val != part2.Val {
			return false
		}
		part1 = part1.Next
		part2 = part2.Next
	}
	return true
}
