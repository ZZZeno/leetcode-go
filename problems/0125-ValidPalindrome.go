package problems

func isPalindrome(s string) bool {
    if len(s) == 0 || len(s) == 1{
        return true
    }
    lo := 0
    hi := len(s)-1
    
    for lo < hi {
        for lo < hi && !isAlphaNum(byte(s[lo])) {
            lo ++
        }
        for lo < hi && !isAlphaNum(byte(s[hi])) {
            hi --
        }
        if !isEqual(byte(s[lo]), byte(s[hi])) {
            return false
        }
        lo ++
        hi --
    }
    return true
}

func isEqual(a, b byte) bool {
    if isAlpha(a) && isAlpha(b) {
        if (isCaps(a) && isCaps(b)) || (!isCaps(a) && !isCaps(b)) {
            return a == b
        }
        reverseA := reverse(a)
        return reverseA == b    
    }
    return a == b
}

func isCaps(a byte) bool {
    return a >= 97 && a <= 122
}

func isAlphaNum(a byte) bool {
    return (a >= 97 && a <= 122) || (a >= 65 && a <= 90) || (a >= 48 && a <= 57)
}

func isAlpha(a byte) bool {
    return (a >= 97 && a <= 122) || (a >= 65 && a <= 90)
}

func reverse(a byte) byte {
    if a >= 97 && a <= 122 {
        return a - 32
    }
    return a + 32
}
