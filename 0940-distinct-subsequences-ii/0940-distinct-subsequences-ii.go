var MOD int = 1e9+7
func distinctSubseqII(s string) int {
    // number of non-unique subsequence, including empty string
    // generally it's 2^n, which can be compute iteratively
    //
    // dp[k] = number of subsequence ending with s[k-1]
    // dp[k] = 2 * dp[k-1], why?
    // because take it for example s = "ab"
    // dp[0] = {""}
    // dp[1] = {"", "a"}
    // dp[2] = dp[1] + for each e element at dp[1] put "b" at the end of it, hence it is twice the size of previous dp

    // now the problem is to count the distinc subsequence
    // let last[char] = i,where is is the last i value when computing the dp[i] where s[i-1] = char.
    // dp[k] = 2 * dp[k-1] - dp[last[s[k-1]]]
    // this is because when we want to add s[k-1] to dp[k-1], we need to eliminate the double counting happened of previous dp that has suffix letter char
    n := len(s)
    dp := make([]int, n + 1)
    for i:=0;i<=n;i++{
        dp[i] = -1
    }
    dp[0] = 1
    var lastDPIndex [26]int
    for i:=0;i<26;i++{
        lastDPIndex[i] = -1
    }

    for k:=0;k<n;k++{
        dp[k+1] = 2 * dp[k] % MOD
        char := int(s[k] - 'a')
        if lastDPIndex[char] >= 0 {
            dp[k+1] -= dp[lastDPIndex[char]]
            if dp[k+1] < 0 {
                dp[k+1] += MOD
            }
        }
        lastDPIndex[char] = k
    }
    dp[n] -= 1 // the empty string
    if dp[n] < 0 {
        dp[n] += MOD
    }

    return dp[n] % MOD

}