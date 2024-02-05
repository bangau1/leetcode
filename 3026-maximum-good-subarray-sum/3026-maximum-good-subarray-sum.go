func maximumSubarraySum(nums []int, k int) int64 {
    n := len(nums)
    prefixSum:= make([]int64, n+1)
    for i:=0;i<n;i++{
        prefixSum[i+1] = prefixSum[i] + int64(nums[i])
    }

    track := make(map[int][]int)
    for i:=0;i<n;i++{
        num := nums[i]
        track[num+k] = append(track[num+k], i) 
        track[num-k] = append(track[num-k], i)
    }

    rangeSum := func(start, end int) int64 {
        return prefixSum[end+1] - prefixSum[start]
    }

    res := int64(0)
    found := false
    for i:=0;i<n;i++{
        candidates := track[nums[i]]

        for _, cand := range candidates {
            if cand >= i {
                // fmt.Println(i, cand)
                if !found {
                    res = rangeSum(i, cand)
                }else{
                    res = max(res, rangeSum(i, cand))
                }
                found = true
            }
        }
    }

    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}