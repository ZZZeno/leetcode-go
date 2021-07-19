package problems

import (
	"fmt"
	"testing"
)

func Test_findCombination2(t *testing.T) {
	nums := []int{10,1,2,7,6,1,5}
	//nums = []int{1, 1, 7}
	ret := combinationSum2(nums, 8)
	fmt.Println(ret)
}
