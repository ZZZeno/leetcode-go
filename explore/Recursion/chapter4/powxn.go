package chapter4

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	return tailRecursionPow(x, n)
}

func tailRecursionPow(x float64, n int) float64 {
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

func GetPowXN(x float64, n int) float64 {
	return myPow(x, n)
}
