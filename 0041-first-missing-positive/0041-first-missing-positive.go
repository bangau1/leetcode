func firstMissingPositive(nums []int) int {
/*    
observations:
- we only care about the positive integers, 
- perhaps we can edit in place in the nums. if found A = nums[i], the set nums[A-1]=A
*/
    n := len(nums)
    var a, temp int
    for i:=0;i<n;i++{
        if nums[i] > 0 {
            a = nums[i]
            for a > 0 && a-1 < n && nums[a-1]!=a {
                temp = nums[a-1]
                nums[a-1] = a
                a = temp
            }
        }
    }
    for i:=0;i<n;i++{
        if nums[i] != i+1 {
            return i+1
        }
    }
    return n+1
}