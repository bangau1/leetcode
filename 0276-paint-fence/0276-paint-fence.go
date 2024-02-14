func numWays(n int, k int) int {
    if n < 3 {
        res := 1
        for i:=0;i<n;i++{
            res = res * k
        }
        return res
    }
    // let dp[i] = number of ways to coloring i post with k
    // dp[1] = k
    // dp[2] = k * dp[1]
    // dp[3] = 
    //  (k-1)*dp[2] + (k-1)*dp[1]
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = k
    dp[2] = k*dp[1]
    for i:=3;i<=n;i++{
        dp[i] = dp[i-1]*(k-1) + dp[i-2]*(k-1)
    }
    return dp[n]
}
