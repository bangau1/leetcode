func maxSlidingWindow(nums []int, k int) []int {
    // we can use monotonic-dequeu to maintain the max element on the window
    // the idea is to only insert the element in the window when it is bigger from current queue's back element (non-increaseing queue)
    q := make([]int, 0)
    n := len(nums)
    // add the first k element
    for i:=0;i<k;i++{
        for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
            q = q[:len(q)-1] // pop the back
        }
        q = append(q, i)
    }
    res := make([]int, 0)
    res = append(res, nums[q[0]])
    var validIdx int
    for i:=k;i < n;i++{
        // remove old element from the queue
        validIdx = i - k + 1
        for len(q) > 0 && q[0] < validIdx {
            q = q[1:]
        }
        for len(q) > 0 && nums[q[len(q)-1]] < nums[i] {
            q = q[:len(q)-1] // pop the back
        }
        q = append(q, i)
        res = append(res, nums[q[0]])
    }
    return res
}
