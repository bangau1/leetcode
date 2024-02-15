func maxSlidingWindow(nums []int, k int) []int {
    if k == 1 {
        return nums
    }
    var res []int
    deque := list.New()

    for i:=0;i<k;i++{
        for deque.Len() > 0 && nums[i] >= nums[deque.Back().Value.(int)] {
            deque.Remove(deque.Back())
        }
        deque.PushBack(i)
    }

    res = append(res, nums[deque.Front().Value.(int)])
    n := len(nums)
    
    for i:=k;i<n;i++{
        for deque.Len() > 0 && nums[i] >= nums[deque.Back().Value.(int)] {
            deque.Remove(deque.Back())
        }
        deque.PushBack(i)
        res = append(res, nums[deque.Front().Value.(int)])
    }
    return res
}