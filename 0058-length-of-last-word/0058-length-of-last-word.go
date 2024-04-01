func lengthOfLastWord(s string) int {
    n := len(s)
    length := 0
    i := n-1
    for i >= 0 {
        for i >= 0 && s[i] == ' '{
            i--
        }
        for i >= 0 && s[i] != ' '{
            i--
            length++
        }
        return length
    }
    return length
}