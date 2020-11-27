package problems

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	if n == 2 {
		return []string{"()()", "(())"}
	}
	ret := make([]string, 0)
	mp := make(map[string]bool, 0)
	for i := 1; i <= n/2; i++ {
		Order(i, n-i, mp)
		Wrapper(i, n-i, mp)
		//ret = append(ret, Order(i, n-i)...)

	}
	for k, _ := range mp {
		ret = append(ret, k)
	}
	return ret
}

func Order(front, behind int, mp map[string]bool) map[string]bool {
	frontParenthesis := generateParenthesis(front)
	backParenthesis := generateParenthesis(behind)
	lf := len(frontParenthesis)
	bf := len(backParenthesis)
	//res := make([]string, lf*bf)
	for i := 0; i < lf; i++ {
		for j := 0; j < bf; j++ {
			mp[frontParenthesis[i]+backParenthesis[j]] = true
			mp[backParenthesis[j]+frontParenthesis[i]] = true
		}
	}
	return mp
}

func Wrapper(outer, inner int, mp map[string]bool) {
	innerParenthesis := generateParenthesis(inner)
	//outerParenthesis := make([]string, len(innerParenthesis))
	for _, v := range innerParenthesis {
		mp[repeat('(', outer)+v+repeat(')', outer)] = true
	}
}

func repeat(cha byte, n int) string {
	res := ""
	for i := 0; i < n; i++ {
		res += string(cha)
	}
	return res
}

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}
