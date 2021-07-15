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
	_permutate(&ret, nums, 0)

	return ret
}
