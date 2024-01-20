
func twoSum(sorted []int, start, end, target int) [][]int {
    l, r := start, end
    if sorted[l] + sorted[l+1] > target {
        return nil
    }
    if sorted[r-1] + sorted[r] < target {
        return nil
    }
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
    res := make([][]int, 0)
    for i:=0;i<len(nums)-2;i++{
        data := twoSum(nums, i+1, len(nums)-1, -nums[i])
        for _, datum := range data {
            arr := []int{nums[i], datum[0], datum[1]}
                res = append(res, arr)
        }
        for i + 1 <len(nums) && nums[i+1] == nums[i] {
            i++
        }
    }
    return res
}