func longestPalindrome(s string) string {
    n := len(s)
    res := string(s[0])

    for i:=1;i<n;i++{
        l, r := i-1, i+1

        for l >= 0 && r < n {
            if s[l] != s[r] {
                l = l+1
                r = r - 1
                if len(res) < (r-l+1){
                    res = s[l:r+1]
                }
                break
            }
            l--
            r++
        }

        l, r = i-1, i
        for l >= 0 && r < n {
            if s[l] != s[r] {
                l = l+1
                r = r - 1
                if len(res) < (r-l+1){
                    res = s[l:r+1]
                }
                break
            }
            l--
            r++
        }
    }

    return res
}