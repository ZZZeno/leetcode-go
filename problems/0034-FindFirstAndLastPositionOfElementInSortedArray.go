package problems

func searchRange(nums []int, target int) []int {
	startIndex := -1
	endIndex := -1

	lo := 0
	hi := len(nums) - 1

	startFlag := false
	for lo <= hi {
		if nums[lo] == target {
			startIndex = lo
			startFlag = true
		}

		if !startFlag {
			lo ++
		} else {
			if nums[hi] == target {
				endIndex = hi
				break
			}
			hi --
		}
	}

	return []int{startIndex, endIndex}
}
