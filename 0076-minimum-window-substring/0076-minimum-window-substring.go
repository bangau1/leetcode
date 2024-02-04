func minWindow(s string, t string) string {
    // can be solved by using hashmap + 2 pointers
    tMap := counter(t)
    sMap := counter(s)
    // initial check if s have all chars on t
    for char, count := range tMap {
        if sMap[char] < count {
            return ""
        }
    }

    isMatched := func(sMap map[byte]int) bool {
        for char, count := range tMap {
            if sMap[char] < count{
                return false
            }
        }
        return true
    }
    n := len(s)
    l, r := 0,0
    windowMap := make(map[byte]int)
    windowMap[s[r]]++
    res, minLen := "", n+1
    for r < n {
        ok := isMatched(windowMap)
        if ok {
            if minLen > r-l+1{
                minLen = r - l + 1
                res = s[l:r+1]
            }
            windowMap[s[l]]--
            l++
        }else{
            r++
            if r < n {
                windowMap[s[r]]++
            }
        }
    }
    return res
}

func counter(s string) map[byte]int {
    res := make(map[byte]int)
    for i:=0;i<len(s);i++{
        res[s[i]]++
    }

    return res
}