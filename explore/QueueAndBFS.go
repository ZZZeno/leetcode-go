package explore

import "fmt"

func NumIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	i := 0
	j := 0
	num := 0
	count1 := 0
	marked := 0
	for _, cols := range grid {
		for _, cell := range cols {
			count1 += int(cell-'0')
		}
	}
	if row == col && row == 1 {
		if grid[0][0] == '1' {
			return 1
		}
		return 0
	}
	
	for marked < count1 {
		if grid[i][j] == '1' {
			num += 1
			marked += Mark(grid, i, j, row, col)
			Print2DimMatrix(grid)
		}

		i = (i + 1) % row
		if i == 0 && j < col {
			j += 1
		}
		if grid[i][j] == '1' {
			num += 1
			marked += Mark(grid, i, j, row, col)
		}
	}

	return num
}

func Mark(grid [][]byte, i, j, row, col int) int {
	marked := 0
	rowQueue := Constructor(row)
	colQueue := Constructor(col)
	rowQueue.EnQueue(i)
	colQueue.EnQueue(j)
	if i < 0 || j < 0 || i >= row || j >= col {
		return 0
	}
	if grid[i][j] == '0' {
		return 0
	}

	grid[i][j] = '0'
	marked += 1

	marked += Mark(grid, i+1, j, row, col)
	marked += Mark(grid, i, j+1, row, col)
	marked += Mark(grid, i-1, j, row, col)
	marked += Mark(grid, i, j-1, row, col)
	return marked
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

func Print2DimMatrix(data [][]byte) {
	for _, item := range data {
		fmt.Printf("%v \n", string(item))
	}
	fmt.Println()
	fmt.Println()
}
