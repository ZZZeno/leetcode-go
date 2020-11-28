package Queues

type MyCircularQueue struct {
	items []int
	head  int
	tail  int
	size  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		items: make([]int, k),
		head:  -1,
		tail:  -1,
		size: 0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = (this.head+1) % len(this.items)
	}
	this.tail = (this.tail+1) % len(this.items)
	this.items[this.tail] = value
	this.size += 1
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head+1) % len(this.items)
	this.size -= 1
	if this.IsEmpty() {
		this.tail = this.head
	}
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.items[this.tail]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.items)
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
