func searchInsert(nums []int, target int) int {
    return sort.Search(len(nums), func(idx int) bool {
        return nums[idx] >= target
    })
}