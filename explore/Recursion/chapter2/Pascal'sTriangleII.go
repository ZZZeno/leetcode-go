package chapter2

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	return moveOneIndexAdd(getRow(rowIndex-1))
}

func moveOneIndexAdd(arr []int) []int {
	res := make([]int, len(arr)+1)
	res[0] = 1
	res[len(arr)] = 1
	for i := 1; i < len(arr); i ++ {
		res[i] = arr[i] + arr[i-1]
	}
	return res
}