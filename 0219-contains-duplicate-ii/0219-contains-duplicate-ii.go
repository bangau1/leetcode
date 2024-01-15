func containsNearbyDuplicate(nums []int, k int) bool {
    prevIndex := make(map[int]int)

    for i:=0;i<len(nums);i++ {
        num := nums[i]

        if prevIdx, ok:=prevIndex[num]; ok {
            if i-prevIdx <= k{
                return true
            }
        }
        prevIndex[num] = i
    }
    return false
}