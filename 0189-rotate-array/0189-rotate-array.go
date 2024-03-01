func rotate(nums []int, k int)  {
    n := len(nums)
    k = k % n
    if k == 0 {
        return
    }

    // the basic idea is to reverse the nums
    // [1,2,3,4,5] -> [5,4,3,2,1]
    // then for the first k and the last n-k element, we should reverse it back independently
    // [4,5,1,2,3]
    reverse(nums)
    reverse(nums[:k])
    reverse(nums[k:])
}

func reverse(nums []int) {
    n := len(nums)
    l, r := 0, n-1
    for l < r {
        nums[l], nums[r] = nums[r], nums[l]
        l++
        r--
    }
}
