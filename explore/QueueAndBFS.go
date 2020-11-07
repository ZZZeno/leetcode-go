package explore

import "fmt"

// 按理来说leetcode的接口定义的grid是[][]byte类型的，传入的矩阵每个cell是'0'或者'1'，但是我实在是不知道如何把string类型转成byte类型，所以只能写一个[][]string的接口了

func NumIslands(grid [][]string) int {
	row := len(grid)
	col := len(grid[0])
	i := 0
	j := 0
	num := 0
	for i+j != row+col-2 {
		if grid[i][j] == "1" {
			num += 1
			Mark(grid, i, j, row, col)
			//Print2DimMatrix(grid)
		}
		if i+j != row+col-2 {
			i = (i + 1) % row
			if i == 0 && j < col {
				j += 1
			}
		}
	}


	return num
}

func Mark(grid [][]string, i, j, row, col int) {
	rowQueue := Constructor(row)
	colQueue := Constructor(col)
	rowQueue.EnQueue(i)
	colQueue.EnQueue(j)
	for !colQueue.IsEmpty() {
		r, c := getItem(&rowQueue, &colQueue)
		grid[r][c] = "0"
		if r+1 < row && grid[r+1][c] == "1" {
			addItem(r+1, c, &rowQueue, &colQueue)
		}
		if c+1 < col && grid[r][c+1] == "1" {
			addItem(r, c+1, &rowQueue, &colQueue)
		}
	}
}

func addItem(i, j int, rowQueue, colQueue *MyCircularQueue) {
	rowQueue.EnQueue(i)
	colQueue.EnQueue(j)
}

func getItem(rowQueue, colQueue *MyCircularQueue) (int, int) {
	i := rowQueue.Front()
	j := colQueue.Front()
	rowQueue.DeQueue()
	colQueue.DeQueue()
	return i, j
}

func Print2DimMatrix(data [][]string) {
	for _, item := range data {
		fmt.Printf("%v \n", item)
	}
	fmt.Println()
	fmt.Println()
}
