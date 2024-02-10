// var charPos [26][]int
var dp [][]int
func longestPalindromeSubseq(s string) int {
    n := len(s)
    maxLen := 1
    // charPos = [26][]int{}
    dp = make([][]int,n)
    for i:=0;i<n;i++{
        dp[i] = make([]int, n)
        for ii:=0;ii<n;ii++{
            dp[i][ii] = -1
        }
        // dp[i][i] = 1
    }

    // for i:=0;i<n;i++{
    //     letter := int(s[i]-'a')
    //     charPos[letter] = append(charPos[letter], i)
    // }
    for i:=0;i<n;i++{
        l, r := i-1, i + 1
        maxLen = max(maxLen, dfs(l, r, &s) + 1)
        
        l, r = i-1, i
        maxLen = max(maxLen, dfs(l, r, &s))
    }

    return maxLen
}

func dfs(l,r int, sPtr *string) int {
    if l == r {
        panic("unexpected")
    }
    s := *sPtr
    n := len(s)
    if l < 0 || r >= n {
        return 0
    }
    if dp[l][r] != -1 {
        return dp[l][r]
    }

    result := 0
    if s[l] == s[r] {
        result = 2 + dfs(l-1, r+1, sPtr)
    }else{
        result = max(dfs(l-1, r, sPtr), dfs(l, r+1, sPtr))
        // choose nothing
        // maxi = dfs(l-1, r+1, sPtr)
        // maxi = max(maxi, dfs(l-1))
        
        // choose left
        // letter := s[l]-'a'
        // leftPos := charPos[letter]
        // idx := sort.SearchInts(leftPos, r+1)
        // if idx < len(leftPos) && leftPos[idx] > r {
        //     if leftPos[idx-1] >= r {
        //         panic("unexpected error")
        //     }
        //     maxi = max(maxi, dfs(l-1, leftPos[idx]+1, sPtr) + 2 )
        // }

        // chose right
        // letter = s[r]-'a'
        // rightPos := charPos[letter]
        // idx = sort.SearchInts(rightPos, l-1)
        // if rightPos[idx] >= 0 && rightPos[idx] < l {
        //     if rightPos[idx+1] <= l {
        //         panic("unexpected error")
        //     }
        //     maxi = max(maxi, dfs(rightPos[idx]-1, r+1, sPtr) + 2 )
        // }

        // result = maxi
    }

    dp[l][r] = max(dp[l][r], result)
    return dp[l][r]
}