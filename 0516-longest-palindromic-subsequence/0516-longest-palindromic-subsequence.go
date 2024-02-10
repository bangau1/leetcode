func longestPalindromeSubseq(s string) int {
    return longestCommonSubsequence(s, reverse(s))
}

func reverse(s string) string {
    n := len(s)
    res := make([]byte, n)
    
    for i:=0;i<n;i++{
        res[i] = s[n-1-i]
    }
    return string(res)
}

func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    curr := make([]int, n+1)
    prev := make([]int, n+1)

    for i:=1;i<=m;i++{
        curr, prev = prev, curr
        for ii:=1;ii<=n;ii++{
            if text1[i-1] == text2[ii-1] {
                curr[ii] = prev[ii-1] + 1
            }else{
                curr[ii] = max(prev[ii], curr[ii-1])
            }
        }
    }

    return curr[n]
}