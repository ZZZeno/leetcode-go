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

func ThreeSum(nums []int) {
	fmt.Printf("%v\n", threeSum(nums))
}
