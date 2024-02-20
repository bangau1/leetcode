
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    n := len(nums)
    for i:=0;i<n;i++{
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        // fmt.Println("consider ", i)
        twoRes := twoSum(nums[i+1:], -nums[i])
        for _, twoSums := range twoRes {
            res = append(res, []int{nums[i], twoSums[0], twoSums[1]})
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
            for l + 1 < n && sorted[l+1] == sorted[l] {
                l++
            }
        }
        l++
    }
    return res
}