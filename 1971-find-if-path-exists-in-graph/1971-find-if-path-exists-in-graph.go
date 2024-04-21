type UnionSet struct {
    arr []int
    // maxSize int
}

type Edge struct {
    from int
    to int
}

func NewUnionSet(nodeSize int, edges []Edge) *UnionSet {
    res := UnionSet{}
    res.arr = make([]int, nodeSize)
    for i := 0; i<nodeSize;i++{
        res.arr[i] = -1 // each is group set with 1 size (negative value)
    }

    for _, edge := range edges {
        res.Union(edge.from, edge.to)
    }

    return &res
}
func (u *UnionSet) FindGroup(i int) int {
    // fmt.Println("FindGroup", i)
    // fmt.Println("Arr", u.arr)
    // check if i is the root of the group 
    if u.arr[i] < 0 {
        return i
    }
    path := make([]int, 0)
    
    // while it's not a group
    for u.arr[i] >= 0 {
        fmt.Println(i, u.arr[i])
        path = append(path, i)
        i = u.arr[i]
    }
    // now i should be the root
    
    // path compression optimization
    for _, p := range path {
        u.arr[p] = i
    }

    return i
}

func (u *UnionSet) Union(f, p int) {
    // fmt.Println("??Union", f, p, u.arr)
    fGroup := u.FindGroup(f)
    pGroup := u.FindGroup(p)
    if fGroup == pGroup {
        return
    }

    fGroupSize := u.arr[fGroup] * -1
    pGroupSize := u.arr[pGroup] * -1

    

    // fmt.Println("fGroup", fGroup, "pGroup", pGroup)
    // fmt.Println("fGroupSize", fGroupSize, "pGroupSize", pGroupSize)
    if fGroupSize <= 0 || pGroupSize <= 0 {
        panic("unexpected logic")
    }
    // merge f to p set
    if fGroupSize < pGroupSize {
        u.arr[pGroup] = -fGroupSize - pGroupSize // store as negative value
        u.arr[fGroup] = pGroup
        // fmt.Println("Merge f to p", u.arr)
    } else { // merge p to f set   
        u.arr[fGroup] = -fGroupSize - pGroupSize // store as negative value
        u.arr[pGroup] = fGroup
        // fmt.Println("Merge p to f", u.arr)
    }
}

func validPath(n int, edges [][]int, source int, destination int) bool {
    edgesArr := make([]Edge, 0)
    for _, edge := range edges {
        edgeStruct := Edge {
            from: edge[0],
            to: edge[1],
        }
        edgesArr = append(edgesArr, edgeStruct)
    }

    unionSet := NewUnionSet(n, edgesArr)
    // fmt.Println("===========", unionSet.arr)
    sourceGroup := unionSet.FindGroup(source)
    destGroup := unionSet.FindGroup(destination)

    return sourceGroup == destGroup
}