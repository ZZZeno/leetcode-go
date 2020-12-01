package problems

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int {{1},}
	}
	res := [][]int{
		{1},
	}
	for i := 1; i < numRows; i ++ {
		res = append(res, generateNext(res[i-1]))
	}
	return res
}

func generateNext(base []int) []int {
	res := make([]int, len(base)+1)
	res[0] = 1
	for i := 1; i <= len(base)-1; i ++ {
		res[i] = base[i] + base[i-1]
	}
	res[len(base)] = 1
	return res
}
