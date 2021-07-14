package whatever

import (
	"container/heap"
	"fmt"
)

type IntLittleTopHeap []int

func (h IntLittleTopHeap) Len() int {
	return len(h)
}

func (h IntLittleTopHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntLittleTopHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntLittleTopHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntLittleTopHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntBigTopHeap []int

func (h IntBigTopHeap) Len() int {
	return len(h)
}

func (h IntBigTopHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntBigTopHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntBigTopHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntBigTopHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
