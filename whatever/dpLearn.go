package whatever

// there are n steps. a frog can jump one step or two steps. how many different solutions here if the frog want
// to get to step n
func FrogJumpDP(n int) int {
	if n <= 2 {
		return n
	}
	solution := make([]int, n+1)
	solution[0] = 0
	solution[1] = 1
	solution[2] = 2
	for i := 3; i < n+1; i++ {
		solution[i] = solution[i-1] + solution[i-2]
	}
	return solution[n]
}



