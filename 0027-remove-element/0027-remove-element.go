func removeElement(nums []int, val int) int {
    n := len(nums)
    temp := make([]int, 0)
    for i:=0;i<n;i++{
        if nums[i] != val {
            temp = append(temp, nums[i])
        }
    }

    copy(nums[0:len(temp)], temp)

    return len(temp)
}