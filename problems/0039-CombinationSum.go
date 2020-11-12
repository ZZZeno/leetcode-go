package problems

import (
	"sort"
)

// 没有比较好的思路，参考了halforst的代码，自己手写了一遍

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	res := [][]int{}
	temp := []int{}
	findCombination(candidates, target, temp, &res, 0)
	return res
}

func findCombination(candidates []int, target int, tempRes []int, result *[][]int, index int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(tempRes))
		copy(temp, tempRes)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(candidates); i++ {
		if target < candidates[i] {
			break
		}

		tempRes = append(tempRes, candidates[i])
		findCombination(candidates, target-candidates[i], tempRes, result, i)
		tempRes = tempRes[:len(tempRes)-1]

	}
}


func CombinationSum(candidates []int, target int) [][]int {
	return combinationSum(candidates, target)

}
