func removeDuplicates(nums []int) int {
    prev := nums[0]

    n := len(nums)
    i := 1
    left := 1
    for i < n {
        for i < n && nums[i] == prev {
            i++
        }
        if i >= n {
            break
        }

        prev = nums[i]
        nums[left] = nums[i]
        i++
        left++
    }
    return left
}