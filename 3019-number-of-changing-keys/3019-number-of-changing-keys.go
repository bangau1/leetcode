func countKeyChanges(s string) int {
    count := 0
    s = strings.ToLower(s)
    prev := s[0]
    for i:=1;i<len(s);i++{
        if s[i] != prev{
            count++
        }
        prev = s[i]
    }
    return count
}