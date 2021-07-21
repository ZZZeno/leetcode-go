package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


type Queue []*TreeNode

func (q *Queue) Pop() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	ret := (*q)[0]
	*q = (*q)[1:len(*q)]
	return ret
}

func (q *Queue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func IntArr2TreeNode(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	queue := Queue{}
	var head = &TreeNode{Val: arr[0]}
	queue.Push(head)
	var queueHead *TreeNode
	for i := 1; i < len(arr); i++ {
		for {
			queueHead = queue.Pop()
			if queueHead != nil {
				break
			}
		}
		queueHead.Left = newNode(arr[i])
		queue.Push(queueHead.Left)
		i += 1
		if i >= len(arr) {
			break
		}
		queueHead.Right = newNode(arr[i])
		queue.Push(queueHead.Right)
	}
	return head
}

func newNode(i int) *TreeNode {
	if i == 0 {
		return nil
	}
	return &TreeNode{Val: i}
}
