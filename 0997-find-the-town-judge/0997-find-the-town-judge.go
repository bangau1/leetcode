func findJudge(n int, trust [][]int) int {
    // this is basically can be solved by counting the out-edge number
    out := make([]int, n + 1)
    in := make([]int, n + 1)
    for _, t := range trust {
        out[t[0]]++
        in[t[1]]++
    }

    for i:=1;i<=n;i++{
        if out[i] == 0 && in[i] == n-1 {
            return i
        }
    }
    return -1
}