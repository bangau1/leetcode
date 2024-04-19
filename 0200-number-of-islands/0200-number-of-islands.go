type DisjointSet struct{
    parents []int
}

func NewDisjointSet(n int) DisjointSet{
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

func (d *DisjointSet) Union(a, b int) bool {
    if a == b {
        return false
    }

    aId, bId := d.Find(a), d.Find(b)
    if aId == bId {
        return false
    }

    aSize, bSize := d.Size(aId), d.Size(bId)
    if aSize > bSize {
        d.parents[bId] = aId
        d.parents[aId] = -aSize -bSize
    }else{
        d.parents[aId] = bId
        d.parents[bId] = -aSize -bSize
    }
    return true
}

func (d *DisjointSet) Size(a int) int {
    aId := d.Find(a)
    return -d.parents[aId]
}

func (d *DisjointSet) SetSize() int {
    c := 0
    for _, p := range d.parents {
        if p < 0 {
            c++
        }
    }
    return c
}

const (
    one = byte('1')
    zero = byte('0')
) 


func numIslands(grid [][]byte) int {
    m, n := len(grid), len(grid[0])
    set := NewDisjointSet(m*n)
    getIdx := func(row, col int) int {
        return row * n + col
    } 
    // getRC := func(idx int) (int, int) {
    //     return idx / n, idx % n
    // }
    var topIdx, leftIdx int
    var zeroIdx = -1
    var tempIdx int
    for i:=0;i<m;i++{
        for j:=0;j<n;j++{

            if grid[i][j] == zero {
                tempIdx = getIdx(i, j)

                if zeroIdx != -1 {
                    set.Union(zeroIdx, tempIdx)  
                }   
                zeroIdx = tempIdx
                // r, c := getRC(set.Find(zeroIdx))
                // assert(grid[r][c] == zero)
                continue
            }

            idx := getIdx(i, j)
            if i > 0 && grid[i-1][j] == one {
                topIdx = getIdx(i-1, j)
                set.Union(idx, topIdx)
            }

            if j > 0 && grid[i][j-1] == one {
                leftIdx = getIdx(i, j-1)
                set.Union(idx, leftIdx)
            }
            // r, c := getRC(set.Find(idx))
            // assert(grid[r][c] == one)
        }
    }

    if zeroIdx != -1 {
        return set.SetSize() - 1
    }
    return set.SetSize()
}


func assert(cond bool) {
    if !cond {
        panic("assert failed")
    }
}