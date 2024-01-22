func isOneEditDistance(s string, t string) bool {
    m, n := len(s), len(t)

    if m == 0 {
        return n == 1
    }else if n == 0 {
        return m == 1
    }else if m == n && m == 0 {
        return false
    }

    if abs(m-n) >= 2 {
        return false
    }

    s1, s2 := s, t
    if len(s1) < len(s2) {
        s1, s2 = s2, s1
    }

    // same length case
    if len(s1) == len(s2) {
        diffCount := 0
        for i:=0;i<len(s1);i++{
            if s1[i] != s2[i]{
                diffCount++
            }

            if diffCount > 1 {
                return false
            }
        }
        return diffCount == 1
    }


    // differ by 1 length case
    l, r := 0, 0
    for r < len(s2) {
        if s1[l] != s2[r] {
           return s1[l+1:] == s2[r:] 
        }
        l++
        r++
    }

    
    return true
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func editDistance(s, t string) int {
    m, n := len(s), len(t)
    if m == 0 {
        return n
    }else if n == 0 {
        return m
    }

    if m == n && s == t {
        return 0
    }

    dp := make([][]int, m + 1)
    for i:=0;i<m+1;i++{
        dp[i] = make([]int, n + 1)
    }

    for i:=0;i<=m;i++{
        dp[i][n] = m-i 
    }

    for i:=0;i<=n;i++{
        dp[m][i] = n-i
    }

    for i:=m-1;i>=0;i--{
        for ii:=n-1;ii>=0;ii--{
            if s[i] == t[ii] {
                dp[i][ii] = dp[i+1][ii+1]
            }else{
                dp[i][ii] = min(
                    dp[i+1][ii],
                    dp[i][ii+1],
                    dp[i+1][ii+1],
                ) + 1
            }
        }
    }

    return dp[0][0]
}