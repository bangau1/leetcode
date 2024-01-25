func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m + 1)
    for i:=0;i<=m;i++{
        dp[i] = make([]int, n + 1)
    }

    for i:=1;i<=m;i++{
        for ii:=1;ii<=n;ii++{
            if text1[i-1] == text2[ii-1] {
                dp[i][ii] = dp[i-1][ii-1] + 1
            }else{
                dp[i][ii] = max(dp[i-1][ii], dp[i][ii-1])
            }
        }
    }

    return dp[m][n]
}