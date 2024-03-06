func minCostConnectPoints(points [][]int) int {
    n := len(points)
    // since n max constraint is quite small, 1000, we can calculate all possible manhattan distance between each of them
    // then what we need to do is to create MST on all those points, ensuring all distance to connect all of them are minimum
    // we can use Kruskal, it's easier to implement
    // but we need to implement disjoint set, to ensure the edge that we keep adding doesn't cause a loop

    edges := make([][3]int, 0)
    var src, dst []int
    var dist int
    for i:=0;i<n;i++{
        src = points[i]
        for ii:=i+1;ii<n;ii++{
            dst = points[ii]
            dist = abs(src[0]-dst[0]) + abs(src[1]-dst[1])
            edges = append(edges, [3]int{i, ii, dist})
        }
    }
    sort.Slice(edges, func(a, b int) bool {
        return edges[a][2] < edges[b][2]
    })

    set := NewDisjointSet(n)
    var totalCost, totalEdge int
    for i:=0;i<len(edges);i++{
        a, b, cost := edges[i][0], edges[i][1], edges[i][2]
        if set.Union(a, b) {
            totalCost += cost
            totalEdge++

            if totalEdge == n-1{
                break
            }
        }
    }
    return totalCost
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

type DisjointSet struct {
    parents []int
}

func NewDisjointSet(n int) *DisjointSet{
    res := DisjointSet{
        parents: make([]int, n),
    }
    for i:=0;i<n;i++{
        res.parents[i] = -1 // -1 means it is its own set, with size -(-1) = 1
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
    aSize, bSize := -d.parents[aId], -d.parents[bId]
    if aSize > bSize {
        d.parents[bId] = aId
        d.parents[aId] = -aSize-bSize
    }else{
        d.parents[aId] = bId
        d.parents[bId] = -aSize-bSize

    }

    return true
}