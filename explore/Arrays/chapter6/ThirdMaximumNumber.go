package chapter6

import "sort"

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	var max int
	var secondMax int
	var thirdMax int
	mp := map[int]int{}
	for _, v := range nums {
		mp[v] = 1
		if len(mp) == 3 {
			break
		}
	}
	tempArr := make([]int, 0)
	for k, _ := range mp {
		tempArr = append(tempArr, k)
	}

	max, secondMax, thirdMax = firstThree(tempArr)
	if len(tempArr) < 3 {
		return thirdMax
	}
	for i := 3; i < len(nums); i ++ {
		if nums[i] == max || nums[i] == secondMax || nums[i] == thirdMax {
			continue
		}
		if nums[i] > max {
			temp := max
			max = nums[i]
			thirdMax = secondMax
			secondMax = temp
			continue
		}
		if nums[i] > secondMax {
			temp := secondMax
			secondMax = nums[i]
			thirdMax = temp
			continue
		}
		if nums[i] > thirdMax {
			thirdMax = nums[i]
		}
	}
	return thirdMax
}

func firstThree(nums []int) (int, int, int) {
	sort.Ints(nums)
	if len(nums) == 1 {
		return nums[0], nums[0], nums[0]
	}
	if len(nums) == 2 {
		return nums[1], nums[1], nums[1]
	}

	return nums[2], nums[1], nums[0]
}
