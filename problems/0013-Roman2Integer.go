package problems

var romaIntMap = map[string]int {
	"I": 1,
	"IV": 4,
	"V": 5,
	"IX": 9,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	ret := 0
	index := 0
	for {
		next, step := findNextSymbol(s, index)
		if next < 0 {
			break
		}
		ret += next
		index += step
	}
	return ret
}

func findNextSymbol(s string, startIndex int) (int, int) {
	if startIndex == len(s) {
		return -1, -1
	}
	if startIndex == len(s)-1 {
		return romaIntMap[s[startIndex:]], 1
	}
	if romaIntMap[s[startIndex: startIndex+1]] < romaIntMap[s[startIndex+1: startIndex+2]] {
		return romaIntMap[s[startIndex+1: startIndex+2]]-romaIntMap[s[startIndex: startIndex+1]], 2
	}
	return romaIntMap[s[startIndex:startIndex+1]], 1
}
