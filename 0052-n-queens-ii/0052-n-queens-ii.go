var rowColDiff map[int]bool
var rowColSum map[int]bool
var colFlag map[int]bool

func putNQueens(n int, remain int, row int) int {
    
    if remain == 0 {
        return 1
    }
    total := 0
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
        total += putNQueens(n, remain - 1, row+1)
        
        colFlag[j] = false
        rowColDiff[diff] = false
        rowColSum[sum] = false
    }
    return total
}

func totalNQueens(n int) int {
    rowColDiff = make(map[int]bool)
    rowColSum = make(map[int]bool)
    colFlag = make(map[int]bool)
    return putNQueens(n, n, 0)
}