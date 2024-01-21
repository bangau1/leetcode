func isSubsequence(s string, t string) bool {
    if len(s) >= len(t) && s != t {
        return false
    }

    c := 0
    count := 0
    for i:=0;i<len(s);i++{
        for c < len(t) && s[i]!=t[c]{
            c++
        }
        if c < len(t) && s[i] == t[c]{
            count++
            c++
        }else{
            return false
        }
    }
    return true
}