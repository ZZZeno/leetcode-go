package problems

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var ret [][]int
	//_permutate(&ret, nums, 0)
	sort.Ints(nums)
	used := make([]bool, len(nums))
	var depth = 0
	var length = len(nums)
	_permuteWithTempArr(&ret, nums, []int{}, used, depth, length)
	return ret
}

func _permuteWithTempArr(ret *[][]int, ori, arr []int, used []bool, depth, length int) {
	if len(arr) == length {
		var temp = make([]int, length)
		copy(temp, arr)
		*ret = append(*ret, temp)
		return
	}
	for i := 0; i < length; i ++ {
		if used[i] || (i > 0 && ori[i] == ori[i-1] && !used[i-1]){
			continue
		}
		arr = append(arr, ori[i])
		used[i] = true
		_permuteWithTempArr(ret, ori, arr, used, depth+1, length)
		arr = arr[0:len(arr)-1]
		used[i] = false
	}
}
