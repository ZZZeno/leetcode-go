package whatever

func kmpNext(pattern string) []int {
	length := len(pattern)
	next := make([]int, length)
	now := 0

	for i := 1; i < length;  {
		if pattern[now] == pattern[i] {
			now += 1
			next[i] = now
			i++
		} else if now > 0 {
			now = next[now-1]
		} else {
			next[i] = 0
			i ++
		}
	}
	return next
}
