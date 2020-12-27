package pattern

func Sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func Calc(nums []int) int {
	in := make(chan int)
	for n := range nums {
		in <- n
	}
	ret := <-Sum(Square(in))
	return ret
}