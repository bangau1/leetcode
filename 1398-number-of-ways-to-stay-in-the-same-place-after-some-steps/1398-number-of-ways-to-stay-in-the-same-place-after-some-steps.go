const mod = 1000000000+7
func numWays(steps int, arrLen int) int {
    
    // let dp[steps][i] is number of ways to reach pos 0 from pos i, with steps left.

    // dp[s][i] = dp[s-1][i-1]  + // left
    //            dp[s-1][i+1] + // right
    //            dp[s-1][i] // stay

    n := min(steps, arrLen) // we only care about max position at n = minimum(steps, arrLen)
    dp := make([]int, n)
    prev := make([]int, n)
    dp[0] = 1
    // fmt.Println(dp[0])
    for s:=1;s<=steps;s++{
        dp, prev = prev, dp
        for pos:=0;pos<n;pos++{
            dp[pos] = prev[pos] // stay
            if pos-1 >= 0{
                dp[pos] += prev[pos-1] % mod // left
            }
            if pos + 1 < n {
                dp[pos] += prev[pos+1] % mod // right
            }
        }
        // fmt.Println(dp[s])
    }
    return dp[0] % mod
}