package problems

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	min := nums[0]+nums[1]+nums[2]
	for i := 0; i < l-2; i++ {
		lo := i + 1
		hi := len(nums) - 1
		for lo < hi {
			base := nums[i] + nums[lo] + nums[hi]
			if absRange(base, target) < absRange(min, target) {
				min = base
			}
			if base == target {
				return base
			}
			if base < target {
				lo++
			} else {
				hi--
			}
		}
	}
	return min
}

func absRange(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func ThreeSumClosest(nums []int, target int) int {
	return threeSumClosest(nums, target)
}
