func findRedundantConnection(edges [][]int) []int {
    // this problem can be solved by using Disjoint Set
    // intuition:
    // - when try adding the edge, check whether they are already in the same union or not. if yes, then that edge can be removed
    // - otherwise keep adding and update the union 
    
    // check max n value
    n := 0
    var a, b int
    for _, edge := range edges {
        a, b = edge[0], edge[1]

        if n < max(a, b) {
            n = max(a, b)
        }
    }

    ds := NewDisjointSet(n)
    var res []int 
    for _, edge := range edges {
        a, b = edge[0]-1, edge[1]-1

        if !ds.Union(a, b){
            res = edge
        }
    }
    return res
}

type DisjointSet struct {
    parents []int
}

func NewDisjointSet(n int) *DisjointSet {
    d := DisjointSet{
        parents: make([]int, n),
    } 
    for i:=0;i<n;i++{
        d.parents[i] = -1 // everyone is it's own union, with size = -(-1) = 1
    }

    return &d
}

// Find return the union id
func (d *DisjointSet) Find(x int) int {
    if d.parents[x] < 0 {
        return x
    }

    // do path compression
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

// Union will merge a and b into a single union. It returns false if it's already the same union, otherwise it's true
func (d *DisjointSet) Union(a, b int) bool {
    if a == b {
        return false
    }
    aId, bId := d.Find(a), d.Find(b)
    if aId == bId {
        return false
    }

    aSize, bSize := -d.parents[aId], -d.parents[bId]

    if bSize > aSize {
        d.parents[aId] = bId
        d.parents[bId] = -aSize-bSize
    }else{
        d.parents[bId] = aId
        d.parents[aId] = -aSize-bSize
    }
    return true
}