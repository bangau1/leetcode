func maxProfit(prices []int) int {
    n := len(prices)
    // another approach: increasing monotonic stack 
    stack := make([]int, 0)
    profit := 0
    for i:=0;i<n;i++{
        for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
            profit = max(profit, stack[len(stack)-1] - stack[0])
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, prices[i])
    }
    profit = max(profit, stack[len(stack)-1] - stack[0])
    return profit
}