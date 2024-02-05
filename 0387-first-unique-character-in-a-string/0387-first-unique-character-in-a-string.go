func firstUniqChar(s string) int {
    counter := make(map[byte]int)

    for i:=0;i<len(s);i++{
        counter[s[i]]++
    }

    for i:=0;i<len(s);i++{
        if counter[s[i]] == 1 {
            return i
        }
    }
    return -1
}