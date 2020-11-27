package chapter6

import (
	"fmt"
	"sort"
)

func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		nextIndex := nums[i] - 1
		for nums[nextIndex] != 0 {
			temp := nums[nextIndex] - 1
			nums[nextIndex] = 0
			nextIndex = temp
		}
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res = append(res, i+1)
		}
	}
	sort.Ints(res)
	return res
}

func FindDisappearedNumbers(nums []int) []int {
	return findDisappearedNumbers(nums)
}
