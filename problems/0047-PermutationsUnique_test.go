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

func TestArrEQ(t *testing.T) {
	testcase := []int{1, 1, 3}
	testcase1 := []int{1, 1, 3}
	fmt.Println(arrEQ(testcase1, testcase))
}
