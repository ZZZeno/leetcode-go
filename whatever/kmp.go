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

func pairing(target string, pattern string) int {
	var index = -1
	next := kmpNext(pattern)
	now := 0
	for i := 0; i < len(target);  {
		if target[i] == pattern[now] {
			now ++
			i ++
		} else if now == 0 {
			i ++
		} else {
			if next[now-1] == 0 {
				now = 0
			} else {
				now = next[now-1] + 1
			}
		}
		if now == len(pattern) {
			return i - now
		}
	}
	return index
}
