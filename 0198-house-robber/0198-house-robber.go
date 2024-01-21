func rob(nums []int) int {
    n := len(nums)
    dp := make([]int, n)

    dp[0] = nums[0]
    res := dp[0]
    for i:=1;i<n;i++{
        if i-2 >= 0 {
            dp[i] = dp[i-2] + nums[i]
            if i-3 >= 0 {
                dp[i] = max(dp[i], dp[i-3]+ nums[i])
            }
        }else{
            dp[i] = nums[i]
        }
        res = max(res, dp[i])
    }
    return res
}