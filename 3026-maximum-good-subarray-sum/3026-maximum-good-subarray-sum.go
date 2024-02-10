func maximumSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    prefixSum:= make([]int64, n+1)
    for i:=0;i<n;i++{
        prefixSum[i+1] = prefixSum[i] + int64(nums[i])
    }

    rangeSum := func(start, end int) int64 {
        if start > end {
            start, end = end, start
        }
        return prefixSum[end+1] - prefixSum[start]
    }

    res := int64(math.MinInt)
    track := make(map[int][]int)
    found := false
    for i:=0;i<n;i++{
        num := nums[i]
        if track[num+k] != nil {
            res = max(res, rangeSum(track[num+k][0], i))
            found = true
        }
        if track[num-k] != nil {
            res = max(res, rangeSum(track[num-k][0], i))
            found = true
        }

        if track[num] == nil {
            track[num] = []int{i}
        }else{
            if prefixSum[i] < prefixSum[track[num][0]] {
                track[num][0] = i
            }
        }
    }
    // fmt.Println(track)
    if !found {
        return 0
    }
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}