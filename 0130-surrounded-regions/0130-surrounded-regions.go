const (
    x = 'X'
    o = 'O'
)

func solve(board [][]byte)  {
    // this is basically to group the 0 as a union set
    // the 0 that is in the edge, basically the one that's not needed to be flipped
    
    m, n := len(board), len(board[0])
    set := NewDisjointSet(m * n)

    getIdx := func(row, col int) int {
        return row * n + col
    }
    
    notFlipIdx := -1
    addNotFlip := func(row, col int) {
        idx := getIdx(row, col)
        if notFlipIdx != -1 {
            set.Union(idx, notFlipIdx)
        }
        notFlipIdx = idx
    }

    var lastX = -1
    addX := func(row, col int) {
        idx := getIdx(row, col)
        if lastX != -1 {
            set.Union(idx, lastX)
        }
        lastX = idx
    }

    for r := 0;r < m;r++{
        if board[r][0] == o {
            addNotFlip(r, 0)
        }
        if board[r][n-1] == o {
            addNotFlip(r, n-1)
        }
    }
    for c:=0;c<n;c++{
        if board[0][c] == o {
            addNotFlip(0, c)
        }
        if board[m-1][c] == o {
            addNotFlip(m-1, c)
        }
    }
    var topIdx, leftIdx, idx int
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if board[i][j] == x {
                addX(i, j)
            }else{
                if i > 0 && board[i-1][j] == o {
                    idx = getIdx(i, j)
                    topIdx = getIdx(i-1, j)
                    set.Union(idx, topIdx)
                }
                if j > 0 && board[i][j-1] == o {
                    idx = getIdx(i, j)
                    leftIdx = getIdx(i, j-1)
                    set.Union(idx, leftIdx)
                }
                
            }
        }
    }

    // check the groupId of notFlip
    notFlipId := -1
    if notFlipIdx != -1 {
        notFlipId = set.Find(notFlipIdx)
    }

    for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if board[i][j] == o {
                if notFlipId == -1 || notFlipId != set.Find(getIdx(i, j)){
                    board[i][j] = x
                }
            }
        }
    }
}

type DisjointSet struct {
    parents []int
}

func NewDisjointSet(n int) DisjointSet {
    p := make([]int, n)
    for i:=0;i<n;i++{
        p[i] = -1
    }
    return DisjointSet{
        parents: p,
    }
}

func (d *DisjointSet) Find(x int) int {
    if d.parents[x] < 0 {
        return x
    }
    path := make([]int, 0)

    for d.parents[x] >= 0 {
        path = append(path, x)
        x = d.parents[x]
    }

    for _, p := range path {
        d.parents[p] = x
    }
    return x
}

func (d *DisjointSet) Union(a, b int) {
    if a == b {
        return
    }
    aId, bId := d.Find(a), d.Find(b)
    if aId == bId {
        return
    }
    aSize, bSize := -d.parents[aId], -d.parents[bId]
    if aSize > bSize {
        d.parents[bId] = aId
        d.parents[aId] = -aSize-bSize
    }else{
        d.parents[aId] = bId
        d.parents[bId] = -aSize-bSize

    }
}