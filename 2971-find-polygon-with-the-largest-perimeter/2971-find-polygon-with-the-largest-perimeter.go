func largestPerimeter(nums []int) int64 {
    sort.Ints(nums)
    n := len(nums)
    res := int64(-1)
    prevSum := int64(nums[0] + nums[1])
    for i:=2;i<n;i++{
        l := int64(nums[i])
        if prevSum > l {
            res = max(res, l + prevSum)
        }
        prevSum += l
    }
    return res
}