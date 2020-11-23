package problems


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
