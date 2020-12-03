func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	num9 := 0
	for _, c := range digits {
		if c == '9' || c == '7' {
			num9 += 1
		}
	}
	resCnt := intPow(4, num9) * intPow(3, len(digits)-num9)
	res := make([]string, resCnt)
	mp := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	lastLen := 1
	for i := len(digits) - 1; i >= 0; i-- {
		curDigit := digits[i]
		curLi := mp[curDigit]
		curLen := len(curLi)
		for j := 0; j < resCnt; j++ {
			charIndex := 0
			if i == 0 {
				charIndex = j / lastLen
			} else if i == len(digits)-1 {
				charIndex = j % curLen
			} else {
				//fmt.Println("sadfagasdg")
				mod := j / lastLen
				charIndex = mod % curLen
			}
			res[j] = curLi[charIndex] + res[j]
		}
		lastLen *= curLen
	}
	return res
}

func intPow(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	return intPow(x, n-1) * x
}
