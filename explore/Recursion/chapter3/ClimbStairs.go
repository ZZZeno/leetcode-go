package chapter3

func climbStairs(n int) int {
	temp := make(map[int]int)
	temp[1] = 1
	temp[2] = 2
	return withCache(n, temp)
}

func withCache(n int, cache map[int]int) int {
	if cache[n] != 0 {
		return cache[n]
	}
	cacheN_1 := withCache(n-1, cache)
	cacheN_2 := withCache(n-2, cache)
	cache[n-1] = cacheN_1
	cache[n-2] = cacheN_2
	return cacheN_1 + cacheN_2
}