func firstUniqChar(s string) int {
    counter := make([]int, 26)

    for i:=0;i<len(s);i++{
        counter[int(s[i]-'a')]++
    }

    for i:=0;i<len(s);i++{
        if counter[int(s[i]-'a')] == 1 {
            return i
        }
    }
    return -1
}