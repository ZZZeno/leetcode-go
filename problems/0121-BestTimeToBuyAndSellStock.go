package problems

func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1 {
        return 0
    }
    buy := 0
    sell := 1
    maxProfit := prices[sell] - prices[buy]
    min := prices[0]
    for i := 0; i < len(prices); i ++ {
        if prices[i] < min {
            min = prices[i]
        }
        if prices[i] - min > maxProfit{
            maxProfit = prices[i] - min
            sell = i
        }
    }
    return maxProfit
}
