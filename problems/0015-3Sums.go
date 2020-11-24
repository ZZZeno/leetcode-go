package problems

import (
	"fmt"
	"sort"
)

func checkDup(arr [][]int, item []int) bool {
	for _, v := range arr {
		if v[0] == item[0] && v[1] == item[1] && v[2] == item[2] {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums) - 2; i ++ {
		lo := i + 1
		hi := len(nums) - 1
		for lo < hi {
			if nums[lo] + nums[hi] + nums[i] == 0{
				if !checkDup(res, []int{nums[i], nums[lo], nums[hi]}) {
					res = append(res, []int{nums[i], nums[lo], nums[hi]})
				}
				lo ++
				hi --
				continue
			}
			if nums[lo] + nums[hi] + nums[i] < 0 {
				lo ++
			} else if nums[lo] + nums[hi] + nums[i] > 0 {
				hi --
			}
		}
	}
	return res
}

func ThreeSum(nums []int) {
	fmt.Printf("%v\n", threeSum(nums))
}
