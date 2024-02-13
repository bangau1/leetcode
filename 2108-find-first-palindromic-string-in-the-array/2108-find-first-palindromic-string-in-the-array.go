func firstPalindrome(words []string) string {
    for _, word := range words {
        if isPalindrome(word) {
            return word
        }
    }
    return ""
}

func isPalindrome(s string) bool {
    n := len(s)
    if n == 1 {
        return true
    }
    if n <= 3 {
        return s[0] == s[n-1]
    }
    var l, r int
    if n % 2 == 0 {
        l, r = n/2 -1, n/2
    }else{
        l, r = n/2-1, n/2 + 1
    }

    for l >= 0 && r < n {
        if s[l] != s[r] {
            return false
        }
        l--
        r++
    }
    return true
}