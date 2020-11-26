package chapter4

import "fmt"

func validMountainArray(arr []int) bool {
	max := arr[0]
	maxIndex := 0
	for i, v := range arr {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	if maxIndex == 0 || maxIndex == len(arr) - 1 {
		return false
	}
	return isMonotonic(arr, 0, maxIndex) && isMonotonic(arr, maxIndex, len(arr))
}

func isMonotonic(arr []int, start, end int) bool {
	if end - start == 0 {
		return false
	}
	if end == len(arr) {
		if end - start == 2 {
			return arr[end-1] < arr[start]
		}
		for i := start; i < end-1; i ++ {
			if arr[i] <= arr[i+1] {
				return false
			}
		}
	} else {
		if end - start == 1 {
			return arr[end] > arr[start]
		}
		for i := 0; i < end; i ++ {
			if arr[i] >= arr[i+1] {
				return false
			}
		}
	}
	return true
}

func ValidMountainArray()  {
	fmt.Println(validMountainArray([]int{3, 5, 5}))
}
