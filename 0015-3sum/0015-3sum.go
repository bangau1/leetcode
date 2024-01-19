type pair struct{
    a, b, c int
}
func twoSum(sorted []int, start, end, target int) [][]int {
    l, r := start, end
    if sorted[l] + sorted[l+1] > target {
        return nil
    }
    if sorted[r-1] + sorted[r] < target {
        return nil
    }
    res := make([][]int, 0)
    for i:=start;i<=end;i++{
        search := target - sorted[i]
        idx := sort.Search(len(sorted[i+1:end+1]), func (idx int)bool {
            return sorted[i+1+idx] >= search
        }) + i+1

        if idx < len(sorted) && sorted[idx] == search {
            res = append(res, []int{sorted[i], sorted[idx]})
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