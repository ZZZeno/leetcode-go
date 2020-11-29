package chapter4

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	temp := map[int]float64{
		1:  x,
		-1: 1 / x,
	}

	return fromCache(x, n, temp)
}

func fromCache(x float64, n int, cache map[int]float64) float64 {
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

func GetPowXN(x float64, n int) float64 {
	return myPow(x, n)
}