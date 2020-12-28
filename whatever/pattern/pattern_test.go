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

func TestVisitor(t *testing.T) {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}
	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)

}
