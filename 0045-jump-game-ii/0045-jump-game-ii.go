func jump(nums []int) int {
    // let dp[i] = minimum jump to reach target=n-1. If it can't then it's assign with math.MaxInt
    n := len(nums)
    dp := make([]int, n)
    for i:=0;i<n;i++{
        dp[i] = math.MaxInt
    }
    dp[n-1] = 0
    for i:=n-2;i>=0;i--{
        if nums[i] + i >= n-1{
            dp[i] = 1
        }else if nums[i] > 0{
            for jump:=nums[i];jump>=1;jump--{
                if dp[i+jump] != math.MaxInt {
                    dp[i] = min(dp[i], dp[i+jump]+1)
                }
            }
        }
    }
    return dp[0]
}