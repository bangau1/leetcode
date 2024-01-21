func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    
    if m == 0 {
        return n
    }else if n == 0 {
        return m
    }else if m == n && word1==word2{
        return 0
    }


    dp := make([][]int, m+1)
    for i:=0;i<=m;i++{
        dp[i] = make([]int, n+1)
    }
    dp[m][n] = 0 // base case
    for i:=0;i<n;i++{
        dp[m][i] = n-i
    }
    for i:=0;i<=m;i++{
        dp[i][n] = m-i
    }

    for i:=0;i<=m;i++{
        fmt.Println(dp[i])
    }

    for i:=m-1;i>=0;i--{
        for ii:=n-1;ii>=0;ii--{
            if word1[i] == word2[ii] {
                dp[i][ii] = dp[i+1][ii+1]
            }else{
                dp[i][ii] = min(
                    dp[i+1][ii],
                    dp[i][ii+1],
                    dp[i+1][ii+1],
                )+1
            }
        }
    }

    return dp[0][0]
}