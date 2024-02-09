func findTargetSumWays(nums []int, target int) int {
    // naive approach: dfs on all possible solution
    // var total int
    // dfs(0, 0, nums, target, &total)
    // return total

    // optimized version. some observations:
    // - at index i, we have 2 options, either to take +nums[i] or -nums[i]
    // - sums[i] at certain branch of dfs, refer to the sum of nums[0...i]
    // - we need to find whether the combination nums[i+1....n-1] == target - sums[i], if yes how many of them
    // - to help with that, we need to use dp[i]map[int]int
    // that we can build iteratively
    // dp[0][0] = 1
    // dp[1] = nums[i-1] + dp[0]... and -nums[i-1] + dp[0]

    n := len(nums)
    dp := make([]map[int]int, n + 1)
    for i:=0;i<=n;i++{
        dp[i] = make(map[int]int)
    }
    dp[0][0] = 1 
    for i:=1;i<=n;i++{
        for sum, count := range dp[i-1]{
            dp[i][nums[i-1]+sum]+=count
            dp[i][-nums[i-1]+sum]+=count
        }
    }
    return dp[n][target]
}

func dfs(index, sum int, nums []int, target int, total *int) {
    if index >= len(nums) {
        if sum == target {
            *total = *total + 1
        }
        return
    }

    dfs(index+1, sum + nums[index], nums, target, total)
    dfs(index+1, sum - nums[index], nums, target, total)
}

