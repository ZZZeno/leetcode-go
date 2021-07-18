package problems

func _swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func _permutate(ret *[][]int, arr []int, start int) {
	if start == len(arr) {
		var temp = make([]int, len(arr))
		copy(temp, arr)
		*ret = append(*ret, temp)
		return
	}

	for i := start; i < len(arr); i ++ {
		_swap(arr, i, start)
		_permutate(ret, arr, start+1)
		_swap(arr, i, start)
	}

}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}

	var ret [][]int
	//_permutate(&ret, nums, 0)
	used := make([]bool, len(nums))
	var depth = 0
	var length = len(nums)
	permuteWithTempArr(&ret, nums, []int{}, used, depth, length)
	return ret
}

func permuteWithTempArr(ret *[][]int, ori, arr []int, used []bool, depth, length int) {
	if len(arr) == length {
		var temp = make([]int, length)
		copy(temp, arr)
		*ret = append(*ret, temp)
		return
	}
	for i := 0; i < length; i ++ {
		if used[i] {
			continue
		}
		arr = append(arr, ori[i])
		used[i] = true
		permuteWithTempArr(ret, ori, arr, used, depth+1, length)
		arr = arr[0:len(arr)-1]
		used[i] = false
	}
}