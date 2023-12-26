func firstMissingPositive(nums []int) int {
    // we can't use min/max, since the missing can be in the middle of min-max
    // observations:
    // n = length(nums)
    // the smallest missing positive number can be anywhere of 1-n (inclusive)
    // in order to avoid additional space, we can modify the existing nums array
    // if we find a[i] that is positive, we set a[a[i]] = a[i]
    //
    // basically we want to convert the nums as hash map to indicate that the number has been marked
    // the unmarked one is basically the missing positive integer


    for i:=1;i<=len(nums);i++{
        tmp := nums[i-1]
        for tmp <= len(nums) && tmp >= 1 && tmp != nums[tmp-1]{
            tmp, nums[tmp-1] = nums[tmp-1], tmp
        }

    }
    // fmt.Println(nums)
    for i:=1;i<=len(nums);i++{
        if nums[i-1] != i {
            return i
        }
    }
    return len(nums)+1
}