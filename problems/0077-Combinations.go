package problems

func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	combineVer1(&res, n, k, path, 1)
	return res
}

func combineVer1(res *[][]int, n, k int, path []int, idx int) {
	if len(path) == k {
		var temp = make([]int, k)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}
	for i := idx; i <= n; i ++ {
		path = append(path, i)
		combineVer1(res, n, k, path, i+1)
		path = path[0:len(path)-1]
	}
}
