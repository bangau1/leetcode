var zero = "0"[0]


func numDecodings(s string) int {
    n := len(s)
    if n == 0 {
        return 1
    }

    if s[0] == zero {
        return 0
    }

    dp := make([]int, n+1)
    dp[0] = 1

    for i:=1;i<=n;i++{
        if s[i-1] != zero{
            dp[i] = dp[i-1]

            if i>=2 && s[i-2] != zero {
                num, _ := strconv.Atoi(s[i-2:i])
                if num <= 26 {
                    dp[i]+= dp[i-2]
                }
            }
        }else{
            if i>=2 && s[i-2] != zero {
                num, _ := strconv.Atoi(s[i-2:i])
                if num <= 26 {
                    dp[i]= dp[i-2]
                }
            }
        }
    }
    // fmt.Println(dp[1:])

    return dp[n]
}