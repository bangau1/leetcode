func canJump(nums []int) bool {
    // let dp[i] = maximum index it can reach from i
    // dp[i] = dp[i + nums[i]]
    n := len(nums)
    dp := make([]int, n)
    for i:=0;i<n;i++{
        dp[i] = -1
    }
    dp[n-1] = n-1
    dfs(0, dp, nums)
    return dp[0] == n -1
}

func dfs(idx int, dp []int, nums []int){
    n := len(nums)
    if idx == n-1 {
        dp[idx] = n-1
        return
    }
    if nums[idx] == 0{
        dp[idx] = idx
        return
    }
    if dp[idx] != -1 {
        return
    }

    for jump:=nums[idx];jump >=1;jump--{
        if idx + jump >= n-1{
            dp[idx] = n-1
            break
        }else{
            dfs(idx+jump, dp, nums)
            dp[idx] = max(dp[idx], dp[idx+jump])
        }
    }
}