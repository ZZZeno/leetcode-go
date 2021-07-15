package problems

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	testcase := []int{1, 2, 3}
	ret := permute(testcase)
	fmt.Println(ret)
}
