func outerTrees(trees [][]int) [][]int {
    if len(trees) <= 3 {
        return trees
    }
    bottomTree := trees[0]

    for i:=1;i<len(trees);i++{
        if bottomTree[1] > trees[i][1] {
            bottomTree = trees[i]
        }else if bottomTree[1] == trees[i][1] && bottomTree[0] > trees[i][0] {
            bottomTree = trees[i]
        }
    }
    sorted := make([][]int, 0)
    for i:=0;i<len(trees);i++{
        if !(bottomTree[0] == trees[i][0] && bottomTree[1] == trees[i][1]) {
            sorted = append(sorted, trees[i])
        }
    }

    

    // sort all remaining trees relative to bottomTree by its -cotan
    // quadran 1 to 2
    // 0 degree = -cotant(x,y) = -inf
    // 180 degree = -cotant(x,y) = inf
    sort.Slice(sorted, func(a, b int) bool {
        aVal := -cotant(bottomTree, sorted[a])
        bVal := -cotant(bottomTree, sorted[b])
        if aVal == bVal {
            distA := (sorted[a][0]-bottomTree[0])*(sorted[a][0]-bottomTree[0]) +  (sorted[a][1]-bottomTree[1])*(sorted[a][1]-bottomTree[1])
            distB := (sorted[b][0]-bottomTree[0])*(sorted[b][0]-bottomTree[0]) +  (sorted[b][1]-bottomTree[1])*(sorted[b][1]-bottomTree[1])
            return distA < distB
        }

        return aVal < bVal
    })

    i:=len(sorted)-1
    for i>=0 {
        if crossProduct(bottomTree, sorted[len(sorted)-1], sorted[i]) != 0{
            break
        }
        i--
    }
    for l,r:=i+1,len(sorted)-1; l<r; {
        sorted[l], sorted[r] = sorted[r], sorted[l]
        l++
        r--
    }

    // fmt.Println(bottomTree)
    // fmt.Println("sorted", sorted)
    

    res := make([][]int, 0)
    res = append(res, bottomTree)
    res = append(res, sorted[0])

    for i:=1;i<len(sorted);i++{
        a := res[len(res)-2]
        b := res[len(res)-1]

        c := sorted[i]

        for len(res) >= 3 && !isCounterClockwise(a, b, c) {
            res = res[0:len(res)-1]
            a = res[len(res)-2]
            b = res[len(res)-1]
        }
        res = append(res, c)
        
    }
    return res
}

var inf =  float64(math.MaxInt)
var nInf = float64(math.MinInt)

// a is the relative position
func cotant(a, b []int) float64 {
    ax, ay := float64(a[0]), float64(a[1])
    bx, by := float64(b[0]), float64(b[1])

    

    // contant = dx/dy
    dx := bx - ax
    dy := by - ay

    if dx == 0 && dy == 0 {
        panic("undefined")
    }

    if dx == 0 {
        return 0
    }

    if dy == 0 {
        if dx < 0 {
            return nInf
        }else{
            return inf
        }
    }
    return dx/dy

}


func isCounterClockwise(a, b, c []int) bool {

    cross := crossProduct(a, b, c)
    // fmt.Println(a, b, c, cross)

    // if the cross product positive, it means it's counter clockwise
    return cross >= 0
}

func crossProduct(a, b, c []int) int {
    ab := []int{b[0]-a[0], b[1]-a[1]}
    bc := []int{c[0]-b[0], c[1]-b[1]}

    return ab[0]*bc[1] - ab[1]*bc[0]
}