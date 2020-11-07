package explore

import (
	"fmt"
	"testing"
)

func TestMyCircularQueue(t *testing.T) {

	obj := Constructor(6)
	fmt.Println(obj.EnQueue(6))
	fmt.Println(obj.Rear())
	fmt.Println(obj.Rear())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.EnQueue(5))
	fmt.Println(obj.Rear())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.Front())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj.DeQueue())
	fmt.Println(obj)
}

type row []int
type column []row

func TestNumsIsland(t *testing.T) {
	a := make([][]uint8, 10)
	for i := range a {
		a[i] = make([]uint8, 5)
	}
	fmt.Println(len(a))
	fmt.Println(len(a[0]))
}

func TestNums(t *testing.T) {
	//grid := [][]string{
	//	{"1","1","0","0","0"},
	//	{"1","1","0","0","0"},
	//	{"0","0","1","0","0"},
	//	//{"0","0","0","1","1"},
	//	{"1","1","0","0","0"},
	//}
	grid := [][]string{
		{"1","1","1","1","0"},
		{"1","1","0","1","0"},
		{"1","1","0","0","0"},
		{"0","0","0","0","0"},
	}
	fmt.Println(NumIslands(grid))
}

func TestMark(t *testing.T) {
	//grid := [][]string{
	//	{"1","1","0","0","0"},
	//	{"1","1","0","0","0"},
	//	{"0","0","1","0","0"},
	//	{"1","1","0","0","0"},
	//}

	grid := [][]string{
		{"1","1","1","1","0"},
		{"1","1","0","1","0"},
		{"1","1","0","0","0"},
		{"0","0","0","0","0"},
	}
	Print2DimMatrix(grid)
	Mark(grid, 0, 0, 4, 5)
	Print2DimMatrix(grid)
}
