
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    n := len(nums)
    dict := make(map[[3]int]bool)
    for i:=0;i<n;i++{
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        twoRes := twoSum(nums[i+1:], -nums[i])
        for _, twoSums := range twoRes {
            if _, ok:=dict[[3]int{nums[i], twoSums[0], twoSums[1]}]; !ok {
                res = append(res, []int{nums[i], twoSums[0], twoSums[1]})
                dict[[3]int{nums[i], twoSums[0], twoSums[1]}] = true
            }
        }
    }

    return res
}

func twoSum(sorted []int, target int) [][]int {
    n := len(sorted)
    if n < 2 {
        return nil
    }
    var res [][]int
    l, r := 0, n-1
    for l < r {
        sum := sorted[l] + sorted[r]
        if sum < target {
            l++
            continue
        }else if sum > target {
            r--
            continue
        }else{
            res = append(res, []int{sorted[l], sorted[r]})
        }
        l++
    }
    return res
}