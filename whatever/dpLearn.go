package whatever

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

func RobotMoveDP(m, n int) int {
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
