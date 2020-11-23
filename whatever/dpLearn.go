package whatever

import "fmt"

// there are n steps. a frog can jump one step or two steps. how many different solutions here if the frog want
// to get to step n
func FrogJumpDP(n int) int {
	if n <= 2 {
		return n
	}
	solution := make([]int, n+1)
	solution[0] = 0
	solution[1] = 1
	solution[2] = 2
	for i := 3; i < n+1; i++ {
		solution[i] = solution[i-1] + solution[i-2]
	}
	return solution[n]
}

// leetcode 62nd problem
func RobotMoveDP(m, n int) int {
	grid := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		grid[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		grid[i][1] = 1
	}
	for j := 0; j < n+1; j++ {
		grid[1][j] = 1
	}

	for i := 2; i < m+1; i++ {
		for j := 2; j < n+1; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	//for _, item := range grid {
	//	fmt.Printf("%v \n", item)
	//}
	return grid[m][n]
}

// leetcode 64th problem
func MinimumPathSum(grid [][]int) int {
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
	fmt.Printf("%v \n", resArr)
	return resArr[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// leetcode 72nd problem   Edit Distance
func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	grid := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		grid[i] = make([]int, l2+1)
	}
	grid[0][0] = 0
	for i := 1; i < l1+1; i++ {
		grid[i][0] = i
	}
	for j := 1; j < l2+1; j++ {
		grid[0][j] = j
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if word1[i-1] == word2[j-1] {
				grid[i][j] = grid[i-1][j-1]
			} else {
				grid[i][j] = min3(grid[i-1][j], grid[i][j-1], grid[i-1][j-1]) + 1
			}
		}
	}
	return grid[l1][l2]
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func MinDistance(word1, word2 string) int {
	return minDistance(word1, word2)
}
