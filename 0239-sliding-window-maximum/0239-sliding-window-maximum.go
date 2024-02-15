func maxSlidingWindow(nums []int, k int) []int {
    if k == 1 {
        return nums
    }
    n := len(nums)
    res := make([]int, n - k + 1)
    deque := list.New()

    for i:=0;i<k;i++{
        for deque.Len() > 0 && nums[i] >= nums[deque.Back().Value.(int)] {
            deque.Remove(deque.Back())
        }
        deque.PushBack(i)
    }
    idx := 0
    res[idx] = nums[deque.Front().Value.(int)]
    idx++
    
    var start int
    for i:=k;i<n;i++{
        for deque.Len() > 0 && nums[i] >= nums[deque.Back().Value.(int)] {
            deque.Remove(deque.Back())
        }
        // valid range: start=i-k+1 end=i
        start = i - k + 1
        for deque.Len() > 0 && deque.Front().Value.(int) < start {
            deque.Remove(deque.Front())
        }
        deque.PushBack(i)
        res[idx] = nums[deque.Front().Value.(int)]
        idx++
    }
    return res
}