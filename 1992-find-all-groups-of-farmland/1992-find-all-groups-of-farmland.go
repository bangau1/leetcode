func findFarmland(land [][]int) [][]int {
    m, n := len(land), len(land[0])
    ds := NewDisjointSet(m * n)

    posMap := make(map[int][4]int)
    cellId := 0
    for r:=0;r<m;r++{
        for c:=0;c<n;c++{
            if land[r][c] == 1 {
                posMap[cellId] = [4]int{r,c,r,c}
            }
            cellId++
        }
    }
    cellId = 0
    prevCellId := -1
    for r:=0;r<m;r++{
        for c:=0;c<n;c++{
            
            if land[r][c] != 1 {
                continue
            }
            cellId = getCellId(r, c, n)
            if r > 0 && land[r-1][c] == 1 {
                prevCellId = getCellId(r-1, c, n)
                uid, puid := ds.Find(cellId), ds.Find(prevCellId)
                newUnion := union(posMap[uid], posMap[puid])

                if ds.Union(cellId, prevCellId) {
                   if ds.Find(uid) == uid {
                        delete(posMap, puid)
                        posMap[uid] = newUnion
                   }else{
                        delete(posMap, uid)
                        posMap[puid] = newUnion
                   } 
                }
            }

            if c > 0 && land[r][c-1] == 1 {
                prevCellId = getCellId(r, c-1, n)
                uid, puid := ds.Find(cellId), ds.Find(prevCellId)
                newUnion := union(posMap[uid], posMap[puid])

                if ds.Union(cellId, prevCellId) {
                   if ds.Find(uid) == uid {
                        delete(posMap, puid)
                        posMap[uid] = newUnion
                   }else{
                        delete(posMap, uid)
                        posMap[puid] = newUnion
                   } 
                }
            }
        }
    }

    res := make([][]int, 0)
    for _, pos := range posMap {
        res = append(res, []int{pos[0], pos[1], pos[2], pos[3]})
    }
    return res
}

func union(a, b [4]int) [4]int{
    return [4]int{ 
        min(a[0], b[0]),
        min(a[1], b[1]),
        max(a[2], b[2]),
        max(a[3], b[3]),
    }
}

func getCellId(r, c int, colSize int) int {
    return r * colSize + c
}

func getPos(cellId, colSize int) (int, int) {
    return cellId/colSize, cellId % colSize
}


type DisjointSet struct {
	parents []int
}

func NewDisjointSet(n int) *DisjointSet {
	res := DisjointSet{
		parents: make([]int, n),
	}

	for i := 0; i < n; i++ {
		res.parents[i] = -1 // initialize the union size to 1 (negative value to store the size of the union, while also saying that it's a union leaderid)
	}

	return &res
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

	aSize, bSize := d.GetUnionSize(aId), d.GetUnionSize(bId)

	if bSize > aSize {
		d.parents[aId] = bId
		d.parents[bId] = -(aSize + bSize)
	} else {
		d.parents[bId] = aId
		d.parents[aId] = -(aSize + bSize)
	}
	return true
}

func (d *DisjointSet) GetUnionSize(a int) int {
	return -d.parents[d.Find(a)]
}

func (d *DisjointSet) GetTotalUnions() int {
	var count int
	for i := 0; i < len(d.parents); i++ {
		if d.parents[i] < 0 {
			count += 1
		}
	}
	return count
}
