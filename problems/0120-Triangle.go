package problems

import "fmt"

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	lastRow := triangle[rows-1]
	steps := make([][]int, rows)
	for i := range steps {
		steps[i] = make([]int, len(lastRow))
	}
	for row := range triangle {
		for col := range triangle[row] {
			if col == 0 || col == row {
				if row == 0 {
					steps[row][col] = triangle[row][col]
				} else if col == 0 {
					steps[row][col] = steps[row-1][col] + triangle[row][col]
					//fmt.Println(steps[row-1][col], triangle[row][col])
				} else if col == row {
					steps[row][col] = steps[row-1][col-1] + triangle[row][col]
				}
			} else {
				steps[row][col] = min(steps[row-1][col-1]+triangle[row][col], steps[row-1][col]+triangle[row][col])
			}
		}
	}
	minStep := steps[rows-1][0]
	for i := range steps[rows-1] {
		if minStep > steps[rows-1][i] {
			minStep = steps[rows-1][i]
		}
	}
	return minStep
}

func MinimumTotal() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}
