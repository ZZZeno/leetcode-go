package pattern

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	nums := []int{1, 2}
	//fmt.Println(Echo(nums))
	ch := Sum(Square(Echo(nums)))
	for a := range ch {
		fmt.Println(a)
	}
}