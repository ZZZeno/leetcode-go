package problems

func maxCount(m int, n int, ops [][]int) int {
    if len(ops) == 0 {
        return m * n
    }
    minM := m
    minN := n
    
    for _, v := range ops {
        if v[0] < minM {
            minM = v[0]
        }
        if v[1] < minN {
            minN = v[1]
        }
    }
    return minM * minN
}
