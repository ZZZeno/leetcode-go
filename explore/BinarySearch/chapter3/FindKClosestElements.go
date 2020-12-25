package chapter3

func FindClosestElements(arr []int, k int, x int) []int {
	return findClosestElements(arr, k, x)
}

func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[0:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:len(arr)]
	}
	if k == len(arr) {
		return arr
	}
	kindex := FindKIndex(arr, x)
	left := kindex
	right := kindex+1

	for right - left < k {
		if left == 0{
			// need return
			right += (k-(right-left))
			break
		}
		if right == len(arr) {
			left -= (k-(right-left))
			break
		}
		if arr[right]-x < x-arr[left-1] {
			right ++
		} else {
			left --
		}
	}
	return arr[left:right]
}

func FindKIndex(arr []int, x int) int {
	lo := 0
	hi := len(arr) - 1
	var mid int

	for lo < hi {
		mid = lo + (hi-lo)/2
		if arr[mid] == x {
			return mid
		}
		if hi - lo == 1 && arr[hi] != x && arr[lo] != x {
			if arr[hi] - x < x - arr[lo] {
				return hi
			} else {
				return lo
			}
		}
		if arr[mid] < x {
			lo = mid
		} else {
			hi = mid
		}
	}
	return mid
}
