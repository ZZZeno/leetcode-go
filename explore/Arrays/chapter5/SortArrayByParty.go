package chapter5

import "fmt"

func sortArrayByParity(A []int) []int {
	if len(A) == 1 || len(A) == 0 {
		return A
	}
	lo := 0
	hi := len(A) - 1
	for lo <= hi {
		if A[lo]%2 == 0 {
			lo++
		}
		if A[hi]%2 == 1 {
			hi--
		}
		if lo <= hi {
			exch(A, lo, hi)
		}
	}
	return A
}

func SortArrayByParity() {
	a := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(a))
}
