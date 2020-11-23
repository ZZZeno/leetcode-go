package problems

import "fmt"

func climbStairs(n int) int {
	tmp1 := 1
	tmp2 := 2
	if n == 1 {
		return tmp1
	}
	if n == 2 {
		return tmp2
	}
	times := 2
	var res int
	for times < n {
		res = tmp1 + tmp2
		times ++
		tmp3 := tmp2
		tmp2 = tmp1 + tmp3
		tmp1 = tmp3
	}

	return res
}

func ClimbStairs(n int) {
	fmt.Println(climbStairs(n))
}
