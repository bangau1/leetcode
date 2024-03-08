func maxFrequencyElements(nums []int) int {
    counter := make([]int, 101)
    max := 0
    for _, num := range nums{
        counter[num]++
        if counter[num] > max {
            max = counter[num]
        }
    }
    res := 0
    for i:=1;i<=100;i++{
        if counter[i] == max {
            res += max
        }
    }
    return res

}