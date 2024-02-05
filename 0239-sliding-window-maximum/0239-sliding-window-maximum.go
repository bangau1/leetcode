func maxSlidingWindow(nums []int, k int) []int {
    if k == 1 {
        return nums
    }
    n := len(nums)
    if k > n {
        panic("unexpected k > n")
    }

    queue := make([]int, 0)

    for i:=0;i<k;i++{
        for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
            queue = queue[0:len(queue)-1]
        }
        queue = append(queue, i)
    }
    res := make([]int, 0)
    res = append(res, nums[queue[0]])
    
    start := 1
    for i:=k;i<n;i++{
        start = i - k + 1
        for len(queue) > 0 && queue[0] < start {
            queue = queue[1:]
        }
        for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
            queue = queue[0:len(queue)-1]
        }
        queue = append(queue, i)
        
        res = append(res, nums[queue[0]])
    }
    return res
}