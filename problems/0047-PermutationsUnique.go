package problems

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var ret [][]int
	_permutate_1(&ret, nums, 0)
	return ret
}

func _swap_1(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func arrEQ(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func arrContains(ret *[][]int, arr []int) bool {
	for _, item := range *ret {
		if arrEQ(item, arr) {
			return true
		}
	}
	return false
}

func _permutate_1(ret *[][]int, arr []int, start int) {
	if start == len(arr) {
		if arrContains(ret, arr) {
			return
		}
		var temp = make([]int, len(arr))
		copy(temp, arr)
		*ret = append(*ret, temp)
		return
	}

	for i := start; i < len(arr); i ++ {
		_swap_1(arr, i, start)
		_permutate_1(ret, arr, start+1)
		_swap_1(arr, i, start)
	}
}
