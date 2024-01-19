func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    dp := make([][]int, n)
    for i:=0;i<n;i++ {
        dp[i] = make([]int, n)
        for ii:=0;ii<n;ii++{
            dp[i][ii] = math.MaxInt
        }
    }

    moves := [][]int {
        {-1, 0},
        {-1, -1},
        {-1, 1},
    }
    for i:=0;i<n;i++{
        for ii:=0;ii<n;ii++{
            if i == 0 {
                dp[i][ii] = matrix[i][ii]
            }else{
                for _, move := range moves {
                    r, c := i+move[0], ii+move[1]
                    if r >= 0 && r < n && c >= 0 && c < n {
                        dp[i][ii] = min(dp[i][ii], matrix[i][ii] + dp[r][c])
                    }
                }
            }
        }
    }
    res := dp[n-1][0]
    for i:=1;i<n;i++{
        res = min(res, dp[n-1][i])
    }
    return res
}