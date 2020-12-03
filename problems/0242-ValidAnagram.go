package problems

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := map[byte]int{}

	for i := 0; i < len(s); i ++ {
		mp[s[i]] += 1
		mp[t[i]] -= 1
	}
	for _, v := range mp {
		if v != 0 {
			return false
		}
	}
	return true
}
