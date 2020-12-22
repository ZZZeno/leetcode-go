package chapter2

import "fmt"

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	for left < right {
		if mid > left {
			if nums[mid] < nums[mid-1] {
				return nums[mid]
			} else {
				if nums[mid] > nums[right] {
					left = mid
				} else {
					right = mid
				}
				mid = (left + right) / 2
			}
		} else {
			if nums[mid] < nums[right] {
				return nums[mid]
			} else {
				return nums[right]
			}
		}
	}
	return nums[mid]
}

func FindMin(nums []int) {
	fmt.Println(findMin(nums))
}
