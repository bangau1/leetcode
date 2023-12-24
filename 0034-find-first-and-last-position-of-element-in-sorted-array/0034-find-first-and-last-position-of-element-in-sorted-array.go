func searchRange(nums []int, target int) []int {
    n := len(nums)
    if n == 0 {
        return []int{-1, -1}
    }

    lowerBound := sort.Search(n, func(idx int) bool {
        return nums[idx] >= target
    })

    if lowerBound < n && nums[lowerBound] == target {
        
        upperBound := sort.Search(n, func(idx int) bool {
            return nums[idx] > target
        })
        return []int{lowerBound, upperBound-1}
    }

    return []int{-1, -1}

}