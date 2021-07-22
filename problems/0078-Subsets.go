package problems

func subsets(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return [][]int{}
	}
	if l == 1 {
		return [][]int{
			{},
			{nums[0]},
		}
	}
	tmp := subsets(nums[1:])
	res := subsets(nums[1:])
	//fmt.Println(tmp)
	for _, v := range tmp {
		v = append(v, nums[0])
		res = append(res, v)
	}
	return res
}

func subsetsVerBackTracking(nums []int) [][]int {
	var res [][]int
	var path []int
	subsetsBT(&res, nums, path, 0)
	return res
}

func subsetsBT(res *[][]int, arr []int, path []int, idx int) {
	var temp = make([]int, len(path))
	copy(temp, path)
	*res = append(*res, temp)

	for i := idx; i < len(arr); i++ {
		path = append(path, arr[i])
		subsetsBT(res, arr, path, i+1)
		path = path[0 : len(path)-1]
	}
}
