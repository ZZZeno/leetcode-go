package problems

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	min := nums[0] + nums[1] + nums[2]
	for i := 0; i < l; i++ {
		for j := i+1; j < l; j ++ {
			for k := j+1; k < l; k ++ {
				sum := nums[i] + nums[j] + nums[k]
				if absRange(sum, target) < absRange(min, target) {
					min = sum
				}
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

func ThreeSumClosest()  {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
