var rowColDiff map[int]bool
var rowColSum map[int]bool
var colFlag map[int]bool

type result [][]string

func putNQueens(n int, remain int, row int, path [][]int, answer *result) {
    if remain == 0 {
        
        combi := make([]string, 0)
        for i:= 0;i<len(path);i++{
            y := path[i][1]
            queenStr := []byte(strings.Repeat(".", n))
            queenStr[y] = "Q"[0]
            combi = append(combi, string(queenStr))
        }
        *answer = append(*answer, combi)
        return
    }
    prevLen := len(path)
    i := row
    for j:=0;j<n;j++{
        if colFlag[j] {
            continue
        }
        sum := i+j
        diff := i-j

        if rowColDiff[diff]{
            continue
        }
        if rowColSum[sum] {
            continue
        }

        colFlag[j] = true
        rowColDiff[diff] = true
        rowColSum[sum] = true
        path = append(path, []int{i, j})
        putNQueens(n, remain - 1, row+1, path, answer)
        colFlag[j] = false
        rowColDiff[diff] = false
        rowColSum[sum] = false
        path = path[0:prevLen]

    }
}

func solveNQueens(n int) [][]string {
    // to check if we can place the queen withouut attack each other, we need 3 set of flags
    rowColDiff = make(map[int]bool)
    rowColSum = make(map[int]bool)
    colFlag = make(map[int]bool)

    var answer result
    putNQueens(n, n, 0, make([][]int, 0), &answer)
    return answer    
}