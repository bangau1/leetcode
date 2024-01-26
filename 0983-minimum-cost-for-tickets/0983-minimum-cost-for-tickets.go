func mincostTickets(days []int, costs []int) int {
    // dp[i] = min(
    //   dp[i-1] + cost(1d)
    //   dp[i-7] + cost(7d)
    //   dp[i-30] + cost(30d)   
    // ) when i is days

    // if i not in days, then dp[i] = dp[i-1]

    travel := make([]bool, 365+1)
    for i:=0;i<len(days);i++{
        travel[days[i]] = true
    }

    dp := make([]int, 365+1)
    f, l := days[0], days[len(days)-1]
    for i:=f;i<=l;i++{
        if travel[i] {
            dp[i] = dp[i-1] + costs[0]
            if i>=7 {
                dp[i] = min(dp[i], dp[i-7] + costs[1])
            }else{
                dp[i] = min(dp[i], costs[1])
            }
            if i>=30 {
                dp[i] = min(dp[i], dp[i-30] + costs[2])
            }else{
                dp[i] = min(dp[i], costs[2])
            }
        }else{
            dp[i] = dp[i-1]
        }
    }

    // fmt.Println(dp)

    return dp[l]
}