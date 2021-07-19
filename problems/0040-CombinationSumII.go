package problems

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	if candidates[0] > target {
		return nil
	}
	var res [][]int
	findCombination2(&res, candidates, []int{}, 0, target)
	return res
}

func findCombination2(res *[][]int, candidates, path []int, index, target int) {
	if target == 0 {
		var temp = make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	if index == len(candidates) {
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i ++ {
		if index != i && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		findCombination2(res, candidates, path, i+1, target-candidates[i])
		path = path[0: len(path)-1]
	}
}
//
//func arrEQ(arr, arr1 []int) bool {
//	if len(arr) != len(arr1) {
//		return false
//	}
//	for i := range arr {
//		if arr[i] != arr1[i] {
//			return false
//		}
//	}
//	return true
//}
//
//func arrIn(res *[][]int, arr []int) bool {
//	for _, item := range *res {
//		if arrEQ(arr, item) {
//			return true
//		}
//	}
//	return false
//}
