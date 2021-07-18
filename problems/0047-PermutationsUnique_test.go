package problems

import (
	"fmt"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	testcase := []int{1, 1, 3}

	ret := permuteUnique(testcase)
	fmt.Println(ret)
}
