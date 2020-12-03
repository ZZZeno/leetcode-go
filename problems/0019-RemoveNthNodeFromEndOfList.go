package problems


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	arr := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	amount := len(arr)
	index := amount - n
	if index == 0 {
		return head.Next
	} else if index == amount-1 {
		arr[amount-2].Next = nil
	} else {
		arr[index-1].Next = arr[index].Next
	}
	return head
}
