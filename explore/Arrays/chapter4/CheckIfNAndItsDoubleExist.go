package chapter4

import "sort"

func checkIfExist(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	sort.Ints(arr)
	for i := 0; i < len(arr); i ++ {
		//res := findBinary(arr, 2*arr[i], i, len(arr))
		res := findLinear(arr, 2*arr[i])
		if res {
			return true
		}
	}
	return false
}

func findBinary(arr []int, targetValue, lo, hi int) bool {
	mid := (lo+hi)/2
	for lo < mid {
		if arr[mid] == targetValue {
			return true
		}
		if arr[mid] < targetValue {
			lo = mid
		} else {
			hi = mid
		}
		mid = (lo+hi)/2
	}
	return false
}

func findLinear(arr []int, targetValue int) bool {
	for i, _ := range arr {
		if arr[i] == targetValue{
			if arr[i] == 0 {
				if (i > 0 && arr[i-1] == 0) || (i < len(arr)-1 && arr[i+1] == 0) {
					return true
				}
				return false
			}
			return true
		}
	}
	return false
}

func CheckIfExist(arr []int) bool {
	return checkIfExist(arr)
}
