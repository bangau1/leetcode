func removeDuplicates(nums []int) int {
    n := len(nums)
    prev := nums[0]
    fill := 1

    for i:=1;i<n;i++{
        if prev != nums[i] {
            nums[fill] = nums[i]
            prev = nums[i]
            fill++
        }
    }
    return fill
}