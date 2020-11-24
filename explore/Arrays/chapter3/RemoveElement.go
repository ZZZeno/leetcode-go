package chapter3

func removeElement(nums []int, val int) int {
	count := 0
	for _, v := range nums {
		if v == val {
			count += 1
		}
	}
	if count == len(nums) {
		return 0
	}
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		if nums[lo] == val {
			for nums[hi] == val {
				hi --
			}
			tmp := nums[lo]
			nums[lo] = nums[hi]
			nums[hi] = tmp
		}
		lo ++
		if lo + count == len(nums) {
			break
		}
	}
	return len(nums) - count
}

func RemoveElements(nums []int, val int) int {
	return removeElement(nums, val)
}
