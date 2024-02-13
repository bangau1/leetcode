const (
    plus byte = '+'
    minus byte = '-'
)

func generatePossibleNextMoves(str string) []string {
    n := len(str)
    if n < 2 {
        return nil
    }

    res := make([]string, 0)
    for i:=0;i<n-1;i++{
        next := make([]byte, n)
        copy(next, str)
        if str[i] == plus && str[i] == str[i+1] {
            next[i], next[i+1] = minus, minus
            res = append(res, string(next))
        }
    }
    
    return res
}