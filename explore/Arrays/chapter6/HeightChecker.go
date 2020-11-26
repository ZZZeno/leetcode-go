package chapter6

import "sort"

func heightChecker(heights []int) int {
	disorderCnt := 0
	temp := make([]int, len(heights))
	copy(temp, heights)
	sort.Ints(temp)
	for i := range heights {
		if temp[i] != heights[i] {
			disorderCnt += 1
		}
	}
	return disorderCnt
}
