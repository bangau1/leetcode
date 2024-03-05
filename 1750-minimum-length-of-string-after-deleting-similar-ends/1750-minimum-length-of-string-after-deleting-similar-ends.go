func minimumLength(s string) int {
    n := len(s)
    l, r := 0, n-1
    var char byte
    // use two pointer
    // check if l, r same character, if yes, proceed until it breaks
    // if not same, break
    for l < r {
        char = s[l]
        // fmt.Println(l, r)
        if char != s[r] {
            break
        }else{
            for l < r && s[l] == char{
                l++
            }
            for r >= l && s[r] == char {
                r--
            }
        }

    }

    return r-l+1
}