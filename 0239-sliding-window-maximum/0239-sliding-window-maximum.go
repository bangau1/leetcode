type data struct {
    val, index int
}
func printVal(arrs []data) {
    res := make([]int, 0)
    for _, datum := range arrs{
        res = append(res, datum.val)
    }
    fmt.Println(res)
}
func maxSlidingWindow(nums []int, k int) []int {
    if k == 1 {
        return nums
    }
    n := len(nums)
    if k > n {
        panic("unexpected k > n")
    }

    queue := make([]data, 0)

    for i:=0;i<k;i++{
        for len(queue) > 0 && nums[i] > queue[len(queue)-1].val {
            queue = queue[0:len(queue)-1]
        }
        queue = append(queue, data{nums[i], i})
    }
    res := make([]int, 0)
    res = append(res, queue[0].val)
    
    for i:=k;i<n;i++{
        start:= i-k+1
        for len(queue) > 0 && queue[0].index < start {
            queue = queue[1:]
        }
        for len(queue) > 0 && nums[i] > queue[len(queue)-1].val {
            queue = queue[0:len(queue)-1]
        }
        queue = append(queue, data{nums[i], i})
        
        res = append(res, queue[0].val)
    }
    return res
}