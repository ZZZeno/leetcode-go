package chapter5

type ListNode struct {
    Val int
    Next *ListNode
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cur := &ListNode{
		Val:  0,
		Next: nil,
	}
	if l1.Val < l2.Val {
		cur = l1
		mergeWithCur(l1.Next, l2, cur)
	} else {
		cur = l2
		mergeWithCur(l1, l2.Next, cur)
	}
	return cur
}

func mergeWithCur(l1Tail *ListNode, l2Tail *ListNode, cur *ListNode) {
	if l1Tail == nil {
		cur.Next = l2Tail
		return
	}
	if l2Tail == nil {
		cur.Next = l1Tail
		return
	}
	if l1Tail.Val < l2Tail.Val {
		cur.Next = l1Tail
		mergeWithCur(l1Tail.Next, l2Tail, cur.Next)
	} else {
		cur.Next = l2Tail
		mergeWithCur(l1Tail, l2Tail.Next, cur.Next)
	}
}

