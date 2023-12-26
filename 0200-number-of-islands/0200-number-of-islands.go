func find(parents []int, x int) int {
    if parents[x] < 0 {
        return x
    }

    path := make([]int, 0)
    
    for parents[x] >= 0 {
        path = append(path, x)
        x = parents[x]
    }
    for _, p := range path {
        parents[p] = x
    }
    return x
}

func union(parents []int, a, b int) {
    if a == b {
        return
    }
    aId := find(parents, a)
    bId := find(parents, b)
    
    if aId == bId {
        return
    }

    aSize := -parents[aId]
    bSize := -parents[bId]

    if aSize > bSize {
        parents[bId] = aId
        parents[aId] = -aSize -bSize
    }else{
        parents[aId] = bId
        parents[bId] = -aSize -bSize
    }
}

var land = "1"[0]
var toplefts = [][]int {{-1, 0}, {0, -1}}
func getTopLeftLands(grid [][]byte, row, col int) [][]int {
    
    res := make([][]int, 0)
    for _, tl := range toplefts{
        cellX, cellY := row + tl[0], col + tl[1]
        if cellX >= 0 && cellY >= 0 && grid[cellX][cellY] == land {
            res = append(res, []int{cellX, cellY})
        }
    }
    return res

}
func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    parents := make([]int, m*n)
    for i:=0;i<len(parents);i++{
        parents[i] = -1 // each is its own
    }

    for r:=0;r<m;r++{
        for c:=0;c<n;c++{
            curIdx := r * n + c 
            if grid[r][c] == land {
                // check the left and top only
                lands := getTopLeftLands(grid, r, c)
                for _, land := range lands {
                    landIdx := land[0] * n + land[1]
                    union(parents, curIdx, landIdx)
                }
            }else{
                parents[curIdx] = math.MaxInt
            }
        }
    }
    total := 0
    for _, x := range parents {
        if x < 0 {
            total++
        }
    }
    return total
}