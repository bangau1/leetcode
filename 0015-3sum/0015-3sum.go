type pair struct{
    a, b, c int
}
func twoSum(sorted []int, start, end, target int) [][]int {
    l, r := start, end
    res := make([][]int, 0)
    for l < r {
        sum := sorted[l] + sorted[r]
        if sum < target {
            l++
        }else if sum > target {
            r--
        }else{
            res = append(res, []int{sorted[l], sorted[r]})
            l++
        }
    }
    return res
}

func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    flag := make(map[pair]bool)
    res := make([][]int, 0)
    for i:=0;i<len(nums)-2;i++{
        data := twoSum(nums, i+1, len(nums)-1, -nums[i])
        for _, datum := range data {
            p := pair{nums[i], datum[0], datum[1]}
            if !flag[p] {
                res = append(res, []int{nums[i], datum[0], datum[1]})
                flag[p] = true
            }
        }
    }
    return res
}