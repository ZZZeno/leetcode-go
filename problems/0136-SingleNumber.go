package problems

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    last := nums[0]
    for i := 1; i < len(nums); i ++ {
        last = last ^ nums[i]
    }
    return last
}

func SingleNumber() {
	fmt.Println(singleNumber([]int{4, 1, 2, 2, 1}))
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
