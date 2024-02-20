
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
    for i:=0;i<n;i++{
        if i > 0 && sorted[i] == sorted[i-1] {
            continue
        }
        search := target-sorted[i]

        idx := sort.SearchInts(sorted[i+1:], search)
        if idx+i+1 < n && sorted[i+1+idx] == search {
            res = append(res, []int{sorted[i], search})
        }
    }
    // fmt.Println(res)
    return res
}