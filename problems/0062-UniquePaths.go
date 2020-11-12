package problems

// original file in whatever/dpLearn.go, function RobotMoveDP(m, n int)

func uniquePaths(m int, n int) int {
	grid := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		grid[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i ++ {
		grid[i][1] = 1;
	}
	for j := 0; j < n+1; j ++ {
		grid[1][j] = 1;
	}

	for i := 2; i < m+1; i ++ {
		for j := 2; j < n+1; j ++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	//for _, item := range grid {
	//	fmt.Printf("%v \n", item)
	//}
	return grid[m][n]

}
