package problems

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	last := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i ++ {
		if nums[i] == last {
			cnt += 1
		} else {
			if cnt == 1 {
				return nums[i-1]
			}
			last = nums[i]
			cnt = 1
		}
	}
	return nums[len(nums)-1]
}

func SingleNumber() {
	fmt.Println(singleNumber([]int{4, 1, 2, 2, 1}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
