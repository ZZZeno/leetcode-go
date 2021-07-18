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

