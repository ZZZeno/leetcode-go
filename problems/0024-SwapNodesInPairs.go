/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	cur := head
	index := 1
	var before *ListNode
	for cur != nil {
		arr1 := cur
		arr2 := cur.Next
        
		if arr2 == nil {
			break
		}
		after := cur.Next.Next
		index += 1
		if index % 2 == 0 {
			before = exch(before, after, arr1, arr2)
			cur = before.Next
			index += 1
		} else {
			cur = cur.Next
		}
	}
	return newHead
}

func exch(before, after, arr1, arr2 *ListNode) *ListNode {
	if before != nil {
		before.Next = arr2
	}
	arr1.Next = after
	arr2.Next = arr1
	return arr1
}
