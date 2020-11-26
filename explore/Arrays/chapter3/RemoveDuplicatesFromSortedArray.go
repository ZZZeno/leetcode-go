package chapter3

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			count++
		}
	}
	lo := 0
	for i := 1; i < len(nums); {

		if nums[i] != nums[i-1] && nums[i] != nums[lo] {
			lo += 1
			exch(nums, lo, i)
			i += 1
			continue
		}
		if nums[i] != nums[i-1] && nums[i-1] == nums[lo]{
			i ++
			continue
		}
		i ++
	}
	return count
}


func exch(nums []int, index1, index2 int) {
	temp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = temp
}

func RemoveDuplicates(nums []int) int {
	return removeDuplicates(nums)
}
