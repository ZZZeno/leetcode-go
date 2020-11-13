package problems

func minPathSum(grid [][]int) int {
	sum := 0
	m := len(grid)
	n := len(grid[0])
	if m == 1 {
		for _, v := range grid[0] {
			sum += v
		}
		return sum
	}
	if n == 1 {
		for _, v := range grid {
			sum += v[0]
		}
		return sum
	}
	resArr := make([][]int, m)
	for i := 0; i < m; i++ {
		resArr[i] = make([]int, n)
	}
	resArr[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		resArr[i][0] += grid[i][0] + resArr[i-1][0]
	}
	for j := 1; j < n; j++ {
		resArr[0][j] += grid[0][j] + resArr[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			resArr[i][j] = min(grid[i][j]+resArr[i][j-1], grid[i][j]+resArr[i-1][j])
		}
	}
	return resArr[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
