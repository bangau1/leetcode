func longestPalindrome(s string) string {
    n := len(s)
    res := string(s[0])

    for i:=1;i<n;i++{
        l, r := i-1, i+1
        found := false
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
            found = true
        }
        l = l+1
        r = r-1
        if found && len(res) < r-l+1{
            res = s[l:r+1]
        }


        l, r = i-1, i
        found = false
        for l >= 0 && r < n && s[l] == s[r] {
            l--
            r++
            found = true
        }
        l = l+1
        r = r-1
        if found && len(res) < r-l+1{
            res = s[l:r+1]
        }
    }

    return res
}