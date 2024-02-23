func find(parents []int, x int) int {
    if parents[x] == -1 {
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

    // small join the large
    if aSize > bSize {
        parents[bId] = aId
        parents[aId] = -aSize -bSize
    }else{
        parents[aId] = bId
        parents[bId] = -aSize -bSize
    }
}

func longestConsecutive(nums []int) int {
    n := len(nums)
    if n <= 1 {
        return n
    }
    numSet := make(map[int]int)
    for i, num := range nums {
        numSet[num] = i
    }

    // to achieve this is O(N)
    // we can incrementally build a disjoint set and amortized them in O(N)
    parents := make([]int, n)
    for i:=0;i<n;i++{
        parents[i] = -1 // it means it is its own disjoint set
    }
    maxSize := 1
    for i:=1;i<n;i++{
        if idx, ok:=numSet[nums[i]-1];ok {
            union(parents, numSet[nums[i]], idx)
        }

        if idx, ok:=numSet[nums[i]+1];ok {
            union(parents, numSet[nums[i]], idx)
        }
    }
    for _, p := range parents {
        if p < 0 && maxSize < -p {
            maxSize = -p
        }
    }
    return maxSize
}