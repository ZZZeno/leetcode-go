package problems

func Multiply(num1 string, num2 string) string {
	return multiply(num1, num2)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	temp := make([]string, len(num1))
	for i := range num1 {
		temp[i] = StringMultiply(num2, num1[len(num1)-1-i]) + pow(10, i)
	}
	return StringAdd(temp...)
}

func pow(x, n int) string {
	if n == 0 {
		return ""
	}
	return pow(x, n-1) + "0"
}

func StringAdd(num ...string) string {
	maxLen := 0
	for _, v := range num {
		if maxLen < len(v) {
			maxLen = len(v)
		}
	}
	carry := 0
	res := make([]byte, maxLen+1)
	for i := 0; i < maxLen; i++ {
		cur := 0
		for _, v := range num {
			if len(v)-1-i < 0 {
				continue
			}
			cur += int(v[len(v)-1-i] - 48)
		}
		left := (carry + cur) % 10
		carry = (carry + cur) / 10
		res[maxLen-i] = byte(left + 48)
	}
	if carry == 0 {
		return string(res[1:])
	}
	res[0] = byte(carry + 48)
	return string(res)
}

func StringMultiply(num string, multiply byte) string {
	l := len(num)
	res := make([]byte, len(num)+1)
	carry := 0
	for i := l - 1; i >= 0; i-- {
		cur := int(num[i]-48)*int(multiply-48) + carry
		carry = cur / 10
		left := cur % 10
		res[i+1] = byte(left + 48)
	}
	if carry == 0 {
		return string(res[1:])
	}
	res[0] = byte(carry + 48)
	return string(res)
}
