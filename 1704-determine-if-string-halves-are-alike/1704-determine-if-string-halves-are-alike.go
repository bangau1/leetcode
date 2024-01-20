var vowel = map[byte]bool{
    'a': true,
    'i': true,
    'u': true,
    'e': true,
    'o': true,
}
func halvesAreAlike(s string) bool {
    s = strings.ToLower(s)
    mid := len(s)/2
    count := 0
    for i:=0;i<mid;i++{
        if vowel[s[i]]{
            count++
        }
    }

    for i:=mid;i<len(s);i++{
        if vowel[s[i]]{
            count--
        }
        if count < 0 {
            return false
        }
    }

    return count == 0
}