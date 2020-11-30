package chapter5

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K == 1 {
		return 0
	}
	if K == 2 {
		return 1
	}
	leavesCnt := myPow(2, N-2)
	if K > leavesCnt {
		//right tree
		leftRes := kthGrammar(N, K-leavesCnt)
		if leftRes == 0 {
			return 1
		}
		return 0

	} else {
		//left tree
		res := kthGrammar(N-1, (K+1)/2)
		return res
	}
}

func myPow(x int, n int) int {
	if n == 0 {
		return 1
	}
	temp := map[int]int{
		1:  x,
		-1: 1 / x,
	}

	return fromCache(x, n, temp)
}

func fromCache(x int, n int, cache map[int]int) int {
	if cache[n] != 0 {
		return cache[n]
	}
	a := fromCache(x, n/2, cache)
	if n%2 == 0 {
		return a * a
	} else {
		if n > 0 {
			return a * a * x
		}
		return a * a * 1/x
	}
}

func KthGrammar(N int, K int) int {
	return kthGrammar(N, K)
}
