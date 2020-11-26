package chapter5

func moveZeroes(nums []int)  {
	noneZeroCount := 0
	for _, v := range nums {
		if v != 0 {
			noneZeroCount += 1
		}
	}
	if noneZeroCount == len(nums) {
		return
	}

	cur := 0
	for i := 0; i < len(nums); i ++ {
		if cur == noneZeroCount {
			break
		}
		if nums[i] == 0 {
			for nums[i] == 0 {
				i ++
			}
		}

		exch(nums, cur, i)
		cur ++
	}

}

func exch(arr []int, lo, hi int) {
	temp := arr[lo]
	arr[lo] = arr[hi]
	arr[hi] = temp
}
