package problems

func solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	for row := 0; row < len(board); row ++ {
		dfs_sr(board, row, 0)
		dfs_sr(board, row, len(board[0])-1)
	}
	for col := 1; col < len(board[0])-1; col ++ {
		dfs_sr(board, 0, col)
		dfs_sr(board, len(board)-1, col)
	}
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			switch board[i][j] {
			case 'X':
				continue
			case 'O':
				board[i][j] = 'X'
			case '@':
				board[i][j] = 'O'
			}
		}
	}
}

func dfs_sr(board [][]byte, row, col int) {
	switch board[row][col] {
	case 'X':
		return
	case 'O':
		board[row][col] = '@'
	}
	neighbors := findNeighbors(board, row, col)
	for _, neighbor := range neighbors {
		if board[neighbor[0]][neighbor[1]] == 'O' {
			dfs_sr(board, neighbor[0], neighbor[1])
		}
	}
}

func findNeighbors(board [][]byte, row, col int) [][2]int {
	var ret [][2]int
	var temp = [][2]int{
		{row-1, col},
		{row, col-1},
		{row+1, col},
		{row, col+1},
	}
	for _, item := range temp {
		if !filterd(board, item[0], item[1]) {
			ret = append(ret, item)
		}
	}
	return ret
}

func filterd(board [][]byte, row, col int) bool {
	maxRow := len(board)
	maxCol := len(board[0])
	return row < 0 || col < 0 || row >= maxRow || col >= maxCol
}
