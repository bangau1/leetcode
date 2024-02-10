func countSubstrings(s string) int {
    total := 0
    n := len(s)
    for i:=0;i<n;i++{
        total += 1 // a single char is palindrom by itself
        // fmt.Println(i)
        l, r := i-1, i+1
        for l >= 0 && r < n {
            if s[l] == s[r] {
                total++
                // fmt.Println(l, r)
            }else{
                break
            }
            l--
            r++
        }
    }
    for i:=0;i<n-1;i++{
        l, r := i, i+1
        for l >= 0 && r < n {
            if s[l] == s[r] {
                total++
                // fmt.Println(l, r)
            }else{
                break
            }
            l--
            r++
        }
    }

    return total
}