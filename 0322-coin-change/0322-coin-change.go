func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    sort.Ints(coins)
    // let dp[i] = minimum coin that sums to amount=i
    // then dp[i] = min(dp[i-coins[a]]+1, dp[i-coins[b]]+1, ....)
    // base case dp[0] = 0
    
    dp := make([]int, amount + 1)
    for i:=0;i<=amount;i++{
        dp[i] = -1
    }
    dp[0] = 0 
    for i:=coins[0];i<=amount;i++{
        for _, coin := range coins {
            if i - coin < 0 {
                break
            }
            if dp[i-coin] != -1 {
                if dp[i] == -1{
                    dp[i] = dp[i-coin] + 1
                }else{
                    dp[i] = min(dp[i], dp[i-coin]+1)
                }
            }
        }
    }
    // fmt.Println(dp)
    return dp[amount]
}