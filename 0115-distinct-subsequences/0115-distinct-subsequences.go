func numDistinct(s string, t string) int {
    ns := len(s)
    nt := len(t)

    if ns < nt {
        return 0
    }

    if ns == nt {
        if s == t {
            return 1
        }
        return 0
    }
    
    dp := make([][]int, ns+1)
    for i:=0;i<ns+1;i++{
        dp[i] = make([]int, nt+1)
        for ii:=0;ii<nt+1;ii++{
            dp[i][ii] = -1
        }
    }

    // dp[ns+1][nt+1]
    // dp[i][j] is number of subsquence of s[i...ns] to t[j....nt]
    // dp[i][j] = dp[i+1][j] + (if s[i] == t[j] --> add dp[i+1][j+1])
    return countDistinct(dp, s, 0, t, 0)
}

func countDistinct(dp [][]int, s string, si int, t string, ti int) int {
    total := 0
    defer func(){
        if dp[si][ti] == -1 {
            dp[si][ti] = total
        }
    }()

    if dp[si][ti] != -1 {
        return dp[si][ti]
    }

    if si >= len(s) && ti < len(t){
        return total
    }

    if ti >= len(t){
        total = 1
        return total
    }

    ns := len(s)-si
    nt := len(t)-ti

    if ns < nt {
        return total
    }

    if ns == nt {
        if s[si:] == t[ti:] {
            total = 1
            return total
        }
        return total
    }

    total = countDistinct(dp, s, si+1, t, ti) 
    if s[si] == t[ti]  {
        total += countDistinct(dp, s, si+1, t, ti+1)
    }
    return total
}