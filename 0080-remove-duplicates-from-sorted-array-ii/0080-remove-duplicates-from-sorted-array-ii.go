func removeDuplicates(nums []int) int {
    n := len(nums)
    prev := nums[0]
    c := 1
    fill := 1
    for i:=1;i<n;i++{
        // we fill the data in place if prev != nums[i]
        // but we still 
        if prev != nums[i] {
            nums[fill] = nums[i]
            fill++
            c = 1
        }else if c < 2 {
            nums[fill] = nums[i]
            fill++
            c++
        }

        // fmt.Println(i, nums[0:fill])
        prev = nums[i]
    }
    return fill
}