func kInversePairs(n int, k int) int {
    return countDp(n, k )
}

func count(n int, k int) int {
    if k == 0 {
        return 1
    }
    if k < 0 {
        return 0
    }
    res := 0
    for i:= 0;i<n;i++{
        res += count(n-1, k-i) // if we chose the first (i=0), then we don't have any inverse pair since the beginning. Hence we calculate it from n -1
    }

    return res
}

var mod = 1000000000+7

func countDp(n, k int) int {
    dp := make([][]int, n+1)
    for i:=0;i<n+1;i++{
        dp[i] = make([]int, k+1)
    }

    for i:=0;i<k+1;i++{
        dp[0][i] = 0
    }
    dp[0][0] = 1
    for i:=1;i<n+1;i++{
        slideWindow := sumSlidingWindow(dp[i-1], i)
        for ii:=0;ii<k+1;ii++{
            dp[i][ii] = slideWindow[ii] % mod
        }
    }

    return dp[n][k]

}


func sumSlidingWindow(data []int, windowSize int) []int {
    n := len(data)
    k := windowSize
    res := make([]int, n)
    for i:=0;i<n;i++{
        if i < k {
            res[i] = data[i]
            if i > 0 {
                res[i] += res[i-1]
            }
        }else{
            res[i] = res[i-1] + data[i] - data[i-k]
        }
    }
    return res
}