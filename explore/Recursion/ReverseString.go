package Recursion

func reverseString(s []byte)  {
	lo := 0
	hi := len(s) - 1
	for lo < hi {
		exch(s, lo, hi)
		lo ++
		hi --
	}
}

func exch(s []byte, lo, hi int) {
	temp := s[lo]
	s[lo] = s[hi]
	s[hi] = temp
}

func ReverseString(s []byte) {
	reverseString(s)
}