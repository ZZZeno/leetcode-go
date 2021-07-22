package problems

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return [][]int{nil, {nums[0]}}
	}
	sort.Ints(nums)
	var path []int
	var res [][]int
	var used = make([]bool, len(nums))
	subsets2BT(&res, path, 0, nums, used)
	return res
}

func subsets2BT(res *[][]int, path []int, idx int, arr []int, used []bool) {
	//if idx == len(arr) {
	//	var temp = make([]int, len(path))
	//	copy(temp, path)
	//	(*res) = append(*res, temp)
	//}
	var temp = make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := idx; i < len(arr); i ++ {
    // if a duplicated elem is selected, the first one of duplicated elements must be selected, or else skip this condition
		if i-1 >= 0 && arr[i] == arr[i-1] && !used[i-1]{
			continue
		}
		path = append(path, arr[i])
		used[i] = true
		subsets2BT(res, path, i+1, arr, used)
		path = path[0: len(path)-1]
		used[i] = false
	}
}
