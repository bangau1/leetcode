func canJump(nums []int) bool {
    // let dp[i] = maximum index it can reach from i
    // dp[i] = dp[i + nums[i]]
    n := len(nums)
    dp := make([]int, n)
    for i:=0;i<n;i++{
        dp[i] = -1
    }
    dp[n-1] = n-1
    for i:=n-2;i>=0;i--{
        if i + nums[i] >= n-1{
            dp[i] = n-1
        }else{
            for jump:=1;jump<=nums[i];jump++{
                if dp[i+jump] == n -1 {
                    dp[i] = n-1
                    break
                }
            }
        }
    }
    return dp[0] == n-1
}