package problems

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		lo := i + 1
		hi := len(nums) - 1
		for lo < hi {
			if nums[lo]+nums[hi]+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
				for hi > lo && hi+1 < len(nums) && nums[hi] == nums[hi+1] {
					hi--
				}
				continue
			}
			if nums[lo]+nums[hi]+nums[i] < 0 {
				lo++
			} else if nums[lo]+nums[hi]+nums[i] > 0 {
				hi--
			}
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i ++
		}
	}
	return res
}

