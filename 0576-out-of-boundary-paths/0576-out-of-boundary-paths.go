func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    var total int
   
    dp := make([][]map[int]int, m)
    for i:=0;i<m;i++{
        dp[i] = make([]map[int]int, n)
        for ii:=0;ii<n;ii++{
            dp[i][ii] = make(map[int]int)
        }
    }

    visitAll(dp, startRow, startColumn, maxMove, &total)
    return total
}
var moves = [][]int{
    {0, -1},
    {0, 1},
    {-1, 0},
    {1, 0},
}
var mod = 1000000000+7
func visitAll(dp [][]map[int]int, row, col int, moveLeft int, total *int) int{
    r, c := len(dp), len(dp[0])
    res := 0
    if row < 0 || row >= r || col < 0 || col >= c {
        *total = (*total + 1) % mod
        res = 1
        return res
    }
  
    if moveLeft > 0 {
        memory := dp[row][col]
        if tot, ok:=memory[moveLeft];ok{
            *total += tot
            return tot
        }else{
            subTot := 0
            for _, move := range moves {
                ri, ci := row + move[0], col + move[1]
                subTot = (subTot + visitAll(dp, ri, ci, moveLeft-1, total)) % mod
            }
            dp[row][col][moveLeft] = subTot
            return subTot
        }
        
    }
    return 0

}