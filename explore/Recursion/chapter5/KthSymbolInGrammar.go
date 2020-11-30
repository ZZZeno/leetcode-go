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
		return leftRes ^ 1

	} else {
		//left tree
		return kthGrammar(N-1, K)
	}
}

func myPow(x int, n int) int {
	if n == 0 {
		return 1
	}

	return tailRecursionPow(x, n)
}

func tailRecursionPow(x int, n int) int {
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	a := tailRecursionPow(x, n/2)
	if n%2 == 0 {
		return a * a
	} else {
		if n > 0 {
			return a * a * x
		}
		return a * a * 1 / x
	}
}
