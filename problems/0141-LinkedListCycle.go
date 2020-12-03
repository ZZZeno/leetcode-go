package problems

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	next1 := head
	next2 := head.Next

	for next2 != nil {
		if next1 == next2 {
			return true
		}
		next1 = next1.Next
		if next2.Next == nil {
			return false
		} else {
			next2 = next2.Next.Next
		}
	}
	return false
}
